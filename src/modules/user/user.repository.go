package user

import (
	"url-shortner/src/config"
	"url-shortner/src/modules/user/model"
	"url-shortner/src/modules/user/request"
	"url-shortner/src/modules/user/response"
	"url-shortner/src/util"
)

type UserRepositoryStruct struct{}

var UserRepository = UserRepositoryStruct{}

func (r *UserRepositoryStruct) CreateOne(user *model.User) error {
	return config.DB.Create(&user).Error
}

func (r *UserRepositoryStruct) GetAllWithPaginate(query request.GetUserRequest, paginate *util.Paginate) ([]response.UserResponse, error) {
	var items []response.UserResponse

	session := config.DB.Model(&model.User{})
	if query.Email != "" {
		session = session.Where("email LIKE ?", "%"+query.Email+"%")
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

func (r *UserRepositoryStruct) GetById(id uint) (response.UserResponse, error) {
	var item = response.UserResponse{}

	err := config.DB.Model(model.User{}).First(&item, id).Error
	return item, err
}

func (r *UserRepositoryStruct) GetOneRecordByQuery(queryMap map[string]interface{}) (model.User, error) {
	var data model.User
	err := config.DB.Model(model.User{}).Where(queryMap).First(&data).Error

	return data, err
}

func (r *UserRepositoryStruct) UpdateOne(id uint, updateBody model.User) (model.User, error) {
	var itemModel model.User
	if err := config.DB.First(&itemModel, id).Error; err != nil {
		return itemModel, err
	}

	if updateBody.Email != "" {
		itemModel.Email = updateBody.Email
	}
	err := config.DB.Save(&itemModel).Error
	return itemModel, err
}

func (r *UserRepositoryStruct) DeleteOne(id uint) error {
	return config.DB.Model(model.User{}).Delete("id = ?", id).Error
}
