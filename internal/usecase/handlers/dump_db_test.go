package handlers

import (
	"bytes"
	"context"
	"errors"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/domain/mocks"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

func TestDumpDb_HandleRequest(t *testing.T) {
	tt := []struct {
		name                string
		request             DumpDbRequest
		wantErr             bool
		createFileReturnArg error
	}{
		{
			name: "happy path",
			request: DumpDbRequest{
				FilePath: "shred_test.txt",
			},
			createFileReturnArg: nil,
			wantErr:             false,
		},
		{
			name: "happy path - file path is not given, use latest modified file",
			request: DumpDbRequest{
				FilePath: "",
			},
			createFileReturnArg: nil,
			wantErr:             false,
		},
		{
			name: "given file not exists",
			request: DumpDbRequest{
				FilePath: "dumpppp_test.txt",
			},
			createFileReturnArg: nil,
			wantErr:             true,
		},
		{
			name: "create file repository error",
			request: DumpDbRequest{
				FilePath: "dump_test.txt",
			},
			createFileReturnArg: errors.New("sqlite error"),
			wantErr:             true,
		},
	}
	ctx := context.TODO()
	var buf bytes.Buffer
	lgr := slog.New(slog.NewJSONHandler(&buf, nil))
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repository := &mocks.Repository{}
			repository.On("CreateFile", ctx, mock.Anything).Return(tc.createFileReturnArg)

			h := &DumpDb{
				logger:     lgr,
				repository: repository,
			}
			if err := h.HandleRequest(ctx, tc.request); (err != nil) != tc.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
