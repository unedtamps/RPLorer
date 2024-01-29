package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/service"
	"github.com/unedtamps/go-backend/util"
)

type UserHandler struct {
	u service.UserServiceI
}

type UserHandlerI interface {
	CreateUser(c *gin.Context)
}

func newUserHandler(userService service.UserServiceI) UserHandler {
	return UserHandler{userService}
}

type createUserParams struct {
	Name     string `json:"name"     binding:"required,min=8,max=255"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
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
	c.JSON(http.StatusCreated, util.ResponseHandler(user))
}
