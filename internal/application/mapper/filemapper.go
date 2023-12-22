package mapper

import (
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"simple-file-upload-go/internal/domain"
	"time"
)

type UploadFileJson struct {
	ID         uuid.UUID `json:"id"`
	Extension  string    `json:"extension"`
	UploadedAt time.Time `json:"uploaded_at"`
}

func UploadedFileToJson(uf domain.UploadedFile) UploadFileJson {
	return UploadFileJson{
		ID:         uf.ID,
		Extension:  uf.Extension,
		UploadedAt: uf.UploadedAt,
	}
}

func DirEntryToUploadedFile(d os.DirEntry) (*domain.UploadedFile, error) {
	info, _ := d.Info()
	fileName, err := uuid.Parse(removeExtension(d.Name()))

	if err != nil {
		return nil, err
	}

	return &domain.UploadedFile{
		ID:         fileName,
		Extension:  filepath.Ext(d.Name()),
		UploadedAt: info.ModTime(),
	}, nil
}

func removeExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
