package domain

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type GetUserRequest struct {
	ID uint `json:"id"`
}

type GetUserResponse struct {
	ID    uint `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
