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

	var num int = 5

	// var myFloat float32 = 6.2
	// println(myFloat + float32(num))

	var myBool bool = true
	println(myBool)

	//variable creada e inicialozada de  manera abreviada

	myString := "texto"
	print(myString)

	// const myConst = "esto es una constante"

	// control de flujo jaja

	num = 2
	myString = "hola"

	if num == 2 && myString == "hola" {
		fmt.Print("el valor es 2")
	} else if num == 5 {
		fmt.Print("el valor es 5")
	} else {
		fmt.Print("el valor no es 2")
	}

	// Myvar = 2 error
}

//esto es un comentario JAJAJAJAJ

/*

esto es un comentario largo
*/
