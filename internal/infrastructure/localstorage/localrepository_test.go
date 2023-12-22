package localstorage_test

import (
	"errors"
	"github.com/estebandem/simple-file-upload-go/internal/application/dtos/requests"
	"github.com/estebandem/simple-file-upload-go/internal/application/repository"
	"github.com/estebandem/simple-file-upload-go/internal/infrastructure/localstorage"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const testUploadLocation = "./uploadsfoldertest"

type testCase struct {
	name        string
	id          uuid.UUID
	expectedErr error
}

func TestLocalRepository_Get(t *testing.T) {
	cleanupTestDirectory()
	localRepository, err := localstorage.NewLocalRepository(testUploadLocation)

	if err != nil {
		panic(err)
	}

	testCases := []testCase{
		{
			name:        "no file found",
			id:          uuid.MustParse("94bfc4de-edb5-4774-85d0-fde18e9951b8"),
			expectedErr: repository.ErrIdNotFound,
		},
		{
			name:        "file found",
			id:          uuid.MustParse("79b39a92-e593-457a-aeb2-df874dea6c22"),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := localRepository.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}

func TestLocalRepository_GetAll(t *testing.T) {
	cleanupTestDirectory()
	localRepository, err := localstorage.NewLocalRepository(testUploadLocation)

	if err != nil {
		panic(err)
	}

	t.Run("get_all_ok", func(t *testing.T) {
		all, err := localRepository.GetAll()
		if err != nil {
			t.Error("error retrieving all files")
		}

		if len(*all) != 1 {
			t.Errorf("should be 1 file, but %d", len(*all))
		}
	})

}

func TestLocalRepository_Upload(t *testing.T) {
	cleanupTestDirectory()
	localRepository, err := localstorage.NewLocalRepository(testUploadLocation)

	if err != nil {
		panic(err)
	}

	t.Run("upload_file_ok", func(t *testing.T) {
		uploadRequest := &requests.UploadRequest{
			FileName: "testfile.txt",
			File:     strings.NewReader("Hello, world!"),
		}

		err := localRepository.Upload(uploadRequest)
		if err != nil {
			t.Errorf("unexpected error uploading file: %v", err)
		}

		// Verify that the file was added successfully
		allFiles, err := localRepository.GetAll()
		if err != nil {
			t.Error("error retrieving all files")
		}

		if len(*allFiles) != 2 {
			t.Errorf("should be 2 files, but %d", len(*allFiles))
		}
	})
	cleanupTestDirectory()
}

func cleanupTestDirectory() {
	err := os.RemoveAll(testUploadLocation)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(testUploadLocation, os.ModePerm)
	if err != nil {
		panic(err)
	}

	createTestFile("79b39a92-e593-457a-aeb2-df874dea6c22.txt", "Hi!")
}

func createTestFile(fileName string, content string) {
	filePath := filepath.Join(testUploadLocation, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}

}
