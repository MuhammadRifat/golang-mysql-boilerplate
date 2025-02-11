package request

type UserRequest struct {
	Name     string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=8"`
}

type GetUserRequest struct {
	Email string

	Limit string
	Page  string
}
