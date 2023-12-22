package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-file-upload-go/internal/application/usecases"
	"simple-file-upload-go/internal/infrastructure/http/routers"
	"simple-file-upload-go/internal/infrastructure/localstorage"
)

const (
	port          = 8089
	uploadsFolder = "./../uploads"
)

func main() {
	// init repositories
	lr, err := localstorage.NewLocalRepository(uploadsFolder)
	if err != nil {
		log.Panic("error initializing the repository")
	}

	// init use cases
	guc := usecases.NewGetAllFilesUseCase(lr)
	iuc := usecases.NewGetByIdUseCase(lr)
	upc := usecases.NewFileUploadUseCase(lr)

	// routers
	fr := routers.FileRouter(guc, iuc, upc)

	// main router
	mr := routers.MainRouter(fr)

	address := fmt.Sprintf(":%d", port)
	err = http.ListenAndServe(address, mr)
	if err != nil {
		log.Panicf("Error trying to launch the app, error: %v", err)
	}
	log.Printf("Server up on localhost:%d", port)
}
