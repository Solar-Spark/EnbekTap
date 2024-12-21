package models

import (
	"errors"

	"gorm.io/gorm"
)

type Vacancy struct {
	VacancyID   int64  `gorm:"primaryKey;autoIncrement"`
	Vacancy     string `gorm:"type:varchar(255);column:vacancy"`
	Salary      int    `gorm:"column:salary"`
	JobType     string `gorm:"type:enum('full-time', 'part-time');column:jobtype"`
	Description string `gorm:"type:varchar(255);column:description"`
}

func CreateVacancy(db *gorm.DB, vacancy, jobType, description string, salary int) error {
	vacancyRecord := Vacancy{
		Vacancy:     vacancy,
		Salary:      salary,
		JobType:     jobType,
		Description: description,
	}
	result := db.Create(&vacancyRecord)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateVacancy(db *gorm.DB, vacancyID int64, newVacancy, newJobType, newDescription string, newSalary int) error {
	var vacancyRecord Vacancy
	result := db.First(&vacancyRecord, vacancyID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("vacancy with this ID not found")
	}

	if newVacancy != "" {
		vacancyRecord.Vacancy = newVacancy
	}
	if newJobType != "" {
		vacancyRecord.JobType = newJobType
	}
	if newDescription != "" {
		vacancyRecord.Description = newDescription
	}
	if newSalary != 0 {
		vacancyRecord.Salary = newSalary
	}

	result = db.Save(&vacancyRecord)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ReadVacancies(db *gorm.DB) ([]Vacancy, error) {
	var vacancies []Vacancy
	result := db.Find(&vacancies)
	if result.Error != nil {
		return nil, result.Error
	}
	return vacancies, nil
}

func ReadOneVacancy(db *gorm.DB, id int64) (*Vacancy, error) {
	var vacancy Vacancy
	result := db.First(&vacancy, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("vacancy not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &vacancy, nil
}

func DeleteVacancy(db *gorm.DB, vacancyID int64) error {
	result := db.Delete(&Vacancy{}, vacancyID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
