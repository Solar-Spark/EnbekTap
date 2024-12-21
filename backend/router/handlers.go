package router

import (
	"enbektap/database"
	"enbektap/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Handlers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}
	defer sqlDB.Close()

	switch r.Method {
	case http.MethodPost:
		var requestData struct {
			JobName     string `json:"Vacancy"`
			Salary      int    `json:"Salary"`
			JobType     string `json:"JobType"`
			Description string `json:"Description"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, "fail", "Invalid JSON data")
			return
		}

		if requestData.JobName == "" || requestData.Salary == 0 || requestData.JobType == "" || requestData.Description == "" {
			jsonResponse(w, http.StatusBadRequest, "fail", "All fields are required")
			return
		}

		err = models.CreateVacancy(db, requestData.JobName, requestData.JobType, requestData.Description, requestData.Salary)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, "fail", "Failed to save vacancy")
			return
		}

		jsonResponse(w, http.StatusOK, "success", "Vacancy successfully created")
		return

	case http.MethodGet:
		if r.URL.Query().Has("id") {
			id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
			if err != nil {
				jsonResponse(w, http.StatusBadRequest, "fail", "Invalid ID format")
				return
			}

			vacancy, err := models.ReadOneVacancy(db, id)
			if err != nil {
				jsonResponse(w, http.StatusNotFound, "fail", err.Error())
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(vacancy)
			return
		}

		vacancies, err := models.ReadVacancies(db)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, "fail", "Failed to fetch vacancies")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(vacancies)
		return

	case http.MethodPut:
		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, "fail", "Invalid ID format")
			return
		}

		var requestData struct {
			JobName     string `json:"Vacancy"`
			Salary      int    `json:"Salary"`
			JobType     string `json:"JobType"`
			Description string `json:"Description"`
		}

		err = json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, "fail", "Invalid JSON data")
			return
		}

		err = models.UpdateVacancy(db, id, requestData.JobName, requestData.JobType, requestData.Description, requestData.Salary)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, "fail", "Failed to update vacancy")
			return
		}

		jsonResponse(w, http.StatusOK, "success", "Vacancy successfully updated")
		return

	case http.MethodDelete:
		// Change this line to use query parameters like GET and PUT
		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, "fail", "Invalid ID format")
			return
		}

		err = models.DeleteVacancy(db, id)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, "fail", "Failed to delete vacancy")
			return
		}

		jsonResponse(w, http.StatusOK, "success", "Vacancy successfully deleted")
		return

	default:
		jsonResponse(w, http.StatusMethodNotAllowed, "fail", "Invalid HTTP method")
	}
}

func jsonResponse(w http.ResponseWriter, statusCode int, status, message string) {
	enableCors(&w) // Ensure CORS headers are set
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ResponseData{
		Status:  status,
		Message: message,
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
