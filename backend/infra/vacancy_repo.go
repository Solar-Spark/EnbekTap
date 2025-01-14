package infra

import (
	"enbektap/entities"

	"gorm.io/gorm"
)

type VacancyRepository interface {
	CreateVacancy(vacancy entities.Vacancy) error
	ReadVacancy(id uint) (entities.Vacancy, error)
	ReadVacancies(filterBy string, sortBy string, page int, pageSize int) ([]entities.Vacancy, int64, error)
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

func (repo *VacancyRepo) ReadVacancies(filterBy string, sortBy string, page int, pageSize int) ([]entities.Vacancy, int64, error) {
	var vacancies []entities.Vacancy
	var total int64

	// Start building the query
	query := repo.DB.Model(&entities.Vacancy{})

	// Apply filtering if jobType is not "none"
	if filterBy != "none" {
		query = query.Where("jobtype = ?", filterBy)
	}

	// Get total count before pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting if sortBy is not "none"
	if sortBy != "none" {
		switch sortBy {
		case "salary-asc":
			query = query.Order("salary ASC")
		case "salary-desc":
			query = query.Order("salary DESC")
		}
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&vacancies).Error
	if err != nil {
		return nil, 0, err
	}

	return vacancies, total, nil
}

func (repo *VacancyRepo) UpdateVacancy(id uint, updated entities.Vacancy) error {
	return repo.DB.Model(&entities.Vacancy{}).Where("vacancy_id = ?", id).Updates(updated).Error
}

func (repo *VacancyRepo) DeleteVacancy(id uint) error {
	return repo.DB.Delete(&entities.Vacancy{}, id).Error
}
