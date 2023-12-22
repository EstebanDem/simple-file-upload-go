package handlers

import (
	"encoding/json"
	"github.com/estebandem/simple-file-upload-go/internal/application/usecases"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func NewGetByIdHandler(guc usecases.GetByIdUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]

		id, err := uuid.Parse(idStr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		}

		file, err := guc.GetById(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		}

		contentType := http.DetectContentType(file)
		w.Header().Set("Content-Type", contentType)
		filename := "download.png"
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)

		w.WriteHeader(http.StatusOK)
		w.Write(file)
	}
}
