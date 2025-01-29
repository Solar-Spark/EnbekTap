package services

import (
	"enbektap/entities"
	"enbektap/infra"
)

type UserService struct {
	Repo infra.UserRepository
}

func (s *UserService) CreateUser(user entities.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUser(id uint) (entities.User, error) {
	return s.Repo.ReadUser(id)
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
	return s.Repo.ReadUsers()
}

func (s *UserService) UpdateUser(id uint, user entities.User) error {
	return s.Repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) GetUserByEmail(email string) (entities.User, error) {
	return s.Repo.ReadUserByEmail(email)
}
