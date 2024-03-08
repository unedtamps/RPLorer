package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/middleware"
	"github.com/unedtamps/go-backend/src/service"
	"github.com/unedtamps/go-backend/util"
)

type AccountHandler struct {
	service.AccountServiceI
}

type AccountHandlerI interface {
	// CreateUser(c *gin.Context)
	// GetAllUser(c *gin.Context)
	// GetMe(c *gin.Context)
	// GetUserByEmail(c *gin.Context)
	// LoginUser(c *gin.Context)
	RegisterUserAccount(c *gin.Context)
	ConfirmRegistrant(c *gin.Context)
	ResendEmailConfirm(c *gin.Context)
	LoginHandler(c *gin.Context)
}

func newAccountHandler(accService service.AccountServiceI) AccountHandlerI {
	return &AccountHandler{accService}
}

type createUserParams struct {
	FirstName string `json:"first_name" binding:"required,max=255"`
	LastName  string `json:"last_name"  binding:"required,max=255"`
	Username  string `json:"user_name"  binding:"required,min=8,max=32"`
	Email     string `json:"email"      binding:"required,email"`
	Password  string `json:"password"   binding:"required,min=8,max=255"`
}

type loginUserParams struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
type emailUri struct {
	Email string `uri:"email" binding:"required,email"`
}

func (h *AccountHandler) RegisterUserAccount(c *gin.Context) {
	var params createUserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		util.BadRequest(c, err)
		return
	}
	account, err := h.CreateUserService(
		c,
		params.FirstName,
		params.LastName,
		params.Username,
		params.Email,
		params.Password,
	)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseCreated(c, "Account created", account)
}

func (h *AccountHandler) ResendEmailConfirm(c *gin.Context) {
	var email emailUri
	if err := c.ShouldBindUri(&email); err != nil {
		util.BadRequest(c, err)
		return
	}
	err := h.ReSendEmailConfirmation(c, email.Email)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseOk(c, "email confirmation sent")
}

func (h *AccountHandler) ConfirmRegistrant(c *gin.Context) {
	acc := c.Value("cred").(middleware.Credentials)
	err := h.ActivatedAccount(c, acc.Id)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseOk(c, "account activated")
}

func (h *AccountHandler) LoginHandler(c *gin.Context) {
	var login loginUserParams
	if err := c.ShouldBindJSON(&login); err != nil {
		util.BadRequest(c, err)
		return
	}
	token, err := h.LoginService(c, login.Email, login.Password)
	if err != nil {
		util.UnknownError(c, err.Error, err.Code)
		return
	}
	util.ResponseData(c, "success login", nil, TokenJwt{Token: *token})
}

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
