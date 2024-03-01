package handlers

import (
	"bytes"
	"context"
	"log/slog"
	"testing"
)

func TestFillData_HandleRequest(t *testing.T) {
	tt := []struct {
		name    string
		request FillDataRequest
		wantErr bool
	}{
		{
			name: "happy path",
			request: FillDataRequest{
				FilePath: "shred_test.txt",
				FileData: "01001000,01100101,01101100,01101100\n00110011,00110011\n00110011,00110001\n00110011,00110001\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01101100,01101100,01101100,01001000,01100101,01100111",
			},
			wantErr: false,
		},
		{
			name: "given file not exists",
			request: FillDataRequest{
				FilePath: "shred_testttt.txt",
				FileData: "01001000,01100101,01101100,01101100\n00110011,00110011\n00110011,00110001\n00110011,00110001\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01101100,01101100,01101100,01001000,01100101,01100111",
			},
			wantErr: true,
		},
		{
			name: "no path given",
			request: FillDataRequest{
				FilePath: "",
				FileData: "\n",
			},
			wantErr: true,
		},
	}
	ctx := context.TODO()
	var buf bytes.Buffer
	lgr := slog.New(slog.NewJSONHandler(&buf, nil))
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			h := &FillData{
				logger: lgr,
			}
			if err := h.HandleRequest(ctx, tc.request); (err != nil) != tc.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
