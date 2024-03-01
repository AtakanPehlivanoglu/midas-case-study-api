package casestudyapi

import (
	"fmt"
	apprequest "github.com/AtakanPehlivanoglu/midas-case-study-api/internal/app/request"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/app/response"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/usecase/handlers"
	"github.com/go-chi/render"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	fileLinesThreshold = 10
)

// FillData handles POST /api/v1/file/fill
func (i *Implementation) FillData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := i.logger

	request := &apprequest.FillDataRequest{}

	if err := render.Bind(r, request); err != nil {
		errMessage := fmt.Errorf("error on binding request")
		logger.Error(errMessage.Error(), "err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.ErrInvalidRequest(errMessage))
		return
	}

	fileExtension := filepath.Ext(request.FilePath)
	if fileExtension != FileExtension {
		errMessage := fmt.Errorf("error on file path, should be '.txt'")
		logger.Error(errMessage.Error(), "filePath", request.FilePath)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.ErrInvalidRequest(errMessage))
		return
	}

	if !isValidBinaryData(request.FileData) {
		logger.Error("error on validating binary data", "fileData", request.FileData)
		w.WriteHeader(http.StatusBadRequest)
		// more explicit explanation to caller
		err := fmt.Errorf("invalid file data, use ',' separator and %q for new lines with 8 bits of binary data", '\n')
		w.Write(response.ErrInvalidRequest(err))
		return
	}

	numberOfLines := countLines(request.FileData)

	if numberOfLines < fileLinesThreshold {
		logger.Error("error on number of lines", "lines", numberOfLines, "threshold", fileLinesThreshold)
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("number of lines are less than expected, lines: %v, threshold: %v", numberOfLines, fileLinesThreshold)
		w.Write(response.ErrInvalidRequest(err))
		return
	}

	err := i.fillDataHandler.HandleRequest(ctx, handlers.FillDataRequest{
		FilePath: request.FilePath,
		FileData: request.FileData,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response.ErrInternalServer(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func isValidBinaryData(data string) bool {
	// Regular expression to match binary data with ',' separator and '\n' for new lines
	binaryRegex := regexp.MustCompile("^([01]{8},)*[01]{8}(\\n([01]{8},)*[01]{8})*\\n?$")

	return binaryRegex.MatchString(data)
}

func countLines(s string) int {
	return strings.Count(s, "\n") + 1
}
