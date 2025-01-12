package infra

import (
	"enbektap/entities"

	"gorm.io/gorm"
)

type VacancyRepository interface {
	CreateVacancy(vacancy entities.Vacancy) error
	ReadVacancy(id uint) (entities.Vacancy, error)
	ReadVacancies() ([]entities.Vacancy, error)
	UpdateVacancy(id uint, updated entities.Vacancy) error
	DeleteVacancy(id uint) error
}

type VacancyRepo struct {
	DB *gorm.DB
}

func (repo *VacancyRepo) CreateVacancy(vacancy entities.Vacancy) error {
	return repo.DB.Create(&vacancy).Error
}

func (repo *VacancyRepo) ReadVacancy(id uint) (entities.Vacancy, error) {
	var vacancy entities.Vacancy
	err := repo.DB.First(&vacancy, id).Error
	return vacancy, err
}

func (repo *VacancyRepo) ReadVacancies() ([]entities.Vacancy, error) {
	var vacancies []entities.Vacancy
	err := repo.DB.Find(&vacancies).Error
	return vacancies, err
}

func (repo *VacancyRepo) UpdateVacancy(id uint, updated entities.Vacancy) error {
	return repo.DB.Model(&entities.Vacancy{}).Where("id = ?", id).Updates(updated).Error
}

func (repo *VacancyRepo) DeleteVacancy(id uint) error {
	return repo.DB.Delete(&entities.Vacancy{}, id).Error
}
