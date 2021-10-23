package main

import (
	"fmt"
	"reflect"
)

func main(){
	var(
		semana1,semana2 int
	)

	fmt.Print("Ingresa un numero:")
	fmt.Scanln(&semana1)

	fmt.Print("Ingresa otro numero:")
	fmt.Scanln(&semana2)

	fmt.Println(reflect.TypeOf(semana1))

	fmt.Println(semana1)
	fmt.Println(semana2)
}