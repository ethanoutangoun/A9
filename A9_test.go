package main

import (
	"reflect"
	"testing"
)

func TestInterpNum(t *testing.T) {

	result := interp(NumC{N: 2}, []Binding{})
	if (result != NumV{N: 2}) {
		t.Errorf("Expected NumV(2)")
	}

}

func TestInterpId(t *testing.T) {

	result := interp(IdC{S: "x"}, []Binding{{Name: "y", Val: NumV{N: 4}},
		{Name: "x", Val: NumV{N: 2}}})
	if (result != NumV{N: 2}) {
		t.Errorf("Expected NumV(2)")
	}

}

func TestInterpString(t *testing.T) {

	result := interp(StrC{S: "hello"}, []Binding{})
	if (result != StrV{S: "hello"}) {
		t.Errorf("Expected NumV(2)")
	}

}

func TestLamC(t *testing.T) {
	result := interp(LamC{Args: []string{"x"}, Body: AppC{Fun: IdC{S: "+"}, Args: []ExprC{IdC{S: "x"}, NumC{N: 4}}}}, []Binding{})
	if !reflect.DeepEqual(result, ClosV{Args: []string{"x"}, Body: AppC{Fun: IdC{S: "+"}, Args: []ExprC{IdC{S: "x"}, NumC{N: 4}}}, Env: []Binding{}}) {

		t.Error("LamC error")
	}
}

func TestLookUpBool(t *testing.T) {
	result := interp(IdC{S: "true"}, []Binding{{Name: "true", Val: BoolV{Bool: true}}})
	if (result != BoolV{Bool: true}) {
		t.Error("BoolV error")
	}

}

func TestAdd(t *testing.T) {
	result := interp(appC{fun: idC{s: "+"}, args: []ExprC{numC{n: 3}, numC{n: 4}}}, topEnv)
	if result != 7 {
		t.Error("Add error")
	}
}

func TestSub(t *testing.T) {
	result := interp(appC{fun: idC{s: "+"}, args: []ExprC{numC{n: 3}, numC{n: 4}}}, topEnv)
	if result != 7 {
		t.Error("Add error")
	}
}
func TestDiv(t *testing.T) {
	result := interp(appC{fun: idC{s: "/"}, args: []ExprC{numC{n: 3}, numC{n: 4}}}, topEnv)
	if result != 0.75 {
		t.Error("Sub error")
	}
}
func TestMult(t *testing.T) {
	result := interp(appC{fun: idC{s: "*"}, args: []ExprC{numC{n: 3}, numC{n: 4}}}, topEnv)
	if result != 12 {
		t.Error("Mult error")
	}
}
func TestAdd(t *testing.T) {
	result := interp(appC{fun: idC{s: "+"}, args: []ExprC{numC{n: 3}, numC{n: 4}}}, topEnv)
	if result != 7 {
		t.Error("Add error")
	}
}
