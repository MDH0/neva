package types

func (expr Expr) String() string {
	var s string

	if expr.Lit.RecLit != nil {
		s += "{"
		for fieldName, fieldExpr := range expr.Lit.RecLit {
			s += " " + fieldName + ": " + fieldExpr.String() + " "
		}
		s += "}"
		return s
	}

	if len(expr.Inst.Args) == 0 {
		return expr.Inst.Ref
	}

	s = expr.Inst.Ref + "<"
	for i, arg := range expr.Inst.Args {
		s += arg.String()
		if i < len(expr.Inst.Args)-1 {
			s += ", "
		}
	}
	s += ">"

	return s
}
