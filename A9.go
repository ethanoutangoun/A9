package main

import (
	"errors"
	"fmt"
)

type ExprC interface{}
type numC struct {
	n float64
}
type idC struct {
	s string
}
type strC struct {
	str string
}
type appC struct {
	fun  ExprC
	args []ExprC
}
type lamC struct {
	args []string
	body ExprC
}
type ifC struct {
	ifExpr   ExprC
	thenExpr ExprC
	elseExpr ExprC
}

type Value interface{}
type closV struct {
	args []string
	body ExprC
	env  []Binding
}

type Binding struct {
	name string
	val  Value
}

var topEnv = []Binding{
	Binding{name: "+", val: "+"},
	Binding{name: "-", val: "-"},
	Binding{name: "*", val: "*"},
	Binding{name: "/", val: "/"},
	Binding{name: "<=", val: "<="},
	Binding{name: "==", val: "=="},
	Binding{name: "true", val: true},
	Binding{name: "false", val: false},
}

func interp(e ExprC, environment []Binding) Value {
	switch e.(type) {
	case numC:
		return e.(numC).n
	case strC:
		return e.(strC).str
	case idC:
		return lookup(e.(idC).s, environment)
	case lamC:
		return closV{args: e.(lamC).args, body: e.(lamC).body, env: environment}
	case ifC:
		ifVal := interp(e.(ifC).ifExpr, environment)
		switch ifVal.(type) {
		case bool:
			if ifVal.(bool) {
				return interp(e.(ifC).thenExpr, environment)
			} else {
				return interp(e.(ifC).elseExpr, environment)
			}
		default:
			return errors.New("if condition should be boolean value")
		}
	case appC:
		f := e.(appC).fun
		args := e.(appC).args
		fd := interp(f, environment)
		fArgs := make([]Value, len(args))
		for i, arg := range args {
			fArgs[i] = interp(arg, environment)
		}
		switch fd.(type) {
		case string:
			switch fd.(string) {
			case "+":
				return getBinop("+", fArgs[0], fArgs[1])
			case "-":
				return getBinop("-", fArgs[0], fArgs[1])
			case "*":
				return getBinop("*", fArgs[0], fArgs[1])
			case "/":
				return getBinop("/", fArgs[0], fArgs[1])
			case "<=":
				return getBinop("<=", fArgs[0], fArgs[1])
			case "==":
				return getBinop("==", fArgs[0], fArgs[1])
			default:
				return errors.New("Error in interp AppC")
			}
		case closV:
			closure := fd.(closV)
			return interp(closure.body, getEnv(closure.args, fArgs, closure.env))
		default:
			return errors.New("Error in interp AppC")
		}
	default:
		return errors.New("Invalid ExprC syntax")
	}
}

func lookup(forName string, environment []Binding) Value {
	if len(environment) == 0 {
		return errors.New("user-error No value match!")
	}
	if forName == environment[0].name {
		return environment[0].val
	}
	return lookup(forName, environment[1:])
}

func getBinop(op Value, l Value, r Value) Value {
	if !(isReal(l) && isReal(r)) {
		panic("Binop left or right value are not numbers")
	}
	switch op := op.(type) {
	case string:
		if op == "+" {
			return l.(float64) + r.(float64)
		}
		if op == "-" {
			return l.(float64) - r.(float64)
		}
		if op == "*" {
			return l.(float64) * r.(float64)
		}
		if op == "/" {
			if r.(float64) == 0 {
				panic("division by zero")
			}
			return l.(float64) / r.(float64)
		}
		if op == "<=" {
			return l.(float64) <= r.(float64)
		}
		if op == "==" {
			return l.(float64) == r.(float64)
		}
	}
	panic("Invalid Binop Syntax")
}

func isReal(x Value) bool {
	_, ok := x.(float64)
	return ok
}

func getEnv(s []string, l []Value, env []Binding) []Binding {
	if len(s) == 0 {
		return env
	}
	return append([]Binding{Binding{s[0], l[0]}}, getEnv(s[1:], l[1:], env)...)
}

func main() {
	testExprC1 := appC{fun: idC{s: "+"}, args: []ExprC{numC{n: 3}, numC{n: 4}}}
	testExprC2 := appC{fun: idC{s: "-"}, args: []ExprC{numC{n: 3}, numC{n: 4}}}
	testExprC3 := appC{fun: idC{s: "/"}, args: []ExprC{numC{n: 3}, numC{n: 4}}}
	testExprC4 := appC{fun: idC{s: "*"}, args: []ExprC{numC{n: 3}, numC{n: 4}}}
	testExprC5 := appC{fun: idC{s: "=="}, args: []ExprC{numC{n: 3}, numC{n: 4}}}
	testExprC6 := appC{fun: idC{s: "<="}, args: []ExprC{numC{n: 3}, numC{n: 4}}}
	testExprC7 := appC{fun: idC{s: "<="}, args: []ExprC{numC{n: 5}, numC{n: 4}}}
	testExprC8 := ifC{ifExpr: appC{fun: idC{s: "<="}, args: []ExprC{numC{n: 4}, numC{n: 5}}}, thenExpr: numC{n: 4}, elseExpr: numC{n: 5}}
	fmt.Println(`appC{idC{"+"}, []{{numC{n: 3}, numC{n: 4}} is : `, interp(testExprC1, topEnv))
	fmt.Println(`appC{idC{"-"}, []{{numC{n: 3}, numC{n: 4}} is : `, interp(testExprC2, topEnv))
	fmt.Println(`appC{idC{"/"}, []{{numC{n: 3}, numC{n: 4}} is : `, interp(testExprC3, topEnv))
	fmt.Println(`appC{idC{"*"}, []{{numC{n: 3}, numC{n: 4}} is : `, interp(testExprC4, topEnv))
	fmt.Println(`appC{idC{"=="}, []{{numC{n: 3}, numC{n: 4}} is : `, interp(testExprC5, topEnv))
	fmt.Println(`appC{idC{"<="}, []{{numC{n: 3}, numC{n: 4}} is : `, interp(testExprC6, topEnv))
	fmt.Println(`appC{idC{"<="}, []{{numC{n: 5}, numC{n: 4}} is : `, interp(testExprC7, topEnv))
	fmt.Println(`ifC{ifExpr: appC{fun: idC{s: "<="}, args: []ExprC{numC{n: 4}, numC{n: 5}}}, thenExpr: numC{n: 4}, elseExpr: numC{n: 5}} is : `, interp(testExprC8, topEnv))

}
