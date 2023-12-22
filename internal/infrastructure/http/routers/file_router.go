package routers

import (
	"github.com/estebandem/simple-file-upload-go/internal/application/usecases"
	"github.com/estebandem/simple-file-upload-go/internal/infrastructure/http/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func FileRouter(guc usecases.GetAllFilesUseCase, iuc usecases.GetByIdUseCase, upc usecases.UploadFileUseCase) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/upload", handlers.NewUploadHandler(upc)).Methods(http.MethodPost)
	r.HandleFunc("/uploads", handlers.NewGetAllHandler(guc)).Methods(http.MethodGet)
	r.HandleFunc("/uploads/{id}", handlers.NewGetByIdHandler(iuc)).Methods(http.MethodGet)

	return r
}
