package handlers

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/domain"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	AssetsPrefix = "./assets"
)

// DumpDbHandler is an abstraction for dumping binary file into db use-case handler.
type DumpDbHandler interface {
	HandleRequest(ctx context.Context, request DumpDbRequest) error
}

type NewDumpDbArgs struct {
	Logger     *slog.Logger
	Repository domain.Repository
}

func NewDumpDb(args NewDumpDbArgs) (*DumpDb, error) {
	if args.Logger == nil {
		return nil, fmt.Errorf("logger is required")
	}
	return &DumpDb{
		logger:     args.Logger,
		repository: args.Repository,
	}, nil
}

// DumpDb is a request handler with all dependencies initialized.
type DumpDb struct {
	logger     *slog.Logger
	repository domain.Repository
}

// DumpDbRequest represents necessary POST /api/v1/file/dump request data for handler.
type DumpDbRequest struct {
	FilePath string
}

func (h *DumpDb) HandleRequest(ctx context.Context, request DumpDbRequest) error {
	logger := h.logger
	fileName := filepath.Base(request.FilePath)
	filePath := fmt.Sprintf("%v/%v", AssetsPrefix, request.FilePath)

	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("error opening file", "err", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		text, _ := binaryToText(line)
		// prepare insert statement here
		// concurrency can be added in batches to execute sql stmt
		// to avoid having very large string in memory
		logger.Warn(text)
	}

	if err = scanner.Err(); err != nil {
		logger.Error("error reading from file", "err", err)
	}

	err = h.repository.CreateFile(ctx, domain.CreateFileArgs{
		FileName: fileName,
	})

	if err != nil {
		logger.Error("error on repository create_file", "err", err)
	}

	return nil
}

func binaryToText(binaryData string) (string, error) {
	dataBytes := strings.Split(binaryData, ",")
	var text bytes.Buffer

	for _, byteStr := range dataBytes {
		decimal, err := strconv.ParseInt(byteStr, 2, 64)
		if err != nil {
			return "", err
		}
		text.WriteByte(byte(decimal))
	}

	return text.String(), nil
}
