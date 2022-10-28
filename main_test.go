package main

import "testing"

func TestMain(t *testing.T) {
	sum := suma(5, 2)

	if sum != 7 {
		t.Error("la suma es incorrecta")
	} else {
		t.Log("test finalizado correctamente")
	}
}
