package funcs

import (
	"context"

	"github.com/nevalang/neva/internal/runtime"
)

type intAdd struct{}

func (intAdd) Create(
	io runtime.FuncIO,
	_ runtime.Msg,
) (func(ctx context.Context), error) {
	seqIn, err := io.In.SingleInport("seq")
	if err != nil {
		return nil, err
	}

	resOut, err := io.Out.SingleOutport("res")
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) {
		var acc int64 = 0

		for {
			var item map[string]runtime.Msg
			select {
			case <-ctx.Done():
				return
			case msg := <-seqIn:
				item = msg.Map()
			}

			acc += item["data"].Int()

			if item["last"].Bool() {
				select {
				case <-ctx.Done():
					return
				case resOut <- runtime.NewIntMsg(acc):
					acc = 0 // reset
					continue
				}
			}
		}
	}, nil
}
