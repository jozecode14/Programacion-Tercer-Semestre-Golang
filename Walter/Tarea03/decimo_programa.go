package main
/*
10.-Escriba un programa que lea dos números, uno entero y uno flotante.
Obtenga la suma. Por ejemplo si el primer número es 10 y el segundo es
5.3, el resultado es 15.3
*/
import (
	"fmt"
	"strconv"
)
func main(){
	var (
		num1 int
		num2 string
	)
	fmt.Print("Ingresa el primer numero entero:")
	fmt.Scanln(&num1)
	fmt.Print("Ingresa el segundo numero flotante:")
	fmt.Scanln(&num2)

	num,_:= strconv.ParseFloat(num2,64)

	fmt.Println(num1,"+",num2,"=",float64(num1)+num)
}