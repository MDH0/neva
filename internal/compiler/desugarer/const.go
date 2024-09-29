package desugarer

import (
	"github.com/nevalang/neva/internal/compiler"
	src "github.com/nevalang/neva/internal/compiler/sourcecode"
)

// handleConst handles case when constant has integer value and type is float.
func (d Desugarer) handleConst(constant src.Const) (src.Const, *compiler.Error) {
	if constant.Message == nil {
		return constant, nil
	}
	if constant.TypeExpr.String() != "float" {
		return constant, nil
	}
	if constant.Message.Float != nil {
		return constant, nil
	}
	return src.Const{
		TypeExpr: constant.TypeExpr,
		Message: &src.Message{
			Float: compiler.Pointer(float64(*constant.Message.Int)),
		},
	}, nil
}
