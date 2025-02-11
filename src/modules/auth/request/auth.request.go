package request

type LoginRequest struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=8"`
}

type RegisterRequest struct {
	Name     string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=8"`
}
