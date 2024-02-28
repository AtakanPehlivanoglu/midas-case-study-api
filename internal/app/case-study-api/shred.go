package casestudyapi

import (
	"fmt"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/app/response"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/usecase/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Shred handles /api/v1/file/shred/{filePath}
func (i *Implementation) Shred(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := i.logger

	var filePath string

	if filePath = chi.URLParam(r, "filePath"); filePath == "" {
		errMessage := fmt.Errorf("filePath is wrong")
		logger.Error(errMessage.Error(), "filePath", filePath)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.ErrInvalidRequest(errMessage))
		return
	}

	err := i.shredHandler.HandleRequest(ctx, handlers.ShredRequest{
		FilePath: filePath,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response.ErrInternalServer(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
