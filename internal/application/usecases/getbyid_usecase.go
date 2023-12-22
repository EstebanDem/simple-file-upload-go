package usecases

import (
	"github.com/google/uuid"
	"simple-file-upload-go/internal/application/repository"
)

type GetByIdUseCase interface {
	GetById(id uuid.UUID) ([]byte, error)
}

type getByIdUseCase struct {
	fileRepository repository.FileRepository
}

func NewGetByIdUseCase(fr repository.FileRepository) GetByIdUseCase {
	return &getByIdUseCase{
		fileRepository: fr,
	}
}

func (guc *getByIdUseCase) GetById(id uuid.UUID) ([]byte, error) {
	return guc.fileRepository.Get(id)
}
