package funcs

import (
	"context"

	"github.com/nevalang/neva/internal/runtime"
)

type intSub struct{}

func (intSub) Create(io runtime.FuncIO, _ runtime.Msg) (func(ctx context.Context), error) {
	seqIn, err := io.In.Port("seq")
	if err != nil {
		return nil, err
	}

	resOut, err := io.Out.Port("res")
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) {
		flag := false
		var res int64
		for {
			select {
			case <-ctx.Done():
				return
			case streamItem := <-seqIn:
				if streamItem == nil {
					select {
					case <-ctx.Done():
						return
					case resOut <- runtime.NewIntMsg(res):
						flag = false
						res = 0
						continue
					}
				}

				if !flag {
					res = streamItem.Int()
					flag = true
				} else {
					res -= streamItem.Int()
				}
			}
		}
	}, nil
}
