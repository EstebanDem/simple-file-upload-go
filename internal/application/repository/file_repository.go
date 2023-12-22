package repository

import (
	"errors"
	"github.com/estebandem/simple-file-upload-go/internal/application/dtos/requests"
	"github.com/estebandem/simple-file-upload-go/internal/domain"
	"github.com/google/uuid"
)

var (
	ErrIdNotFound = errors.New("id not found")
)

type FileRepository interface {
	GetAll() (*[]domain.UploadedFile, error)
	Upload(ur *requests.UploadRequest) error
	Get(uuid uuid.UUID) ([]byte, error)
}
