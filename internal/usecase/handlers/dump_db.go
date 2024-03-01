package handlers

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/domain"
	sqliterepo "github.com/AtakanPehlivanoglu/midas-case-study-api/internal/infra/sqlite"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	AssetsPrefix  = "./assets"
	FileExtension = ".txt"
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
	var fileName string
	// meaning that no path is given hence use latest file used in fill-data operation
	if request.FilePath == "" {
		var err error
		fileName, err = getLastFilledFileName()
		if err != nil {
			logger.Error("error on get_last_filled_file_name", "err", err)
			return err
		}
	} else {
		fileName = filepath.Base(request.FilePath)
	}

	filePath := fmt.Sprintf("%v/%v", AssetsPrefix, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("error opening file", "err", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	createFileQuery := sqliterepo.CreateFileQueryBuilder()
	for scanner.Scan() {
		line := scanner.Text()
		text := binaryToText(line)
		// concurrency can be added to execute sql query in batches
		// rather than having very large query string in memory
		createFileQuery += sqliterepo.CreateFileBulkInsertBuilder(fileName, text)
	}

	// remove trailing comma from SQL
	createFileQuery = strings.TrimSuffix(createFileQuery, ",")

	if err = scanner.Err(); err != nil {
		logger.Error("error on reading from file", "err", err)
		return err
	}

	err = h.repository.CreateFile(ctx, domain.CreateFileArgs{
		Query: createFileQuery,
	})

	if err != nil {
		logger.Error("error on repository create_file", "err", err)
		return err
	}

	return nil
}

// getLastFilledFileName gets the latest file used with fill-data operation
func getLastFilledFileName() (string, error) {
	latestFile := struct {
		ModificationTime time.Time
		FileName         string
	}{}

	files, err := os.ReadDir(fmt.Sprintf("%v/", AssetsPrefix))
	if err != nil {
		return "", err
	}

	for _, file := range files {
		fileExtension := filepath.Ext(fmt.Sprintf("%v/%v", AssetsPrefix, file.Name()))
		if fileExtension == FileExtension {
			fileInfo, statErr := os.Stat(fmt.Sprintf("%v/%v", AssetsPrefix, file.Name()))
			if statErr != nil {
				return "", statErr
			}
			// get the latest modified file used with fill-data
			if fileInfo.ModTime().After(latestFile.ModificationTime) {
				latestFile.FileName = fileInfo.Name()
				latestFile.ModificationTime = fileInfo.ModTime()
			}
		}
	}
	return latestFile.FileName, nil
}

func binaryToText(binaryData string) string {
	dataBytes := strings.Split(binaryData, ",")
	var text bytes.Buffer

	for _, byteStr := range dataBytes {
		decimal, _ := strconv.ParseInt(byteStr, 2, 64)
		text.WriteByte(byte(decimal))
	}

	return text.String()
}
