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
	Logger      *slog.Logger
	ShredConfig *shred.Conf
}

func NewShred(args NewShredArgs) (*Shred, error) {
	if args.Logger == nil {
		return nil, fmt.Errorf("logger is required")
	}
	return &Shred{
		logger:      args.Logger,
		shredConfig: args.ShredConfig,
	}, nil
}

// Shred is a request handler with all dependencies initialized.
type Shred struct {
	logger      *slog.Logger
	shredConfig *shred.Conf
}

// ShredRequest represents necessary DELETE /api/v1/file/shred/{filePath} request data for handler.
type ShredRequest struct {
	FilePath string
}

func (h *Shred) HandleRequest(ctx context.Context, request ShredRequest) error {
	logger := h.logger
	shredConfig := h.shredConfig

	filePath := fmt.Sprintf("%v/%v", AssetsPrefix, request.FilePath)

	err := shredConfig.File(filePath)
	if err != nil {
		logger.Error("error on shredding file", "err", err)
		return err
	}

	return nil
}
