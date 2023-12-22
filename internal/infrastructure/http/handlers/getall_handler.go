package handlers

import (
	"encoding/json"
	"net/http"
	"simple-file-upload-go/internal/application/mapper"
	"simple-file-upload-go/internal/application/usecases"
)

func NewGetAllHandler(uc usecases.GetAllFilesUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		files, err := uc.GetAll()

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		}

		var filesJson []mapper.UploadFileJson

		for _, f := range *files {
			fj := mapper.UploadedFileToJson(f)
			filesJson = append(filesJson, fj)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filesJson)
	}
}

//func NewGetAllHandler(w http.ResponseWriter, r *http.Request) {
//	directory := "./../uploads"
//
//	files, err := os.ReadDir(directory)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, f := range files {
//		fmt.Println(f.Name())
//		fmt.Println(filepath.Ext(f.Name()))
//	}
//}
