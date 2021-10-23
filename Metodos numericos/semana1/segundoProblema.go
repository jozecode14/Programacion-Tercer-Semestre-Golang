package main

import "fmt"

func main(){
	//Declaramos las variables del problema
	var(
		x0 = -5
		fx0 = 18
		x = 3
		x1 = 5
		fx1 = -2
	)
	/*
					SEGUNDO PROBLEMA
	                  -2-18
		f1(x) = 18 + --------- (3-(-5))
					   5+5
	*/
	//Reducimos la fraccion
	var fraccionSuperior = fx1-(fx0)
	var fraccionInferior = x1-x0

	//Reducimos la multiplicación
	var reducirMulti1 = x-(x0)
	var valorDeLaFraccion = fraccionSuperior * reducirMulti1

	//Obtenemos la operacion de la fraccion
	var valorFinalFraccion = valorDeLaFraccion / fraccionInferior

	//Multiplicamos el valor de la fraccion y el primer valor
	var resultado = fx0+valorFinalFraccion

	//Imprimimos el valor de el resultado
	fmt.Println("\t-----------------------------------")
	fmt.Print("\t|\tEl valor de f(x) =",resultado,"\t  |\n")
	fmt.Println("\t-----------------------------------")
}

