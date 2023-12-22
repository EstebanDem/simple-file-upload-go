package requests

import (
	"io"
)

type UploadRequest struct {
	FileName string
	File     io.Reader
}
