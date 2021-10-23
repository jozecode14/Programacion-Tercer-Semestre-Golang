package main

import "fmt"

/*
1.- Escriba 4 funciones para resolver 4 operaciones básicas (sumar, restar, multiplicar y dividir dos números).
	Las funciones deben estar fuera del main y estar asignadas a variables dentro del main.
 */

//RETORNAMOS EL RESULTADO EN CADA UNA
func sumar(n1,n2 int)int{
	return n1+n2
}
func restar(n1,n2 int)int{
	return n1-n2
}
func multiplicar(n1,n2 int)int{
	return n1*n2
}
func dividir(n1,n2 float64)float64{
	return n1/n2
}

func main(){
	var (
		num1,num2 int
	)
	fmt.Println("\t---------------------------------")
	fmt.Println("\t|       OPERACIONES BASICAS     |")
	fmt.Println("\t---------------------------------")

	fmt.Println("\tIngresa dos numeros:")
	fmt.Print("\t->")
	fmt.Scan(&num1)
	fmt.Print("\t->")
	fmt.Scan(&num2)

	fmt.Println("\t---------------------------------")
	fmt.Println("\t|           RESULTADOS          |")
	fmt.Println("\t---------------------------------")

	fmt.Printf("\t|\t%d + %d = %d\t\t|\n",num1,num2,sumar(num1,num2))
	fmt.Printf("\t|\t%d - %d = %d\t\t|\n",num1,num2,restar(num1,num2))
	fmt.Printf("\t|\t%d * %d = %d\t\t|\n",num1,num2,multiplicar(num1,num2))
	fmt.Printf("\t|\t%.2f / %.2f = %.2f\t|\n",float64(num1),float64(num2),dividir(float64(num1),float64(num2))) //cambiamos el tipo de dato porque la division es flotante
}