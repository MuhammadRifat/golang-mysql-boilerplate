package response

import "time"

type UserResponse struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
}
