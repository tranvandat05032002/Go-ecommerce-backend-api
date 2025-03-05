package service

import "example.com/m/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserInfo() string {
	return us.userRepo.GetUserInfo()
}
