package gojq

import "errors"

func (env *env) compileQuery(q *Query) error {
	if len(q.FuncDefs) > 0 {
		return errors.New("funcdef")
	}
	if err := env.compilePipe(q.Pipe); err != nil {
		return err
	}
	env.append(&code{op: opret})
	return nil
}

func (env *env) compilePipe(e *Pipe) error {
	for _, e := range e.Commas {
		if err := env.compileComma(e); err != nil {
			return err
		}
	}
	return nil
}

func (env *env) compileComma(e *Comma) error {
	if len(e.Alts) > 1 {
		return errors.New("compileComma")
	}
	return env.compileAlt(e.Alts[0])
}

func (env *env) compileAlt(e *Alt) error {
	if len(e.Right) > 0 {
		return errors.New("compileAlt")
	}
	return env.compileExpr(e.Left)
}

func (env *env) compileExpr(e *Expr) error {
	if e.Logic != nil && e.Bind == nil && e.Label == nil {
		return env.compileLogic(e.Logic)
	}
	return errors.New("compileExpr")
}

func (env *env) compileLogic(e *Logic) error {
	if len(e.Right) > 0 {
		return errors.New("compileLogic")
	}
	return env.compileAndExpr(e.Left)
}

func (env *env) compileAndExpr(e *AndExpr) error {
	if len(e.Right) > 0 {
		return errors.New("compileAndExpr")
	}
	return env.compileCompare(e.Left)
}

func (env *env) compileCompare(e *Compare) error {
	if e.Right != nil {
		return errors.New("compileCompare")
	}
	return env.compileArith(e.Left)
}

func (env *env) compileArith(e *Arith) error {
	if e.Right != nil {
		return errors.New("compileArith")
	}
	return env.compileFactor(e.Left)
}

func (env *env) compileFactor(e *Factor) error {
	if len(e.Right) > 0 {
		return errors.New("compileFactor")
	}
	return env.compileTerm(e.Left)
}

func (env *env) compileTerm(e *Term) error {
	return errors.New("compileTerm")
}

func (env *env) append(c *code) {
	env.codes = append(env.codes, c)
}