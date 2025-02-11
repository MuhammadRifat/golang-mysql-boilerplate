package auth

import (
	"url-shortner/src/config"
	"url-shortner/src/modules/auth/request"
	"url-shortner/src/modules/user"
	"url-shortner/src/modules/user/model"
	"url-shortner/src/util"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceStruct struct{}

var AuthService = AuthServiceStruct{}

func (s *AuthServiceStruct) Login(request request.LoginRequest) (string, error) {
	user, err := user.UserRepository.GetOneRecordByQuery(map[string]interface{}{"email": request.Email})
	if err != nil {
		return "", util.UnauthorizedErr("invalid email or password")
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", util.UnauthorizedErr("invalid email or password")
	}

	token, err := util.GenerateJWT(int(user.ID))
	if err != nil {
		return "", util.InternalServerErr("error generating jwt")
	}

	return token, nil
}

func (s *AuthServiceStruct) Register(request request.RegisterRequest) (model.User, error) {
	queryMap := map[string]interface{}{"email": request.Email}
	isExist, _ := user.UserRepository.GetOneRecordByQuery(queryMap)
	if isExist.ID != 0 {
		return model.User{}, util.BadRequestErr("email already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	// Access the global DB instance directly
	if err := config.DB.Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
