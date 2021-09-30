package servuser

import (
	"context"
	"errors"
	"log"
	"playlist-saver/app/middleware"
	"playlist-saver/repository/repouser"
	"playlist-saver/utility"
)

type UserServiceImpl struct {
	UserRepository repouser.UserRepository
	jwtAuth        *middleware.ConfigJWT
}

func NewUserService(UserRepository repouser.UserRepository, jwtAuth *middleware.ConfigJWT) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		jwtAuth:        jwtAuth,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, dataUser User) (User, error) {
	err := dataUser.Validate()
	if err != nil {
		return dataUser, err
	}
	password, err := utility.HashPassword(dataUser.Password)
	if err != nil {
		return dataUser, err
	}

	dataUser.Password = password
	dataUser.Status.Name = "FREE"

	userRecord := dataUser.ToRecordUser()
	log.Print("Test Recored", userRecord)
	insertedUser, err := service.UserRepository.Register(ctx, userRecord)
	if err != nil {
		return dataUser, err
	}
	dataUser.FromRecordUser(insertedUser)

	return dataUser, nil
}

func (service *UserServiceImpl) Login(ctx context.Context, email, password string) (string, error) {
	if len(email) == 0 || len(password) == 0 {
		return "", errors.New("dont't leave a blank form")
	}

	userRecord, err := service.UserRepository.Login(ctx, email)
	if err != nil {
		return "", err
	}
	match := utility.CheckPasswordHash(password, userRecord.Password)

	if !match {
		return "", errors.New("password doesn't match")
	}

	jwt := service.jwtAuth.GenerateToken(userRecord.Id, userRecord.Name, userRecord.Role, userRecord.Status.Name)

	return jwt, nil
}

func (service *UserServiceImpl) GetAllUser(ctx context.Context, admin string) ([]User, error) {
	log.Println("Test admin", admin)

	if admin != "ADMIN" {
		return nil, errors.New("you're not admin, can't fetch all data")
	}

	dataFinal := make([]User, 0)
	UserResult, err := service.UserRepository.GetAllUser(ctx)
	if err != nil {
		return nil, errors.New("not Found")
	}
	for _, values := range UserResult {
		dataUser := User{}
		dataUser.Name = values.Name
		dataUser.Email = values.Email
		dataUser.Birthday = values.Birthday
		dataUser.Gender = values.Gender
		dataUser.Role = values.Role
		dataUser.CreatedAt = values.CreatedAt
		dataUser.UpdatedAt = values.UpdatedAt
		dataUser.Status.Id = values.Status.Id
		dataUser.Status.Name = values.Status.Name
		dataUser.Status.CreatedAt = values.Status.CreatedAt
		dataUser.Status.UpdatedAt = values.Status.UpdatedAt
		dataFinal = append(dataFinal, dataUser)
	}
	return dataFinal, nil
}

func (service *UserServiceImpl) GetUserById(ctx context.Context, id int) (User, error) {
	result, err := service.UserRepository.UserFindById(ctx, id)
	finalResult := User{
		Name: result.Name,
		Email: result.Email,
		Gender: result.Gender,
		Role: result.Role,
	}
	if err != nil {
		return User{}, err
	}
	return finalResult,nil
}

func (service *UserServiceImpl) UpdateUser(ctx context.Context, user User, id int) error {
	getUserById, err := service.UserRepository.UserFindById(ctx, id)
	if err != nil {
		return err
	}

	match := utility.CheckPasswordHash(user.Password, getUserById.Password)
	if !match {
		password, err := utility.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = password
	}
	toRecord := user.ToRecordUser()
	err = service.UserRepository.UpdateUser(ctx, toRecord, id)
	if err != nil {
		return err
	}
	return nil
}
