package usecases

import (
	"simple-file-upload-go/internal/application/dtos/requests"
	"simple-file-upload-go/internal/application/repository"
)

type UploadFileUseCase interface {
	Upload(ur *requests.UploadRequest) error
}

type uploadFileUseCase struct {
	fileRepository repository.FileRepository
}

func NewFileUploadUseCase(fr repository.FileRepository) UploadFileUseCase {
	return &uploadFileUseCase{
		fileRepository: fr,
	}
}

func (upc *uploadFileUseCase) Upload(ur *requests.UploadRequest) error {
	return upc.fileRepository.Upload(ur)
}
