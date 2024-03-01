package casestudyapi

import (
	"fmt"
	apprequest "github.com/AtakanPehlivanoglu/midas-case-study-api/internal/app/request"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/app/response"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/usecase/handlers"
	"github.com/go-chi/render"
	"net/http"
	"path/filepath"
)

// DumpDb handles POST /api/v1/file/dump
func (i *Implementation) DumpDb(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := i.logger

	request := &apprequest.DumpDbRequest{}

	if err := render.Bind(r, request); err != nil {
		errMessage := fmt.Errorf("error on binding request")
		logger.Error(errMessage.Error(), "err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.ErrInvalidRequest(errMessage))
		return
	}

	fileExtension := filepath.Ext(request.FilePath)
	if request.FilePath != "" && fileExtension != FileExtension {
		errMessage := fmt.Errorf("error on file path, should be either empty or '.txt'")
		logger.Error(errMessage.Error(), "filePath", request.FilePath)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.ErrInvalidRequest(errMessage))
		return
	}

	err := i.dumpDbHandler.HandleRequest(ctx, handlers.DumpDbRequest{
		FilePath: request.FilePath,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response.ErrInternalServer(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
