package usecases

import (
	"simple-file-upload-go/internal/application/repository"
	"simple-file-upload-go/internal/domain"
)

type GetAllFilesUseCase interface {
	GetAll() (*[]domain.UploadedFile, error)
}

type getAllFilesUseCase struct {
	fileRepository repository.FileRepository
}

func NewGetAllFilesUseCase(fr repository.FileRepository) GetAllFilesUseCase {
	return &getAllFilesUseCase{
		fileRepository: fr,
	}
}

func (guc *getAllFilesUseCase) GetAll() (*[]domain.UploadedFile, error) {
	return guc.fileRepository.GetAll()
}
