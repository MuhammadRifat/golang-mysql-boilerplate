#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: ./generate_module.sh <module_name>"
  exit 1
fi

MODULE_NAME=$1
DIRECTORY="url-shortner/src"
MODULE_PATH="src/modules/$MODULE_NAME"
REQUEST_PATH="src/modules/$MODULE_NAME/request"
RESPONSE_PATH="src/modules/$MODULE_NAME/response"
MODEL_PATH="src/modules/$MODULE_NAME/model"

mkdir -p $MODULE_PATH
mkdir -p $REQUEST_PATH
mkdir -p $RESPONSE_PATH
mkdir -p $MODEL_PATH

touch "$REQUEST_PATH/$MODULE_NAME.request.go"
touch "$RESPONSE_PATH/$MODULE_NAME.response.go"
touch "$MODEL_PATH/$MODULE_NAME.model.go"

# Capitalized module name
CAP_MODULE_NAME=$(echo "$MODULE_NAME" | awk '{print toupper(substr($0,1,1)) tolower(substr($0,2))}')


# Create Router
cat <<EOF > $MODULE_PATH/${MODULE_NAME}.route.go
package $MODULE_NAME

import "github.com/gin-gonic/gin"

func RoutesHandler(router *gin.RouterGroup) {
	group := router.Group("/$MODULE_NAME")
	{
		group.POST("/", ${CAP_MODULE_NAME}Controller.CreateOne)
		group.GET("/", ${CAP_MODULE_NAME}Controller.GetAll)
		group.GET("/:id", ${CAP_MODULE_NAME}Controller.GetOne)
		group.PUT("/:id", ${CAP_MODULE_NAME}Controller.UpdateOne)
		group.DELETE("/:id", ${CAP_MODULE_NAME}Controller.DeleteOne)
	}
}
EOF

# Create Controller
cat <<EOF > $MODULE_PATH/${MODULE_NAME}.controller.go
package $MODULE_NAME

import (
	"$DIRECTORY/modules/$MODULE_NAME/request"
	"$DIRECTORY/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ${CAP_MODULE_NAME}ControllerStruct struct{}

var ${CAP_MODULE_NAME}Controller = ${CAP_MODULE_NAME}ControllerStruct{}

// create one $MODULE_NAME
func (c *${CAP_MODULE_NAME}ControllerStruct) CreateOne(ctx *gin.Context) {
	var ${MODULE_NAME}Request request.${CAP_MODULE_NAME}Request
	if err := ctx.ShouldBindJSON(&${MODULE_NAME}Request); err != nil {
		ctx.Error(util.ValidationErr(err, &request.${CAP_MODULE_NAME}Request{}))
		return
	}

	new${CAP_MODULE_NAME}, err := ${CAP_MODULE_NAME}Service.Create${CAP_MODULE_NAME}(${MODULE_NAME}Request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.AbortWithStatusJSON(util.ResponseCreated(new${CAP_MODULE_NAME}))
}

// get all $MODULE_NAME
func (c *${CAP_MODULE_NAME}ControllerStruct) GetAll(ctx *gin.Context) {
	var queryRequest request.Get${CAP_MODULE_NAME}Request
	if err := ctx.ShouldBindQuery(&queryRequest); err != nil {
		ctx.Error(err)
		return
	}

	${CAP_MODULE_NAME}s, pagination, err := ${CAP_MODULE_NAME}Service.Get${CAP_MODULE_NAME}WithPaginate(queryRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.AbortWithStatusJSON(util.ResponseOK(${CAP_MODULE_NAME}s, pagination))
}

// get one $MODULE_NAME
func (c *${CAP_MODULE_NAME}ControllerStruct) GetOne(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(util.BadRequestErr("invalid id"))
		return
	}

	${MODULE_NAME}, ${MODULE_NAME}Err := ${CAP_MODULE_NAME}Service.Get${CAP_MODULE_NAME}ById(uint(id))
	if ${MODULE_NAME}Err != nil {
		ctx.Error(${MODULE_NAME}Err)
		return
	}

	ctx.AbortWithStatusJSON(util.ResponseOK(${MODULE_NAME}))
}

// update one $MODULE_NAME
func (c *${CAP_MODULE_NAME}ControllerStruct) UpdateOne(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(util.BadRequestErr("invalid id"))
		return
	}

	var ${MODULE_NAME}UpdateRequest request.${CAP_MODULE_NAME}Request
	if err := ctx.ShouldBindJSON(&${MODULE_NAME}UpdateRequest); err != nil {
		ctx.Error(util.ValidationErr(err, &request.${CAP_MODULE_NAME}Request{}))
		return
	}

	new${CAP_MODULE_NAME}, updateErr := ${CAP_MODULE_NAME}Service.Update${CAP_MODULE_NAME}(uint(id), ${MODULE_NAME}UpdateRequest)
	if updateErr != nil {
		ctx.Error(updateErr)
		return
	}
	ctx.AbortWithStatusJSON(util.ResponseOK(new${CAP_MODULE_NAME}))
}

// delete one $MODULE_NAME
func (c *${CAP_MODULE_NAME}ControllerStruct) DeleteOne(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(util.BadRequestErr("invalid id"))
		return
	}
	${MODULE_NAME}Err := ${CAP_MODULE_NAME}Service.Delete${CAP_MODULE_NAME}ById(uint(id))
	if ${MODULE_NAME}Err != nil {
		ctx.Error(${MODULE_NAME}Err)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "${CAP_MODULE_NAME} successfully deleted!",
	})
}
EOF

# Create Service
cat <<EOF > $MODULE_PATH/${MODULE_NAME}.service.go
package $MODULE_NAME

import (
	"$DIRECTORY/modules/$MODULE_NAME/model"
	"$DIRECTORY/modules/$MODULE_NAME/request"
	"$DIRECTORY/modules/$MODULE_NAME/response"
	"$DIRECTORY/util"

	"github.com/jinzhu/copier"
)

type ${CAP_MODULE_NAME}ServiceStruct struct{}

var ${CAP_MODULE_NAME}Service = ${CAP_MODULE_NAME}ServiceStruct{}

// create item
func (s *${CAP_MODULE_NAME}ServiceStruct) Create${CAP_MODULE_NAME}(${MODULE_NAME}Request request.${CAP_MODULE_NAME}Request) (model.${CAP_MODULE_NAME}, error) {
	${MODULE_NAME}Model := model.${CAP_MODULE_NAME}{}
	copier.Copy(&${MODULE_NAME}Model, ${MODULE_NAME}Request)

	// store to db
	err := ${CAP_MODULE_NAME}Repository.CreateOne(&${MODULE_NAME}Model)
	if err != nil {
		return ${MODULE_NAME}Model, err
	}

	return ${MODULE_NAME}Model, nil
}

// get all items with pagination
func (s *${CAP_MODULE_NAME}ServiceStruct) Get${CAP_MODULE_NAME}WithPaginate(query request.Get${CAP_MODULE_NAME}Request) ([]response.${CAP_MODULE_NAME}Response, util.Pagination, error) {
	paginateRequest := util.PaginateDefault(query.Page, query.Limit)

	items, err := ${CAP_MODULE_NAME}Repository.GetAllWithPaginate(query, &paginateRequest)
	if err != nil {
		return items, util.Pagination{}, err
	}
	paginationResponse := util.PaginationMake(paginateRequest)

	return items, paginationResponse, nil
}

// get item by id
func (s *${CAP_MODULE_NAME}ServiceStruct) Get${CAP_MODULE_NAME}ById(id uint) (response.${CAP_MODULE_NAME}Response, error) {
	item, err := ${CAP_MODULE_NAME}Repository.GetById(id)
	if err != nil {
		return item, util.NotFoundErr("${MODULE_NAME} not found")
	}

	return item, nil
}

// delete item by id
func (s *${CAP_MODULE_NAME}ServiceStruct) Delete${CAP_MODULE_NAME}ById(id uint) error {
	oldItem, _ := ${CAP_MODULE_NAME}Repository.GetById(id)
	if oldItem.ID == 0 {
		return util.NotFoundErr("${MODULE_NAME} not found")
	}

	return ${CAP_MODULE_NAME}Repository.DeleteOne(id)
}

// update item by id
func (s *${CAP_MODULE_NAME}ServiceStruct) Update${CAP_MODULE_NAME}(id uint, ${MODULE_NAME}Request request.${CAP_MODULE_NAME}Request) (model.${CAP_MODULE_NAME}, error) {
	${MODULE_NAME}Model := model.${CAP_MODULE_NAME}{}
	copier.Copy(&${MODULE_NAME}Model, ${MODULE_NAME}Request)

	oldItem, _ := ${CAP_MODULE_NAME}Repository.GetById(id)
	if oldItem.ID == 0 {
		return model.${CAP_MODULE_NAME}{}, util.NotFoundErr("${MODULE_NAME} not found")
	}

	updatedItem, _ := ${CAP_MODULE_NAME}Repository.UpdateOne(id, ${MODULE_NAME}Model)
	return updatedItem, nil
}
EOF

# Create Request
cat <<EOF > $REQUEST_PATH/${MODULE_NAME}.request.go
package request

type ${CAP_MODULE_NAME}Request struct {
    Title string \`binding:"required,min=3,max=100"\`
}

type Get${CAP_MODULE_NAME}Request struct {
    Title string

	Limit string
	Page  string
}
EOF

# Create Response
cat <<EOF > $RESPONSE_PATH/${MODULE_NAME}.response.go
package response

import "time"

type ${CAP_MODULE_NAME}Response struct {
    ID        uint
	Title     string
	CreatedAt time.Time
}
EOF

# Create Model
cat <<EOF > $MODEL_PATH/${MODULE_NAME}.model.go
package model

import (
	"gorm.io/gorm"
)

type $CAP_MODULE_NAME struct {
	gorm.Model
	Title string \`gorm:"size:100;not null"\`
}
EOF

# Create Repository
cat <<EOF > $MODULE_PATH/${MODULE_NAME}.repository.go
package $MODULE_NAME

import (
	"$DIRECTORY/config"
	"$DIRECTORY/modules/$MODULE_NAME/model"
	"$DIRECTORY/modules/$MODULE_NAME/request"
	"$DIRECTORY/modules/$MODULE_NAME/response"
	"$DIRECTORY/util"
)

type ${CAP_MODULE_NAME}RepositoryStruct struct{}

var ${CAP_MODULE_NAME}Repository = ${CAP_MODULE_NAME}RepositoryStruct{}

func (r *${CAP_MODULE_NAME}RepositoryStruct) CreateOne(${MODULE_NAME} *model.${CAP_MODULE_NAME}) error {
	return config.DB.Create(&${MODULE_NAME}).Error
}

func (r *${CAP_MODULE_NAME}RepositoryStruct) GetAllWithPaginate(query request.Get${CAP_MODULE_NAME}Request, paginate *util.Paginate) ([]response.${CAP_MODULE_NAME}Response, error) {
	var items []response.${CAP_MODULE_NAME}Response

	session := config.DB.Model(&model.${CAP_MODULE_NAME}{})
	if query.Title != "" {
		session = session.Where("title LIKE ?", "%"+query.Title+"%")
	}

	// Count total items before applying LIMIT and OFFSET
	var totalCount int64
	countErr := session.Count(&totalCount).Error
	if countErr != nil {
		return items, countErr
	}
	paginate.Count = totalCount

	// Fetch paginated items
	err := session.Limit(paginate.Limit).Offset(paginate.Offset).Order("id DESC").Find(&items).Error
	if err != nil {
		return items, err
	}

	return items, nil
}

func (r *${CAP_MODULE_NAME}RepositoryStruct) GetById(id uint) (response.${CAP_MODULE_NAME}Response, error) {
	var item = response.${CAP_MODULE_NAME}Response{}

	err := config.DB.Model(model.${CAP_MODULE_NAME}{}).First(&item, id).Error
	return item, err
}

func (r *${CAP_MODULE_NAME}RepositoryStruct) UpdateOne(id uint, updateBody model.${CAP_MODULE_NAME}) (model.${CAP_MODULE_NAME}, error) {
	var itemModel model.${CAP_MODULE_NAME}
	if err := config.DB.First(&itemModel, id).Error; err != nil {
		return itemModel, err
	}

	if updateBody.Title != "" {
		itemModel.Title = updateBody.Title
	}
	err := config.DB.Save(&itemModel).Error
	return itemModel, err
}

func (r *${CAP_MODULE_NAME}RepositoryStruct) DeleteOne(id uint) error {
	return config.DB.Model(model.${CAP_MODULE_NAME}{}).Delete("id = ?", id).Error
}
EOF

echo "Module '$MODULE_NAME' created successfully."
