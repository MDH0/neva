package funcs

import (
	"context"

	"github.com/nevalang/neva/internal/runtime"
)

type listToStream struct{}

func (c listToStream) Create(
	io runtime.FuncIO,
	_ runtime.Msg,
) (func(ctx context.Context), error) {
	dataIn, err := io.In.SingleInport("data")
	if err != nil {
		return nil, err
	}

	seqOut, err := io.Out.SingleOutport("seq")
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) {
		for {
			var list []runtime.Msg

			select {
			case <-ctx.Done():
				return
			case dataMsg := <-dataIn:
				list = dataMsg.List()
			}

			for idx := 0; idx < len(list); idx++ {
				item := streamItem(
					list[idx],
					int64(idx),
					idx == len(list)-1,
				)

				select {
				case <-ctx.Done():
					return
				case seqOut <- item:
				}
			}
		}
	}, nil
}
