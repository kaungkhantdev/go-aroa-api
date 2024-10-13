package services

import (
	"errors"
	"fmt"
	"go-aora-api/internal/repository"
	"go-aora-api/pkg/hash"
	"go-aora-api/pkg/jwt"
)

type AuthService struct {
	userService	*UserService
}

type Payload struct {
	UserId	int `json:"user_id"`
}

type LoginData struct {
	Email		string `json:"email"`
	Password	string `json:"password"`
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{userService: userService}
}

func (s *AuthService) Register(data repository.CreateData) (string, error) {

	checkEmail := s.userService.CheckEmail(data.Email);
	if checkEmail {
		return "", errors.New("email already exists");
	}

	user, err := s.userService.CreateUserService(data);
	if err != nil {
		return "", errors.New("error: something wrong when create user")
	}

	payload := Payload{
		UserId: int(user.ID),
	}
	token, err := responseJwt(payload);


	if err != nil {
		return "", errors.New("error: something wrong")
	}

	return token, nil;
}

func (s *AuthService) Login(data LoginData) (string, error) {
	hasUser := s.userService.CheckEmail(data.Email);
	if !hasUser {
		return "", errors.New("error: invalid password or email")
	}

	user, err := s.userService.FindByEmail(data.Email);

	if err != nil {
		return "", errors.New("error: something wrong")
	}

	/** check password */
	if hash.CheckPasswordHash(user.Password, data.Password) {
		fmt.Println("password hash", hash.CheckPasswordHash(user.Password, data.Password))
		return "", errors.New("error: invalid password or email",)
	}

	payload := Payload{
		UserId: int(user.ID),
	}
	token, err := responseJwt(payload);


	if err != nil {
		return "", errors.New("error: something wrong")
	}

	return token, nil;
}

func responseJwt(payload  Payload) (string, error) {

	token, err := jwt.GenerateJWT(payload.UserId);

	if err != nil {
		return "", errors.New("error: when create jwt")
	}

	return token, nil
}