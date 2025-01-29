package infra

import (
	"enbektap/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entities.User) error
	ReadUser(id uint) (entities.User, error)
	ReadUsers() ([]entities.User, error)
	ReadUserByEmail(email string) (entities.User, error)
	UpdateUser(id uint, updated entities.User) error
	DeleteUser(id uint) error
}

type UserRepo struct {
	DB *gorm.DB
}

func (repo *UserRepo) CreateUser(user entities.User) error {
	return repo.DB.Create(&user).Error
}

func (repo *UserRepo) ReadUser(id uint) (entities.User, error) {
	var user entities.User
	err := repo.DB.First(&user, id).Error
	return user, err
}

func (repo *UserRepo) ReadUsers() ([]entities.User, error) {
	var user []entities.User

	query := repo.DB.Model(&entities.User{})
	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepo) ReadUserByEmail(email string) (entities.User, error) {
	var user entities.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (repo *UserRepo) UpdateUser(id uint, updated entities.User) error {
	return repo.DB.Model(&entities.User{}).Where("user_id = ?", id).Updates(updated).Error
}

func (repo *UserRepo) DeleteUser(id uint) error {
	return repo.DB.Delete(&entities.User{}, id).Error
}
