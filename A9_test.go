package main

import (
	"testing"
	"reflect"
)


func TestInterpNum(t *testing.T) {
	
	result := interp(numC{n : 2}, []Binding{})

	if (result.(float64) != 2){
		t.Errorf("interp num failed")
	}

}

func TestInterpId(t *testing.T) {
	
	result := interp(idC{s : "x"}, []Binding{{name: "y", val:4 },
											 {name: "x", val: 2}})
	if (result.(int) !=2){
		t.Errorf("idC interp failed")
	}

}


func TestInterpString(t *testing.T) {
	
	result := interp(strC{str : "hello"}, []Binding{})
	if (result != "hello"){
		t.Errorf("interp string")
	}

}



func TestLamC(t *testing.T){
	result := interp(lamC{args : []string{"x"}, body : appC{fun: idC{s: "+"}, args: []ExprC{idC{s: "x"}, numC{n: 4}}}}, []Binding{})
	if !reflect.DeepEqual(result, closV{args : []string{"x"}, body : appC{fun: idC{s: "+"}, args: []ExprC{idC{s: "x"}, numC{n: 4}}}, env : []Binding{}}){
	
		t.Error("LamC error")
	}
}


func TestLookUpBool(t *testing.T){
	result := interp(idC{s : "true"}, []Binding{{name : "true", val : true}})
	if (result!= true){
		t.Error("BoolV error")
	}

}
