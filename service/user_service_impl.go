package service

import (
	"context"
	"log"
	"playlist-saver/app/middleware"
	"playlist-saver/exceptions"
	"playlist-saver/model/domain"
	"playlist-saver/model/web"
	"playlist-saver/repository"
	"playlist-saver/utility"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	jwtAuth        *middleware.ConfigJWT
}

func NewUserService(UserRepository repository.UserRepository, jwtAuth *middleware.ConfigJWT) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		jwtAuth:        jwtAuth,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) web.UserRegisterResponse {
	err := request.Validate()
	exceptions.PanicIfError(err)

	userInsert := domain.User{}
	userInsert.UserInsertRequest(request)
	password, err := utility.HashPassword(userInsert.Password)

	exceptions.PanicIfError(err)
	userInsert.Password = password
	userInsert = service.UserRepository.Register(ctx, userInsert)

	response := userInsert.ToRegisterResponse()

	return response
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) string {
	err := request.Validate()
	exceptions.PanicIfError(err)

	userResponse := service.UserRepository.Login(ctx, request.Email)
	log.Println(userResponse.Password)
	log.Println(request.Password)
	//requestPassword, _ := utility.HashPassword(request.Password)
	match := utility.CheckPasswordHash(request.Password, userResponse.Password)
	log.Println("INI MATCH", match)
	if match {
		jwt := service.jwtAuth.GenerateToken(userResponse.Id, userResponse.Name, userResponse.Role)
		log.Print(jwt)
		exceptions.PanicIfError(err)
		return jwt
	} else {
		panic("Password not match Or not Found Any Data In the Server")
		return ""
	}
}
