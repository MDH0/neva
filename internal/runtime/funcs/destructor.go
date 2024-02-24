package funcs

import (
	"context"

	"github.com/nevalang/neva/internal/runtime"
)

type destructor struct{}

func (d destructor) Create(io runtime.FuncIO, _ runtime.Msg) (func(ctx context.Context), error) {
	outport, err := io.In.Port("msg")
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-outport:
			}
		}
	}, nil
}
