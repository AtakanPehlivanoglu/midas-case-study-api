package handlers

import (
	"context"
	"fmt"
	"github.com/lu4p/shred"
	"log/slog"
)

// ShredHandler is an abstraction for file shredder use-case handler.
type ShredHandler interface {
	HandleRequest(ctx context.Context, request ShredRequest) error
}

type NewShredArgs struct {
	Logger *slog.Logger
}

func NewShred(args NewShredArgs) (*Shred, error) {
	if args.Logger == nil {
		return nil, fmt.Errorf("logger is required")
	}
	return &Shred{
		logger: args.Logger,
	}, nil
}

// Shred is a request handler with all dependencies initialized.
type Shred struct {
	logger *slog.Logger
}

// ShredRequest represents necessary GET /api/v1/file/shred/{filePath} request data for handler.
type ShredRequest struct {
	FilePath string
}

func (h *Shred) HandleRequest(ctx context.Context, request ShredRequest) error {
	logger := h.logger
	filePath := request.FilePath
	shredConfig := shred.Conf{Times: 3, Zeros: true, Remove: false}
	err := shredConfig.File(filePath)
	if err != nil {
		logger.Error("err on shredding file", "err", err)
		return err
	}

	return nil
}
