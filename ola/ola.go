package main

import "fmt"

const espanhol = "espanhol"
const frances = "francês"
const japones = "japonês"
const prefixOlaPortugues = "Olá, "
const prefixOlaespanhol = "Hola, "
const prefixOlafrances = "Bonjour, "
const prefixOlajapones = "Kon'nichiwa, "

// Funções privadas em Go começam com letra minúscula.
func prefix(idioma string) string {
	switch idioma {
	case espanhol:
		return prefixOlaespanhol
	case frances:
		return prefixOlafrances
	case japones:
		return prefixOlajapones
	default:
		return prefixOlaPortugues
	}
}

// Funções públicas em Go começam com letra maiúscula.
func Ola(name string, idioma string) string {
	if name == "" {
		name = "mundo"
	}
	return prefix(idioma) + name
}

func main() {
	fmt.Println(Ola("Davy", "portugues"))
}