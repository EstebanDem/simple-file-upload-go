package domain

import (
	"github.com/google/uuid"
	"time"
)

type UploadedFile struct {
	ID         uuid.UUID
	Extension  string
	UploadedAt time.Time
}

func NewUploadedFile(extension string) *UploadedFile {
	return &UploadedFile{
		ID:         uuid.New(),
		Extension:  extension,
		UploadedAt: time.Now(),
	}
}
