//go:generate mockery --name Repository --output mocks --outpkg mocks --case underscore --dir .
package domain

import (
	"context"
)

type Repository interface {
	CreateFile(ctx context.Context, args CreateFileArgs) error
}

// CreateFileArgs to call CreateFile repository method
type CreateFileArgs struct {
	Query string
}
