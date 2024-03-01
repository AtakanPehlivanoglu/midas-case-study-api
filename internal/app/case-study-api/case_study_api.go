package casestudyapi

import (
	"fmt"
	"github.com/AtakanPehlivanoglu/midas-case-study-api/internal/usecase/handlers"
	"log/slog"
)

// Implementation is struct for whole API implementation in one place.
type Implementation struct {
	logger          *slog.Logger
	shredHandler    handlers.ShredHandler
	fillDataHandler handlers.FillDataHandler
	dumpDbHandler   handlers.DumpDbHandler
}

func NewCaseStudyAPI(args NewCaseStudyAPIArgs) (*Implementation, error) {
	if args.Logger == nil {
		return nil, fmt.Errorf("logger is required")
	}
	if args.ShredHandler == nil {
		return nil, fmt.Errorf("shredHandler is required")
	}
	if args.FillDataHandler == nil {
		return nil, fmt.Errorf("fillDataHandler is required")
	}
	if args.DumpDbHandler == nil {
		return nil, fmt.Errorf("dumpDbHandler is required")
	}

	return &Implementation{
		logger:          args.Logger,
		shredHandler:    args.ShredHandler,
		fillDataHandler: args.FillDataHandler,
		dumpDbHandler:   args.DumpDbHandler,
	}, nil
}

type NewCaseStudyAPIArgs struct {
	Logger          *slog.Logger
	ShredHandler    handlers.ShredHandler
	FillDataHandler handlers.FillDataHandler
	DumpDbHandler   handlers.DumpDbHandler
}
