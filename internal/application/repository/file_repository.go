package repository

import (
	"errors"
	"github.com/google/uuid"
	"simple-file-upload-go/internal/application/dtos/requests"
	"simple-file-upload-go/internal/domain"
)

var (
	ErrIdNotFound = errors.New("id not found")
)

type FileRepository interface {
	GetAll() (*[]domain.UploadedFile, error)
	Upload(ur *requests.UploadRequest) error
	Get(uuid uuid.UUID) ([]byte, error)
}
