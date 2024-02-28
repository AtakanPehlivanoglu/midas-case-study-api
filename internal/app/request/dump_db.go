package request

import "net/http"

type DumpDbRequest struct {
	FilePath string
}

func (req DumpDbRequest) Bind(r *http.Request) error {
	return nil
}
