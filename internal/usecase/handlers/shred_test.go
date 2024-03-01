package handlers

import (
	"bytes"
	"context"
	"github.com/lu4p/shred"
	"log/slog"
	"testing"
)

func TestShred_HandleRequest(t *testing.T) {
	tt := []struct {
		name    string
		request ShredRequest
		wantErr bool
	}{
		{
			name: "happy path",
			request: ShredRequest{
				FilePath: "shred_test.txt",
			},
			wantErr: false,
		},
		{
			name: "given file not exists",
			request: ShredRequest{
				FilePath: "shred_testtt.txt",
			},
			wantErr: true,
		},
		{
			name: "no path given",
			request: ShredRequest{
				FilePath: "",
			},
			wantErr: true,
		},
	}
	ctx := context.TODO()
	var buf bytes.Buffer
	lgr := slog.New(slog.NewJSONHandler(&buf, nil))
	shredConfig := &shred.Conf{
		Times:  3,
		Zeros:  false,
		Remove: false,
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			h := &Shred{
				logger:      lgr,
				shredConfig: shredConfig,
			}
			if err := h.HandleRequest(ctx, tc.request); (err != nil) != tc.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
