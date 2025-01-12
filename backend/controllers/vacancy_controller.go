package controllers

import (
	"enbektap/entities"
	"enbektap/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VacancyController struct {
	Service *services.VacancyService
}

func (c *VacancyController) CreateVacancy(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	var vacancy entities.Vacancy
	if err := ctx.ShouldBindJSON(&vacancy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := c.Service.CreateVacancy(vacancy); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vacancy"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Vacancy created successfully"})
}

func (c *VacancyController) GetVacancy(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	idStr := ctx.DefaultQuery("id", "")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	id := uint(id64)
	vacancy, err := c.Service.GetVacancy(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Vacancy not found"})
		return
	}

	ctx.JSON(http.StatusOK, vacancy)
}

func (c *VacancyController) GetAllVacancies(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	vacancies, err := c.Service.GetAllVacancies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vacancies"})
		return
	}

	ctx.JSON(http.StatusOK, vacancies)
}

func (c *VacancyController) UpdateVacancy(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPut {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	idStr := ctx.DefaultQuery("id", "")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	id := uint(id64)
	var vacancy entities.Vacancy
	if err := ctx.ShouldBindJSON(&vacancy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := c.Service.UpdateVacancy(id, vacancy); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vacancy"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vacancy updated successfully"})
}

func (c *VacancyController) DeleteVacancy(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodDelete {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	idStr := ctx.DefaultQuery("id", "")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	id := uint(id64)
	if err := c.Service.DeleteVacancy(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete vacancy"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vacancy deleted successfully"})
}
