package controllers

import (
	"enbektap/entities"
	"enbektap/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type VacancyController struct {
	Service *services.VacancyService
}

func (c *VacancyController) CreateVacancy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var vacancy entities.Vacancy
	if err := json.NewDecoder(r.Body).Decode(&vacancy); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateVacancy(vacancy); err != nil {
		http.Error(w, "Failed to create vacancy", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vacancy created successfully"})
}

func (c *VacancyController) GetVacancy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64) // Parse as uint64
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	id := uint(id64) // Convert uint64 to uint
	vacancy, err := c.Service.GetVacancy(id)
	if err != nil {
		http.Error(w, "Vacancy not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(vacancy)
}

func (c *VacancyController) GetAllVacancies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vacancies, err := c.Service.GetAllVacancies()
	if err != nil {
		http.Error(w, "Failed to fetch vacancies", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vacancies)
}

func (c *VacancyController) UpdateVacancy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64) // Parse as uint64
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	id := uint(id64) // Convert uint64 to uint
	var vacancy entities.Vacancy
	if err := json.NewDecoder(r.Body).Decode(&vacancy); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.Service.UpdateVacancy(id, vacancy); err != nil {
		http.Error(w, "Failed to update vacancy", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vacancy updated successfully"})
}

func (c *VacancyController) DeleteVacancy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64) // Parse as uint64
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	id := uint(id64) // Convert uint64 to uint
	if err := c.Service.DeleteVacancy(id); err != nil {
		http.Error(w, "Failed to delete vacancy", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vacancy deleted successfully"})
}
