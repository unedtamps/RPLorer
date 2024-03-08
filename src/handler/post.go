package handler

import (
	"errors"
	"fmt"
	"math"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/unedtamps/go-backend/src/middleware"
	"github.com/unedtamps/go-backend/src/service"
	"github.com/unedtamps/go-backend/util"
)

type PostHandler struct {
	service.PostServiceI
}

type PostHandlerI interface {
	CreatePostHandler(c *gin.Context)
	GetPostByAccountHandler(c *gin.Context)
	GetDetailPostByIDHandler(c *gin.Context)
	LikePostHandler(c *gin.Context)
	UnLikePostHandler(c *gin.Context)
	UserDeletePostHandler(c *gin.Context)
	UpdatePostCaptionHandler(c *gin.Context)
	AdminDeletePostHandler(c *gin.Context)
	AddNewCommentHandler(c *gin.Context)
	GetCommentByIdHandler(c *gin.Context)
	GetCommentPostHandler(c *gin.Context)
	EditCommentHandler(c *gin.Context)
	DeleteCommentHandler(c *gin.Context)
}

func newPostHandler(postService service.PostServiceI) PostHandlerI {
	return &PostHandler{postService}
}

type createPostParams struct {
	Caption string                  `form:"caption" binding:"required"`
	File    *[]multipart.FileHeader `form:"file"`
}

func (h *PostHandler) CreatePostHandler(c *gin.Context) {
	var postParams createPostParams
	var fileUrl []uuid.UUID
	var fileExtension []string
	if err := c.Bind(&postParams); err != nil {
		util.BadRequest(c, err)
		return
	}
	id := c.Value("cred").(middleware.Credentials).Id
	post, err := h.CreatePostService(c, id, postParams.Caption, postParams.File)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	if postParams.File != nil {

		for _, f := range *postParams.File {
			if filepath.Ext(f.Filename) != ".jpg" && filepath.Ext(f.Filename) != ".jpeg" &&
				filepath.Ext(f.Filename) != ".png" {
				util.BadRequest(c, errors.New("File must be an image"))
				return
			} else if f.Size > (2 << 20) {
				util.BadRequest(c, errors.New("File size must be less than 2MB"))
				return
			}
		}

		for _, f := range *postParams.File {
			img_id := uuid.MustParse(util.GenerateUUID())
			ext := filepath.Ext(f.Filename)
			img_url := fmt.Sprintf("./storage/post/%s%s", img_id, ext)
			c.SaveUploadedFile(&f, img_url)
			fileUrl = append(fileUrl, img_id)
			fileExtension = append(fileExtension, ext)
		}
		err := h.CreateImagePost(c, post.ID, fileUrl, fileExtension)
		if err != nil {
			util.UnknownError(c, err.Error, err.Code)
			return
		}
	}
	util.ResponseCreated(c, "Post created successfully", post)
}

func (h *PostHandler) GetPostByAccountHandler(c *gin.Context) {
	var paginate paginateForm
	if err := c.ShouldBindQuery(&paginate); err != nil {
		util.BadRequest(c, err)
		return
	}
	limit, offset := newPageiante(&paginate)

	id := c.Value("cred").(middleware.Credentials).Id
	posts, total, err := h.GetPostByAccountService(c, uuid.MustParse(id), limit, offset)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}

	metadata := util.WithMetadata(
		paginate.Page,
		total,
		int64(math.Ceil(float64(total)/float64(paginate.Page_size))),
		nil,
	)
	util.ResponseData(c, "Post retrieved", &metadata, posts)
}

func (h *PostHandler) GetDetailPostByIDHandler(c *gin.Context) {
	id := c.Param("id")
	new_id, e := uuid.Parse(id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	post, err := h.GetDetailPostService(c, new_id)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseData(c, "Post retrieved", nil, post)
}

func (h *PostHandler) LikePostHandler(c *gin.Context) {
	id := c.Param("id")
	acc_id := c.Value("cred").(middleware.Credentials).Id
	new_id, e := uuid.Parse(id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	new_acc_id, e := uuid.Parse(acc_id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}

	err := h.LikePostService(c, new_id, new_acc_id)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseOk(c, "Post liked")
}
func (h *PostHandler) UnLikePostHandler(c *gin.Context) {

	id := c.Param("id")
	acc_id := c.Value("cred").(middleware.Credentials).Id
	new_id, e := uuid.Parse(id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	new_acc_id, e := uuid.Parse(acc_id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	err := h.UnLikePostService(c, new_id, new_acc_id)
	if err != nil {
		util.BadRequest(c, err.Error)
		return
	}
	util.ResponseOk(c, "unlike post")
}

func (h *PostHandler) UserDeletePostHandler(c *gin.Context) {
	id := c.Param("id")
	acc_id := c.Value("cred").(middleware.Credentials).Id
	new_id, e := uuid.Parse(id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	new_acc_id := uuid.MustParse(acc_id)
	err := h.UserDeletePost(c, new_id, new_acc_id)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseOk(c, "Post deleted")
}

type postEditCaptionParams struct {
	Caption string `json:"caption" binding:"required"`
}

func (h *PostHandler) UpdatePostCaptionHandler(c *gin.Context) {
	var captionParams postEditCaptionParams
	if err := c.ShouldBindJSON(&captionParams); err != nil {
		util.BadRequest(c, err)
		return
	}
	id := c.Param("id")
	acc_id := c.Value("cred").(middleware.Credentials).Id
	new_id, e := uuid.Parse(id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	new_acc_id := uuid.MustParse(acc_id)
	post, err := h.UpdatePostCaption(c, new_id, new_acc_id, captionParams.Caption)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseData(c, "Caption updated", nil, post)
}

func (h *PostHandler) AdminDeletePostHandler(c *gin.Context) {
	id := c.Param("id")
	new_id, e := uuid.Parse(id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	err := h.AdminDeletePost(c, new_id)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseOk(c, "Post deleted")
}

type commentParams struct {
	PostId string `json:"post_id"    binding:"required"`
	Body   string `json:"body"       binding:"required"`
	Parent string `json:"parrent_id"`
}

func (h *PostHandler) AddNewCommentHandler(c *gin.Context) {
	var param commentParams
	if err := c.ShouldBindJSON(&param); err != nil {
		util.BadRequest(c, err)
		return
	}
	id := c.Value("cred").(middleware.Credentials).Id
	post_id, e := uuid.Parse(param.PostId)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	var parent_id uuid.UUID
	if param.Parent != "" {
		parent_id, e = uuid.Parse(param.Parent)
		if e != nil {
			util.BadRequest(c, e)
			return
		}
	} else {
		parent_id = uuid.Nil
	}
	fmt.Println(parent_id)
	comment, err := h.AddComment(c, param.Body, post_id, uuid.MustParse(id), parent_id)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseCreated(c, "success created comment", comment)
}

func (h *PostHandler) GetCommentByIdHandler(c *gin.Context) {
	id := c.Param("id")
	new_id, e := uuid.Parse(id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	comment, err := h.GetCommentById(c, new_id)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseData(c, "Comment retrieved", nil, comment)
}

type commentPostParams struct {
	PostId   string `form:"post_id" binding:"required"`
	Paginate paginateForm
}

func (h *PostHandler) GetCommentPostHandler(c *gin.Context) {
	var param commentPostParams
	if err := c.ShouldBindQuery(&param); err != nil {
		util.BadRequest(c, err)
		return
	}
	new_post, e := uuid.Parse(param.PostId)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	limit, offset := newPageiante(&param.Paginate)
	comments, total, err := h.GetCommentPost(c, new_post, limit, offset)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	meta := util.WithMetadata(
		param.Paginate.Page,
		total,
		int64(math.Ceil(float64(total)/float64(param.Paginate.Page_size))),
		nil,
	)
	util.ResponseData(c, "success get post comment", &meta, comments)

}

type commentEdit struct {
	Body      string `json:"body"       binding:"required"`
	CommentId string `json:"comment_id" binding:"required"`
}

func (h *PostHandler) EditCommentHandler(c *gin.Context) {
	var param commentEdit
	if err := c.ShouldBindJSON(&param); err != nil {
		util.BadRequest(c, err)
		return
	}
	id := c.Value("cred").(middleware.Credentials).Id
	comment_id, e := uuid.Parse(param.CommentId)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	comment, err := h.EditComment(c, param.Body, comment_id, uuid.MustParse(id))
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseData(c, "success edited comment", nil, comment)
}

func (h *PostHandler) DeleteCommentHandler(c *gin.Context) {
	id := c.Param("id")
	acc_id := c.Value("cred").(middleware.Credentials).Id
	new_id, e := uuid.Parse(id)
	if e != nil {
		util.BadRequest(c, e)
		return
	}
	err := h.DeleteUserComment(c, new_id, uuid.MustParse(acc_id))
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseOk(c, "Comment deleted")
}
