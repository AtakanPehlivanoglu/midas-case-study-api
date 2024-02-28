package casestudyapi

import (
	apprequest "github.com/AtakanPehlivanoglu/midas-case-study-api/internal/app/request"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/app/response"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/usecase/handlers"
	"github.com/go-chi/render"
	"net/http"
)

// DumpDb handles /api/v1/file/dump
func (i *Implementation) DumpDb(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := i.logger

	request := &apprequest.DumpDbRequest{}

	if err := render.Bind(r, request); err != nil {
		logger.Error("error on binding request", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.ErrInvalidRequest(err))
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
