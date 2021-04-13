package main

import "fmt"

const espanhol = "es"
const frances = "fr"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "


func Ola(nome , idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
		case frances:
			prefixo = prefixoOlaFrances
		case espanhol:
			prefixo = prefixoOlaEspanhol
		default:
			prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
//  Issues 5abr
// Backlog 31