package domain

import (
	"context"
)

type Repository interface {
	CreateFile(ctx context.Context, args CreateFileArgs) error
}

// CreateFileArgs to call CreateFile repository method
type CreateFileArgs struct {
	FileName string
	FileData string
}
