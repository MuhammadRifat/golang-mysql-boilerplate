package auth

import (
	"url-shortner/src/modules/auth/request"
	"url-shortner/src/util"

	"github.com/gin-gonic/gin"
)

type AuthControllerStruct struct{}

var AuthController = AuthControllerStruct{}

func (c *AuthControllerStruct) LoginHandler(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.Error(util.ValidationErr(err, &request.LoginRequest{}))
		return
	}

	// Call the auth service to handle login
	token, err := AuthService.Login(loginRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := map[string]string{"AccessToken": token}
	ctx.AbortWithStatusJSON(util.ResponseOK(res))
}

func (c *AuthControllerStruct) RegisterHandler(ctx *gin.Context) {
	var registerRequest request.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.Error(util.ValidationErr(err, &request.RegisterRequest{}))
		return
	}

	// Call the auth service to handle registration
	user, err := AuthService.Register(registerRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.AbortWithStatusJSON(util.ResponseCreated(user))
}
