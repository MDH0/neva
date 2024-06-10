package funcs

import (
	"context"

	"github.com/nevalang/neva/internal/runtime"
)

type if_ struct{}

func (p if_) Create(io runtime.FuncIO, _ runtime.Msg) (func(ctx context.Context), error) {
	dataIn, err := io.In.Port("data")
	if err != nil {
		return nil, err
	}

	okOut, err := io.Out.Port("then")
	if err != nil {
		return nil, err
	}

	elseOut, err := io.Out.Port("else")
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) {
		var (
			val1 runtime.Msg
		)

		for {
			select {
			case <-ctx.Done():
				return
			case val1 = <-dataIn:
			}

			select {
			case <-ctx.Done():
				return

			default:
				if val1.Bool() {
					okOut <- val1
				} else {
					elseOut <- val1
				}
			}
		}
	}, nil
}
