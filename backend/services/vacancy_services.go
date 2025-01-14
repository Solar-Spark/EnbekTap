package services

import (
	"enbektap/entities"
	"enbektap/infra"
)

type VacancyService struct {
	Repo infra.VacancyRepository
}

func (s *VacancyService) CreateVacancy(vacancy entities.Vacancy) error {
	return s.Repo.CreateVacancy(vacancy)
}

func (s *VacancyService) GetVacancy(id uint) (entities.Vacancy, error) {
	return s.Repo.ReadVacancy(id)
}

func (s *VacancyService) GetAllVacancies(filterBy string, sortBy string) ([]entities.Vacancy, error) {
	return s.Repo.ReadVacancies(filterBy, sortBy)
}

func (s *VacancyService) UpdateVacancy(id uint, vacancy entities.Vacancy) error {
	return s.Repo.UpdateVacancy(id, vacancy)
}

func (s *VacancyService) DeleteVacancy(id uint) error {
	return s.Repo.DeleteVacancy(id)
}
