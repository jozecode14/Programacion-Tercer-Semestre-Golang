package main

import "fmt"

func main(){
	//DECLARAMOS LAS VARIABLES DEL PROBLEMA
	var(
		x0 = -3
		fx0 = 59
		x1 = -1
		fx1 = 11
		x = 2
		x2 = 4
		fx2 = 66
	)
	/*
								DESARROLLAMOS EL TERCER PROBLEMA

									  66 - 11    11 - 59
 				  11-59				 --------- - ---------
	f2(x)= 59+ ---------- (2-(-3)) +  4 - (- 1)  -1 - (-3) (2 - (-3))(2-(-1)
				 -1-(-3)			 ----------------------
											 4-(-3)
	 */
	//RESOLVEMOS LA PRIMERA FRACCION
	var fraccionSuperior1 = fx1-fx0 //11-59 = -48
	var fraccionInferior1 = x1-x0   //-1-(-3) = 2

	var fraccion1Resuelta = fraccionSuperior1/fraccionInferior1 // =-24

	//RESOLVEMOS LA PRIMERA MULTIPLICACION
	var multiplicacion1Resuelta = x-x0 //(2-(-3)) = 5

	//RESOLVEMOS LA PRIMERA PARTE
	var primeraParteResuelta = (fraccion1Resuelta * multiplicacion1Resuelta) + fx0 // (-24 * 5)+59 = -61

	//RESOLVEMOS FRACCION SEGUNDA PARTE
	//FRACCION DE LA IZQUIEDA
	var reduccionFraccionIzq = (fx2 - fx1) / (x2 - x1) //66 - 11 / 4 -(-1) = 11

	//FRACCION DE LA DERECHA
	var reduccionFraccionDer = (fx1 - fx0) / (x1 - x0) //11-59 / -1- (-3) = -24

	//FRACCION SUPERIOR COMPLETADA
	var fraccionSuperior = reduccionFraccionIzq - reduccionFraccionDer // 11 + 24 = 35

	//OBTENEMOS EL VALOR DE LA SUMA DE LA PARTE INFERIOR DE LA FRACCION
	var fraccionInferior = x2 - x0 // 4-(-3) = 7

	//RESOLVEMOS LA SEGUNDA FRACCION
	var fraccion2Resuelta = fraccionSuperior / fraccionInferior // 35 / 7 = 5

	//RESOLVEMOS LA SEGUNDA MULTIPLICACION
	var valorSegundaMultiplicacion = (x-x0)*(x-x1)

	//OBTENEMOS EL RESULTADO DE LA SEGUNDA PARTE
	var segundaParteResuelta = fraccion2Resuelta * valorSegundaMultiplicacion

	var resultado = segundaParteResuelta + primeraParteResuelta

	fmt.Println("\t-----------------------------------")
	fmt.Print("\t|\tEl valor de f(x) =",resultado,"\t  |\n")
	fmt.Println("\t-----------------------------------")
}
