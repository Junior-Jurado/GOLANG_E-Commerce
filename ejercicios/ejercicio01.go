package ejercicios

import (
	"fmt"
	"strconv"
)

func MiFuncion(parametro string) (valor1 int, valor2 string) {
	var text string 
	numero, err := strconv.Atoi(parametro); 
	if err != nil {
		fmt.Println("Error:", err)
	} else if numero > 100 {
		text = parametro + " es mayor a 100"
	} else {
		text = parametro + " es menor a 100"
	}

	return numero, text


}