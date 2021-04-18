package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionar(t *testing.T) {
	soma := Adiciona(2, 2)
	esperando := 4

	if soma != esperando {
		t.Errorf("esperando '%d', resultado '%d'", esperando, soma)
	}
}

func ExampleAdicionar() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	//Output: 6
}