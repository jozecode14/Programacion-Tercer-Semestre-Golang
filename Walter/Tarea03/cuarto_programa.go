package main

//4.-Desplegar la suma de 2 numeros enteros positivos con su resultado.

import "fmt"

func sumaNumeros(num1, num2 int){
	fmt.Printf("%d + %d = %d",num1,num2,num1+num2)

}
func main(){
	num1 := 5
	num2 := 20
	sumaNumeros(num1,num2)
}
