package main

import "fmt"

func main(){
	//DECLARAMOS LAS VARIABLES DEL PROBLEMA
	var(
		x0 = -5
		fx0 = -120
		x1 = -3
		fx1 = -56
		x = -1
		x2 = 5
		fx2 = -40
	)
	//RESOLVEMOS LA PRIMERA FRACCION
	var fraccionSuperior1 = fx1-fx0
	var fraccionInferior1 = x1-x0

	var fraccion1Resuelta = fraccionSuperior1/fraccionInferior1

	//RESOLVEMOS LA PRIMERA MULTIPLICACION
	var multiplicacion1Resuleta = x-x0

	//RESOLVEMOS LA PRIMERA PARTE
	var primeraParteResuelta = (fraccion1Resuelta * multiplicacion1Resuleta) + fx0

	//RESOLVEMOS FRACCION SEGUNDA PARTE
	//FRACCION DE LA IZQUIEDA
	var reduccionFraccionIzq = (fx2 - fx1) / (x2 - x1)

	//FRACCION DE LA DERECHA
	var reduccionFraccionDer = (fx1 - fx0) / (x1 - x0)

	//FRACCION SUPERIOR COMPLETADA
	var fraccionSuperior = reduccionFraccionIzq - reduccionFraccionDer

	//OBTENEMOS EL VALOR DE LA SUMA DE LA PARTE INFERIOR DE LA FRACCION
	var fraccionInferior = x2 - x0

	//RESOLVEMOS LA SEGUNDA FRACCION
	var fraccion2Resuelta = fraccionSuperior / fraccionInferior

	//RESOLVEMOS LA SEGUNDA MULTIPLICACION
	var valorSegundaMultiplicacion = (x-x0)*(x-x1)

	//OBTENEMOS EL RESULTADO DE LA SEGUNDA PARTE
	var segundaParteResuelta = fraccion2Resuelta * valorSegundaMultiplicacion

	var resultado = segundaParteResuelta + primeraParteResuelta

	fmt.Println("\t-----------------------------------")
	fmt.Print("\t|\tEl valor de f(x) =",resultado,"\t  |\n")
	fmt.Println("\t-----------------------------------")
}
