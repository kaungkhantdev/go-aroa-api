package requests

type AuthLoginRequest struct {
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=6"`
}

type AuthRegisterRequest struct {
	Name		string	`json:"name" validate:"required,min=3"`
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=6"`
	Avatar		string	`json:"avatar" validate:"required,min=6"`
}