package localstorage

import (
	"fmt"
	"github.com/estebandem/simple-file-upload-go/internal/application/dtos/requests"
	"github.com/estebandem/simple-file-upload-go/internal/application/mapper"
	"github.com/estebandem/simple-file-upload-go/internal/application/repository"
	"github.com/estebandem/simple-file-upload-go/internal/domain"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type LocalRepository struct {
	path  string
	files map[uuid.UUID]*domain.UploadedFile
	mu    sync.Mutex
}

func NewLocalRepository(path string) (*LocalRepository, error) {
	lr := &LocalRepository{
		path: path,
	}
	err := lr.sync()
	if err != nil {
		return nil, err
	}
	return lr, nil
}

func (l *LocalRepository) GetAll() (*[]domain.UploadedFile, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var all []domain.UploadedFile
	for _, f := range l.files {
		all = append(all, *f)
	}

	return &all, nil

}

func (l *LocalRepository) Upload(ur *requests.UploadRequest) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	newId := uuid.New()
	fileExtension := filepath.Ext(ur.FileName)
	fileName := fmt.Sprintf("%s%s", newId, fileExtension)
	dst, err := os.Create(filepath.Join(l.path, fileName))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, ur.File)

	if err != nil {
		return err
	}

	l.files[newId] = &domain.UploadedFile{
		ID:         newId,
		Extension:  fileExtension,
		UploadedAt: time.Time{},
	}
	return nil
}

func (l *LocalRepository) Get(id uuid.UUID) ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	f, ok := l.files[id]
	if !ok {
		return nil, repository.ErrIdNotFound
	}

	path := fmt.Sprintf("%s/%s%s", l.path, id.String(), f.Extension)
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// sync loads all the files present in the folder
func (l *LocalRepository) sync() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.files = make(map[uuid.UUID]*domain.UploadedFile)

	files, err := os.ReadDir(l.path)
	if err != nil {
		return err
	}

	for _, f := range files {
		uf, err := mapper.DirEntryToUploadedFile(f)
		if err != nil {
			return err
		}
		l.files[uf.ID] = uf
	}

	return nil
}
