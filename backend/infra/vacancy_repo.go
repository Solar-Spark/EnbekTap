package infra

import (
	"enbektap/entities"

	"gorm.io/gorm"
)

type VacancyRepository interface {
	CreateVacancy(vacancy entities.Vacancy) error
	ReadVacancy(id uint) (entities.Vacancy, error)
	ReadVacancies(filterBy string, sortBy string) ([]entities.Vacancy, error)
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

func (repo *VacancyRepo) ReadVacancies(filterBy string, sortBy string) ([]entities.Vacancy, error) {
	var vacancies []entities.Vacancy

	// Start building the query
	query := repo.DB.Model(&entities.Vacancy{})

	// Apply filtering if jobType is not "none"
	if filterBy != "none" {
		query = query.Where("jobtype = ?", filterBy)
	}

	// Apply sorting if sortBy is not "none"
	if sortBy != "none" {
		switch sortBy {
		case "salary-asc":
			query = query.Order("salary ASC")
		case "salary-desc":
			query = query.Order("salary DESC")
		default:
			break
		}
	}

	// Fetch the results, limit to 9 records
	err := query.Limit(9).Find(&vacancies).Error
	if err != nil {
		return nil, err
	}

	return vacancies, nil
}

func (repo *VacancyRepo) UpdateVacancy(id uint, updated entities.Vacancy) error {
	return repo.DB.Model(&entities.Vacancy{}).Where("id = ?", id).Updates(updated).Error
}

func (repo *VacancyRepo) DeleteVacancy(id uint) error {
	return repo.DB.Delete(&entities.Vacancy{}, id).Error
}
