package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/middleware"
	"github.com/unedtamps/go-backend/src/service"
	"github.com/unedtamps/go-backend/util"
)

type UserHandler struct {
	u service.UserServiceI
}

type UserHandlerI interface {
	CreateUser(c *gin.Context)
	GetAllUser(c *gin.Context)
	GetMe(c *gin.Context)
	LoginUser(c *gin.Context)
}

func newUserHandler(userService service.UserServiceI) UserHandlerI {
	return &UserHandler{userService}
}

type createUserParams struct {
	Name     string `json:"name"     binding:"required,min=8,max=255"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

type loginUserParams struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var params createUserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorHandler(err))
		return
	}
	user, err := h.u.CreateUser(c, params.Name, params.Email, params.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ErrorHandler(err))
		return
	}
	fmt.Println("masuk ke sini")
	util.ResponseCreated(c, "User created", user)
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	var params paginateForm
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorHandler(err))
		return
	}
	users, meta, err := h.u.GetAllUser(c, params.Page, params.Page_size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ErrorHandler(err))
		return
	}
	util.ResponseData(c, "Get all user", &meta, users)
}

func (h *UserHandler) GetMe(c *gin.Context) {
	data := middleware.GetCredentials(c)
	util.ResponseData(c, "Get me", nil, data)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	body := loginUserParams{}
	if err := c.ShouldBindJSON(&body); err != nil {
		util.BadRequest(c, err)
		return
	}
	user, err := h.u.LoginUser(c, body.Email, body.Password)
	if err != nil {
		util.UnknownError(c, err)
		return
	}
	token, err := middleware.CreateJwtToken(middleware.Credentials{
		Id:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	})
	if err != nil {
		util.UnknownError(c, err)
		return
	}
	util.ResponseData(c, "Login Success", nil, user, TokenJwt{Token: token})
}
