package request

import (
	"net/http"
)

type FillDataRequest struct {
	FilePath string
	FileData string
}

func (req FillDataRequest) Bind(r *http.Request) error {
	return nil
}
