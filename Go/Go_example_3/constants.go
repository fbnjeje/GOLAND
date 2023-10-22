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

	var num int = 2

	var myFloat float32 = 6.2
	println(myFloat + float32(num))

	var myBool bool = true
	println(myBool)

	//variable creada e inicialozada de  manera abreviada

	myString := "texto"
	print(myString)

	// Myvar = 2 error
}

//esto es un comentario JAJAJAJAJ

/*

esto es un comentario largo
*/
