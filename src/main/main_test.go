package main

import "testing"

func Test_Conta(t *testing.T) {
	resultado := efetuaConta()

	if resultado != 5333333323333.449219 {
		t.Error("Valor inválido para a conta!")
	}
}
