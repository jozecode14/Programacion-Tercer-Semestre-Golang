package main

import (
	"errors"
	"fmt"
	"strconv"
)
//3. Escriba un programa que mediante una función obtenga el factorial de un número entero positivo.

//OBTENEMOS EL RESULTADO DEL FACTORIAL
func obtenerFactorial(n int)(int,error){
	if n < 0{
		return n, errors.New("\terror el numero "+strconv.Itoa(n)+" debe ser mayor a 0")
	}
	fact := 1 //Inicializamos fact = 1
	for i := 1; i <= n; i++{
		fact *= i 				// fact = fact * 1 -> { fact = 1 * 1 = 1 }  {fact = 1 * 2 = 2 }  {fact = 2 * 3 = 6} ...
	}
	return fact, nil

}
//OBTENEMOS LAS OPERACIONES
func factorial (n int){
	fact := 1
	num := n - 1
	for i := 1; i <= n; i++{
		fact *= i
		if i == 2{
			fmt.Print("\t",n,"!=",n,"*",num,"*") //Imprimirá las operaciones del factorial.
			num --
		}else if i > 2{
			fmt.Print("",num,"*")
			num --
		}
	}
	fmt.Print("= ",fact) //Resultado factorial
}

func main(){
	var n int
	fmt.Print("\t---------------------\n")
	fmt.Print("\t| OBTENER FACTORIAL |\n")
	fmt.Print("\t---------------------\n")

	fmt.Println("\tIngresa un numero:")
	fmt.Print("\t->")
	fmt.Scanln(&n)
	fact,err := obtenerFactorial(n)
	if err == nil{
		fmt.Print("\t----------------------\n")
		fmt.Printf("\t| FACTORIAL de %d! = %d|\n",n,fact)
		fmt.Print("\t----------------------\n")
		factorial(n)
	} else{
		fmt.Print(err)
	}
}