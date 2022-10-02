package main

import (
	"context"
	"fmt"

	"github.com/emil14/neva/internal/core"
)

func Print(ctx context.Context, io core.IO) error {
	dataIn, err := io.In.Port("data")
	if err != nil {
		return err
	}

	dataOut, err := io.Out.Port("data")
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-dataIn:
			fmt.Print(msg)
			select {
			case dataOut <- msg:
				continue
			case <-ctx.Done(): // TODO try figure out better
				return ctx.Err()
			}
		}
	}
}
