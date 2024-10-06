package main

import (
	"fmt"

	"github.com/Junior-Jurado/GOLANG_E-Commerce/variables"
)

func main() {
	estado, texto := variables.ConviertoTexto(1)

	fmt.Println(estado)
	fmt.Println(texto)
}
