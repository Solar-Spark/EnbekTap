package infra

import (
	"enbektap/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entities.User) error
	ReadUser(id uint) (entities.User, error)
	ReadUsers(filterBy string, sortBy string, page int, pageSize int) ([]entities.User, int64, error)
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

func (repo *UserRepo) ReadUsers(filterBy string, sortBy string, page int, pageSize int) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64
	if filterBy != "" {
		repo.DB = repo.DB.Where(filterBy)
	}
	repo.DB.Find(&users).Count(&total)
	err := repo.DB.Limit(pageSize).Offset((page - 1) * pageSize).Order(sortBy).Find(&users).Error
	return users, total, err
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
