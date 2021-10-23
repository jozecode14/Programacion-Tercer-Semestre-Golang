package main

import "fmt"

func main(){
	//Declaramos las variables del problema
	var(
		x0 = -4
		fx0 = -35
		x = -1
		x1 = 3
		fx1 = 21
	)
	/*
				PRIMER PROBLEMA
				  21 -(-35)
	f1(x) = -35 + --------- (-1-(-4)
				  3 -(-4)
	 */
	//Reducimos la fraccion
	var fraccionSuperior = fx1-(fx0) //	21 -(-35) = 56
	var fraccionInferior = x1-x0 	 // 1-(-4) = +3

	//Reducimos la multiplicación
	var reducirMulti1 = x-(x0) // (-1-(-4) = 3
	var valorDeLaFraccion = fraccionSuperior * reducirMulti1 // =56 * 3 = 168

	//Obtenemos la operacion de la fraccion
	var valorFinalFraccion = valorDeLaFraccion / fraccionInferior// =168/3 = 56

	//Multiplicamos el valor de la fraccion y el primer valor
	var resultado = fx0+valorFinalFraccion // -35 + 56 = 11

	//Imprimimos el valor de el resultado
	fmt.Println("\t-----------------------------------")
	fmt.Print("\t|\tEl valor de f(x) =",resultado,"\t  |\n")
	fmt.Println("\t-----------------------------------")
}
