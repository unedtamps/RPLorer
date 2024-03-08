// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package repository

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	ChangePostStatus(ctx context.Context, arg ChangePostStatusParams) error
	CountCommentPost(ctx context.Context, postID uuid.UUID) (int64, error)
	CoutAccoutPost(ctx context.Context, accountID uuid.UUID) (int64, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (*CreateAccountRow, error)
	CreateAccountDetail(ctx context.Context, arg CreateAccountDetailParams) error
	CreateAccountFollow(ctx context.Context, arg CreateAccountFollowParams) error
	CreateAdminAccount(ctx context.Context, arg CreateAdminAccountParams) (uuid.UUID, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (uuid.UUID, error)
	CreateLikedPost(ctx context.Context, arg CreateLikedPostParams) error
	CreateNewComment(ctx context.Context, arg CreateNewCommentParams) (*CreateNewCommentRow, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (uuid.UUID, error)
	CreatePostImages(ctx context.Context, arg CreatePostImagesParams) error
	CreteNewCommentWithParent(ctx context.Context, arg CreteNewCommentWithParentParams) (*CreteNewCommentWithParentRow, error)
	DeleteComment(ctx context.Context, id uuid.UUID) error
	DeleteCommentByPostId(ctx context.Context, postID uuid.UUID) error
	DeleteImageByPostId(ctx context.Context, postID uuid.UUID) error
	DeleteLikedByPostId(ctx context.Context, postID uuid.UUID) error
	DeleteLikedPost(ctx context.Context, arg DeleteLikedPostParams) error
	DeletePostById(ctx context.Context, id uuid.UUID) error
	GetAccountByEmail(ctx context.Context, email string) (*GetAccountByEmailRow, error)
	GetAccountById(ctx context.Context, id uuid.UUID) (*Account, error)
	InsertNewPost(ctx context.Context, arg InsertNewPostParams) (*InsertNewPostRow, error)
	InsertNewPostImages(ctx context.Context, arg InsertNewPostImagesParams) error
	QueryCommentById(ctx context.Context, id uuid.UUID) (*QueryCommentByIdRow, error)
	QueryCommentByPost(ctx context.Context, arg QueryCommentByPostParams) ([]*QueryCommentByPostRow, error)
	QueryCommentChild(ctx context.Context, parrentID uuid.NullUUID) ([]*QueryCommentChildRow, error)
	QueryGetAccoutFromLikedByPostId(ctx context.Context, postID uuid.UUID) ([]uuid.UUID, error)
	QueryImageByPostID(ctx context.Context, postID uuid.UUID) ([]string, error)
	QueryPostByUserID(ctx context.Context, arg QueryPostByUserIDParams) ([]*QueryPostByUserIDRow, error)
	QueryPostandImages(ctx context.Context, id uuid.UUID) ([]*QueryPostandImagesRow, error)
	QueryUserIdfromPost(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	UpdateCommentBody(ctx context.Context, arg UpdateCommentBodyParams) (*UpdateCommentBodyRow, error)
	UpdateCommentCount(ctx context.Context, arg UpdateCommentCountParams) error
	UpdateCommentCountIncrement(ctx context.Context, id uuid.UUID) error
	UpdateDecreaseGetLikeUserDetail(ctx context.Context, accountID uuid.UUID) error
	UpdateDecreaseGiveLikeUserDetail(ctx context.Context, accountID uuid.UUID) error
	UpdateFollowersCount(ctx context.Context, accountID uuid.UUID) error
	UpdateFollowingCount(ctx context.Context, accountID uuid.UUID) error
	UpdateGetLikeUserDetail(ctx context.Context, accountID uuid.UUID) error
	UpdateGiveLikeUserDetail(ctx context.Context, accountID uuid.UUID) error
	UpdateLikeCountDecrement(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	UpdateLikeCountIncrement(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	UpdatePostCaption(ctx context.Context, arg UpdatePostCaptionParams) (*UpdatePostCaptionRow, error)
	UpdatePostDetailCount(ctx context.Context, arg UpdatePostDetailCountParams) error
	UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) error
}

var _ Querier = (*Queries)(nil)
