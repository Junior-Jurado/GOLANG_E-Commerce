package main

import (
	"fmt"

	"github.com/Junior-Jurado/GOLANG_E-Commerce/ejercicios"
	// "runtime"
	// "github.com/Junior-Jurado/GOLANG_E-Commerce/variables"
)

func main() {
	num, str := ejercicios.MiFuncion("102")

	fmt.Println(num)
	fmt.Println(str)

	// estado, texto := variables.ConviertoTexto(1)

	// fmt.Println(estado)
	// fmt.Println(texto)

	// if os:= runtime.GOOS; os == "Linux." {
	// 	fmt.Println("Esto no es Windows, es: ", os)
	// } else {
	// 	fmt.Println("Esto es Windows, ", os)
	// }

	// switch os:= runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }
}
