package main

//5.-Desplegar el cuadrado de un numero entero cualquiera.

import "fmt"

func elevarCuadrado(num int){
	fmt.Printf("%d ^ %d = %d",num,num,num*num)
}

func main(){
	var num int =15
	elevarCuadrado(num)
}
