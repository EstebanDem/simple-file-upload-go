package handlers

import (
	"fmt"
	"net/http"
	"simple-file-upload-go/internal/application/dtos/requests"
	"simple-file-upload-go/internal/application/usecases"
)

func NewUploadHandler(upc usecases.UploadFileUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request with Content-Type:", r.Header.Get("Content-Type"))

		// Parse the form data to get the file
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

		w.Write([]byte("File uploaded successfully"))
	}
}
