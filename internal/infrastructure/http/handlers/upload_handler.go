package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/estebandem/simple-file-upload-go/internal/application/dtos/requests"
	"github.com/estebandem/simple-file-upload-go/internal/application/usecases"
	"net/http"
)

func NewUploadHandler(upc usecases.UploadFileUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request with Content-Type:", r.Header.Get("Content-Type"))

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		ur := &requests.UploadRequest{
			FileName: header.Filename,
			File:     file,
		}

		err = upc.Upload(ur)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(jsonResponse{Message: "file uploaded successfully", Status: http.StatusCreated})
	}
}

type jsonResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
