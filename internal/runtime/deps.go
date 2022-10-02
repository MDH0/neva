package runtime

import (
	"context"

	"github.com/emil14/neva/internal/core"
	"github.com/emil14/neva/internal/runtime/src"
)

type Decoder interface {
	Decode([]byte) (src.Program, error)
}

type (
	Connector interface {
		Connect(context.Context, []Connection) error
	}
	Connection struct {
		Src       src.Connection
		Sender    chan core.Msg
		Receivers []chan core.Msg
	}
)

type (
	Effector interface {
		Effect(context.Context, Effects) error
	}
	Effects struct {
		Constants []ConstantEffect
		Operators []OperatorEffect
		Triggers  []TriggerEffect
	}
	ConstantEffect struct {
		OutPort chan core.Msg
		Msg     core.Msg
	}
	OperatorEffect struct {
		Ref src.OperatorRef
		IO  core.IO
	}
	TriggerEffect struct {
		InPort  chan core.Msg
		OutPort chan core.Msg
		Msg     core.Msg
	}
)
