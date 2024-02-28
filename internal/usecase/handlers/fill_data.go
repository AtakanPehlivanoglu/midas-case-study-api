package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// FillDataHandler is an abstraction for filling file with binary data use-case handler.
type FillDataHandler interface {
	HandleRequest(ctx context.Context, request FillDataRequest) error
}

type NewFillDataArgs struct {
	Logger *slog.Logger
}

func NewFillData(args NewFillDataArgs) (*FillData, error) {
	if args.Logger == nil {
		return nil, fmt.Errorf("logger is required")
	}
	return &FillData{
		logger: args.Logger,
	}, nil
}

// FillData is a request handler with all dependencies initialized.
type FillData struct {
	logger *slog.Logger
}

// FillDataRequest represents necessary POST /api/v1/file/fill request data for handler.
type FillDataRequest struct {
	FilePath string
	FileData string
}

func (h *FillData) HandleRequest(ctx context.Context, request FillDataRequest) error {
	filePath := fmt.Sprintf("%v/%v", AssetsPrefix, request.FilePath)
	fileData := request.FileData
	// Open the file with read-write access and truncate existing content
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(fileData))
	if err != nil {
		return err
	}

	return nil
}
