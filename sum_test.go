package main

import "testing" 

func TestSum(t *testing.T){
	x := 5+5
    if x != 10 {
		t.Error("Opa!!!!! 5 + 5 não é igual a 10, obtive", x)
	   }
	}