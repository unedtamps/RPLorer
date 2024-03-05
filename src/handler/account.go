package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/service"
	"github.com/unedtamps/go-backend/util"
)

type AccountHandler struct {
	acc service.AccountServiceI
}

type AccountHandlerI interface {
	// CreateUser(c *gin.Context)
	// GetAllUser(c *gin.Context)
	// GetMe(c *gin.Context)
	// GetUserByEmail(c *gin.Context)
	// LoginUser(c *gin.Context)
	GetOne(c *gin.Context)
}

func newAccountHandler(accService service.AccountService) AccountHandlerI {
	return &AccountHandler{&accService}
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

func (h *AccountHandler) GetOne(c *gin.Context) {
	acc, err := h.acc.GetOne(c)
	if err != nil {
		util.UnknownError(c, err)
		return
	}
	util.ResponseData(c, "get one user", nil, acc)
}

// func (h *AccountHandler) CreateUser(c *gin.Context) {
// 	var params createUserParams
// 	if err := c.ShouldBindJSON(&params); err != nil {
// 		c.JSON(http.StatusBadRequest, util.ErrorHandler(err))
// 		return
// 	}
// 	user, err := h.CreateUser(c, params.Name, params.Email, params.Password)

// 	// send email confirmation
// 	body := make(chan string)
// 	go func() {
// 		body <- util.ParseAccountConfirmation(util.EmailConfirm{
// 			Id:    user.ID,
// 			Name:  user.Name,
// 			Email: user.Email,
// 		})
// 	}()
// 	go helper.NewEmail("Account Confirmation", user.Email, body).Send()

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, util.ErrorHandler(err))
// 		return
// 	}
// 	util.ResponseCreated(c, "User created", user)
// }

// func (h *UserHandler) GetAllUser(c *gin.Context) {
// 	var params paginateForm
// 	if err := c.ShouldBindQuery(&params); err != nil {
// 		c.JSON(http.StatusBadRequest, util.ErrorHandler(err))
// 		return
// 	}
// 	users, meta, err := h.u.GetAllUser(c, params.Page, params.Page_size)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, util.ErrorHandler(err))
// 		return
// 	}
// 	util.ResponseData(c, "Get all user", &meta, users)
// }

// func (h *UserHandler) GetMe(c *gin.Context) {
// 	data := middleware.GetCredentials(c)
// 	util.ResponseData(c, "Get me", nil, data)
// }

// func (h *UserHandler) LoginUser(c *gin.Context) {
// 	body := loginUserParams{}
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		util.BadRequest(c, err)
// 		return
// 	}
// 	user, err := h.u.LoginUser(c, body.Email, body.Password)
// 	if err != nil {
// 		util.UnknownError(c, err)
// 		return
// 	}
// 	token, err := middleware.CreateJwtToken(middleware.Credentials{
// 		Id:    user.ID,
// 		Email: user.Email,
// 		Name:  user.Name,
// 	})
// 	if err != nil {
// 		util.UnknownError(c, err)
// 		return
// 	}
// 	util.ResponseData(c, "Login Success", nil, user, TokenJwt{Token: token})
// }

// func (h *UserHandler) GetUserByEmail(c *gin.Context) {
// 	email := c.Param("email")
// 	user, err := h.u.GetUserByEmail(c, email)
// 	if err != nil {
// 		util.NotFoundError(c, err)
// 		return
// 	}
// 	util.ResponseData(c, "Get user by email", nil, user)
// }