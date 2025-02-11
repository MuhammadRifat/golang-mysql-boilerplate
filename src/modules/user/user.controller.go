package user

import (
	"url-shortner/src/modules/user/request"
	"url-shortner/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserControllerStruct struct{}

var UserController = UserControllerStruct{}

// create one user
func (c *UserControllerStruct) CreateOne(ctx *gin.Context) {
	var userRequest request.UserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.Error(util.ValidationErr(err, &request.UserRequest{}))
		return
	}

	newUser, err := UserService.CreateUser(userRequest)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.AbortWithStatusJSON(util.ResponseCreated(newUser))
}

// get all user
func (c *UserControllerStruct) GetAll(ctx *gin.Context) {
	var queryRequest request.GetUserRequest
	if err := ctx.ShouldBindQuery(&queryRequest); err != nil {
		ctx.Error(err)
		return
	}

	Users, pagination, err := UserService.GetUserWithPaginate(queryRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.AbortWithStatusJSON(util.ResponseOK(Users, pagination))
}

// get one user
func (c *UserControllerStruct) GetOne(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(util.BadRequestErr("invalid id"))
		return
	}

	user, userErr := UserService.GetUserById(uint(id))
	if userErr != nil {
		ctx.Error(userErr)
		return
	}

	ctx.AbortWithStatusJSON(util.ResponseOK(user))
}

// update one user
func (c *UserControllerStruct) UpdateOne(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(util.BadRequestErr("invalid id"))
		return
	}

	var userUpdateRequest request.UserRequest
	if err := ctx.ShouldBindJSON(&userUpdateRequest); err != nil {
		ctx.Error(util.ValidationErr(err, &request.UserRequest{}))
		return
	}

	newUser, updateErr := UserService.UpdateUser(uint(id), userUpdateRequest)
	if updateErr != nil {
		ctx.Error(updateErr)
		return
	}
	ctx.AbortWithStatusJSON(util.ResponseOK(newUser))
}

// delete one user
func (c *UserControllerStruct) DeleteOne(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(util.BadRequestErr("invalid id"))
		return
	}
	userErr := UserService.DeleteUserById(uint(id))
	if userErr != nil {
		ctx.Error(userErr)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "User successfully deleted!",
	})
}
