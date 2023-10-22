package main

import (
	"fmt"
	"reflect"
)

func main() {

	var Myvar string = "Esto se una cadena de texto"
	fmt.Println("hola como estas ", Myvar)

	Myvar = "Esto se modifica Gente"
	fmt.Println(Myvar)

	fmt.Println(reflect.TypeOf(Myvar))

	// Myvar = 2 error
}

//esto es un comentario JAJAJAJAJ

/*

esto es un comentario largo
*/
