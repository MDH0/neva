package funcs

import (
	"context"
	"strings"

	"github.com/nevalang/neva/internal/runtime"
)

type stringsToUpper struct{}

func (p stringsToUpper) Create(io runtime.IO, _ runtime.Msg) (func(ctx context.Context), error) {
	dataIn, err := io.In.Single("data")
	if err != nil {
		return nil, err
	}

	resOut, err := io.Out.Single("res")
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) {
		for {
			data, ok := dataIn.Receive(ctx)
			if !ok {
				return
			}

			result := strings.ToUpper(data.Str())
			if !resOut.Send(ctx, runtime.NewStringMsg(result)) {
				return
			}
		}
	}, nil
}
