package main

import "testing"

func TestOla(t *testing.T) {
	res := Ola("Chris")
	wait := "Olá, Chris"

	if res != wait {
		t.Errorf("Resultado '%s', esperado '%s'", res, wait)
	}
}