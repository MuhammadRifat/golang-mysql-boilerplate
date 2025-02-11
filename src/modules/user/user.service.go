package user

import (
	"url-shortner/src/modules/user/model"
	"url-shortner/src/modules/user/request"
	"url-shortner/src/modules/user/response"
	"url-shortner/src/util"

	"github.com/jinzhu/copier"
)

type UserServiceStruct struct{}

var UserService = UserServiceStruct{}

// create item
func (s *UserServiceStruct) CreateUser(userRequest request.UserRequest) (model.User, error) {
	userModel := model.User{}
	copier.Copy(&userModel, userRequest)

	// store to db
	err := UserRepository.CreateOne(&userModel)
	if err != nil {
		return userModel, err
	}

	return userModel, nil
}

// get all items with pagination
func (s *UserServiceStruct) GetUserWithPaginate(query request.GetUserRequest) ([]response.UserResponse, util.Pagination, error) {
	paginateRequest := util.PaginateDefault(query.Page, query.Limit)

	items, err := UserRepository.GetAllWithPaginate(query, &paginateRequest)
	if err != nil {
		return items, util.Pagination{}, err
	}
	paginationResponse := util.PaginationMake(paginateRequest)

	return items, paginationResponse, nil
}

// get item by id
func (s *UserServiceStruct) GetUserById(id uint) (response.UserResponse, error) {
	item, err := UserRepository.GetById(id)
	if err != nil {
		return item, util.NotFoundErr("user not found")
	}

	return item, nil
}

// delete item by id
func (s *UserServiceStruct) DeleteUserById(id uint) error {
	oldItem, _ := UserRepository.GetById(id)
	if oldItem.ID == 0 {
		return util.NotFoundErr("user not found")
	}

	return UserRepository.DeleteOne(id)
}

// update item by id
func (s *UserServiceStruct) UpdateUser(id uint, userRequest request.UserRequest) (model.User, error) {
	userModel := model.User{}
	copier.Copy(&userModel, userRequest)

	oldItem, _ := UserRepository.GetById(id)
	if oldItem.ID == 0 {
		return model.User{}, util.NotFoundErr("user not found")
	}

	updatedItem, _ := UserRepository.UpdateOne(id, userModel)
	return updatedItem, nil
}
