package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	r "github.com/unedtamps/go-backend/internal/repository"
)

type PostService struct {
	repo  *r.Store
	cache *redis.Client
}

type PostServiceI interface {
	CreatePostService(
		ctx context.Context,
		id string,
		caption string,
		file *[]multipart.FileHeader,
	) (*r.InsertNewPostRow, *ErrorService)
	CreateImagePost(
		ctx context.Context,
		post_id uuid.UUID,
		file []uuid.UUID,
		fileExtension []string,
	) *ErrorService
	GetPostByAccountService(
		ctx context.Context,
		account_id uuid.UUID,
		limit int64,
		offset int64,
	) ([]*r.QueryPostByUserIDRow, int64, *ErrorService)
	GetDetailPostService(ctx context.Context, post_id uuid.UUID) (interface{}, *ErrorService)
	LikePostService(ctx context.Context, post_id uuid.UUID, account_id uuid.UUID) *ErrorService
	UserDeletePost(ctx context.Context, post_id uuid.UUID, account_id uuid.UUID) *ErrorService
	UpdatePostCaption(
		ctx context.Context,
		post_id, account_id uuid.UUID,
		caption string,
	) (*r.UpdatePostCaptionRow, *ErrorService)
	UnLikePostService(ctx context.Context, post_id uuid.UUID, account_id uuid.UUID) *ErrorService
	AdminDeletePost(ctx context.Context, postId uuid.UUID) *ErrorService
	AddComment(
		ctx context.Context,
		body string,
		postId, account_id, parrent_id uuid.UUID,
	) (interface{}, *ErrorService)
	GetCommentPost(
		ctx context.Context,
		post_id uuid.UUID,
		limit int64,
		offset int64,
	) (interface{}, int64, *ErrorService)
	GetCommentById(
		ctx context.Context,
		comment_id uuid.UUID,
	) (interface{}, *ErrorService)
	EditComment(
		ctx context.Context,
		body string,
		comment_id uuid.UUID,
		account_id uuid.UUID,
	) (interface{}, *ErrorService)
	DeleteUserComment(ctx context.Context, comment_id uuid.UUID, accId uuid.UUID) *ErrorService
}

func newPostService(repo *r.Store, cache *redis.Client) PostServiceI {
	return &PostService{repo, cache}
}

func (p *PostService) CreatePostService(
	ctx context.Context,
	account_id string,
	caption string,
	file *[]multipart.FileHeader,
) (*r.InsertNewPostRow, *ErrorService) {
	// insert new post
	post, err := p.repo.InsertNewPost(ctx, r.InsertNewPostParams{
		AccountID: uuid.MustParse(account_id),
		Caption:   caption,
	})
	if err != nil {
		return nil, newError(err, 500)
	}
	return post, nil
}

func (p *PostService) CreateImagePost(
	ctx context.Context,
	post_id uuid.UUID,
	img_id []uuid.UUID,
	fileExtension []string,
) *ErrorService {
	for i, x := range img_id {
		err := p.repo.InsertNewPostImages(ctx, r.InsertNewPostImagesParams{
			ID:     x,
			PostID: post_id,
			ImgUrl: fmt.Sprintf("/static/post/%s%s", x, fileExtension[i]),
		})
		if err != nil {
			return newError(err, 500)
		}
	}
	return nil
}

func (p *PostService) GetPostByAccountService(
	ctx context.Context,
	account_id uuid.UUID,
	limit int64,
	offset int64,
) ([]*r.QueryPostByUserIDRow, int64, *ErrorService) {
	posts, err := p.repo.QueryPostByUserID(ctx, r.QueryPostByUserIDParams{
		AccountID: account_id,
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, 0, newError(errors.New("post not found"), 404)
		default:
			return nil, 0, newError(err, 500)
		}
	}
	if posts == nil {
		return nil, 0, newError(errors.New("post not found"), 404)
	}
	total, err := p.repo.CoutAccoutPost(ctx, account_id)
	if err != nil {
		return nil, 0, newError(err, 500)
	}
	return posts, total, nil
}

type postDetailResponse struct {
	Id           uuid.UUID `json:"id"`
	Caption      string    `json:"caption"`
	LikeCount    int64     `json:"like_count"`
	CommentCount int64     `json:"comment_count"`
	ImageUrl     []string  `json:"image_url"`
}

func (p *PostService) GetDetailPostService(
	ctx context.Context,
	post_id uuid.UUID,
) (interface{}, *ErrorService) {
	post, err := p.repo.QueryPostandImages(ctx, post_id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, newError(errors.New("post not found"), 404)
		default:
			return nil, newError(err, 500)
		}
	}
	if post == nil {
		return nil, newError(errors.New("post not found"), 404)
	}
	var post_rest postDetailResponse
	post_rest.Id = post[0].PostID
	post_rest.Caption = post[0].Caption
	post_rest.LikeCount = post[0].LikesCount
	post_rest.CommentCount = post[0].CommentCount

	for _, p := range post {
		if p.ImgUrl.Valid {
			post_rest.ImageUrl = append(post_rest.ImageUrl, p.ImgUrl.String)
		}
	}
	return post_rest, nil
}

func (p *PostService) LikePostService(
	ctx context.Context,
	post_id uuid.UUID,
	account_id uuid.UUID,
) *ErrorService {
	err := p.repo.LikePostWithTx(ctx, account_id, post_id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return newError(errors.New("post already liked"), 400)
		} else if strings.Contains(err.Error(), "violates foreign key constraint") {
			return newError(errors.New("post not found"), 404)
		}
		return newError(err, 500)
	}
	return nil
}

func (p *PostService) UnLikePostService(
	ctx context.Context,
	post_id uuid.UUID,
	account_id uuid.UUID,
) *ErrorService {
	err := p.repo.UnlikePostWithTx(ctx, account_id, post_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return newError(errors.New("post already unliked"), 404)
		}
		return newError(err, 500)
	} else {
		return nil
	}
}

func (p *PostService) UserDeletePost(
	ctx context.Context,
	post_id uuid.UUID,
	account_id uuid.UUID,
) *ErrorService {
	err := p.repo.ChangePostStatus(ctx, r.ChangePostStatusParams{
		ID:        post_id,
		Status:    r.StatusDELETED,
		AccountID: account_id,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return newError(errors.New("post not found"), 404)
		}
		return newError(err, 500)
	}
	return nil
}

func (p *PostService) UpdatePostCaption(
	ctx context.Context,
	post_id, account_id uuid.UUID,
	caption string,
) (*r.UpdatePostCaptionRow, *ErrorService) {
	post, err := p.repo.UpdatePostCaption(ctx, r.UpdatePostCaptionParams{
		Caption: caption,
		ID:      post_id,
	})
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, newError(errors.New("post not found"), 404)
		default:
			return nil, newError(err, 500)
		}
	}
	if post.AccountID != account_id {
		return nil, newError(errors.New("unauthorized"), 401)
	}
	return post, nil
}

func (p *PostService) AdminDeletePost(ctx context.Context, postId uuid.UUID) *ErrorService {
	err := p.repo.DeletePostWithTx(ctx, postId)
	if err != nil {
		if err == sql.ErrNoRows {
			return newError(errors.New("post not found"), 404)
		}
		return newError(err, 500)
	}
	return nil
}

func (p *PostService) AddComment(
	ctx context.Context,
	body string,
	postId, account_id, parrent_id uuid.UUID,
) (interface{}, *ErrorService) {
	if parrent_id == uuid.Nil {
		comm, err := p.repo.CreateNewComment(ctx, r.CreateNewCommentParams{
			AccountID: account_id,
			PostID:    postId,
			Body:      body,
		})
		if err != nil {
			return nil, newError(err, 500)
		}
		err = p.repo.UpdateCommentCountIncrement(ctx, postId)
		if err != nil {
			return nil, newError(err, 500)
		}
		return comm, nil
	} else {
		comm, err := p.repo.CreteNewCommentWithParent(ctx, r.CreteNewCommentWithParentParams{
			AccountID: account_id,
			PostID:    postId,
			ParrentID: uuid.NullUUID{UUID: parrent_id, Valid: true},
			Body:      body,
		})
		if err != nil {
			return nil, newError(err, 500)
		}
		err = p.repo.UpdateCommentCountIncrement(ctx, postId)
		if err != nil {
			return nil, newError(err, 500)
		}
		return comm, nil
	}
}

func (p *PostService) EditComment(
	ctx context.Context,
	body string,
	comment_id uuid.UUID,
	account_id uuid.UUID,
) (interface{}, *ErrorService) {
	id, err := p.repo.QueryCommentById(ctx, comment_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, newError(errors.New("comment not found"), 404)
		}
		return nil, newError(err, 500)
	}
	if id.AccountID != account_id {
		return nil, newError(errors.New("unauthorize"), 403)
	}
	if id.Status == r.StatusDELETED {
		return nil, newError(errors.New("comment already deleted"), 442)
	}
	comment, err := p.repo.UpdateCommentBody(ctx, r.UpdateCommentBodyParams{
		Body: body,
		ID:   comment_id,
	})
	if err != nil {
		return nil, newError(err, 500)
	}
	return comment, nil
}

func (p *PostService) GetCommentById(
	ctx context.Context,
	comment_id uuid.UUID,
) (interface{}, *ErrorService) {
	comm, err := p.repo.QueryCommentById(ctx, comment_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, newError(errors.New("comment not found"), 404)
		}
		return nil, newError(err, 500)
	}
	return comm, nil
}

type commentResponse struct {
	Parent *r.QueryCommentByPostRow  `json:"parent"`
	Child  []*r.QueryCommentChildRow `json:"child"`
}

func (p *PostService) GetCommentPost(
	ctx context.Context,
	post_id uuid.UUID,
	limit int64,
	offset int64,
) (interface{}, int64, *ErrorService) {
	comment_map := make(map[*r.QueryCommentByPostRow][]*r.QueryCommentChildRow)
	comments_query, err := p.repo.QueryCommentByPost(ctx, r.QueryCommentByPostParams{
		PostID: post_id,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 0, newError(errors.New("comment not found"), 404)
		}
		return nil, 0, newError(err, 500)
	}
	count, err := p.repo.CountCommentPost(ctx, post_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 0, newError(errors.New("comment not found"), 404)
		}
		return nil, 0, newError(err, 500)
	}
	for _, c := range comments_query {
		child, err := p.repo.QueryCommentChild(ctx, uuid.NullUUID{
			UUID:  c.ID,
			Valid: true,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, 0, newError(err, 500)
		} else {
			comment_map[c] = child
		}
	}
	var res []commentResponse
	for i, v := range comment_map {
		res = append(res, commentResponse{
			Parent: i,
			Child:  v,
		})
	}
	return res, count, nil
}

func (p *PostService) DeleteUserComment(
	ctx context.Context,
	comment_id uuid.UUID,
	accid uuid.UUID,
) *ErrorService {
	comm, err := p.repo.QueryCommentById(ctx, comment_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return newError(errors.New("comment not found"), 404)
		}
		return newError(err, 500)
	}
	if comm.AccountID != accid {
		return newError(errors.New("unauthorize"), 403)
	}
	if comm.Status == r.StatusDELETED {
		return newError(errors.New("comment already deleted"), 404)
	}
	err = p.repo.DeleteComment(ctx, comment_id)
	if err != nil {
		return newError(err, 500)
	}
	return nil
}
