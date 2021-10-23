package main

import "fmt"

// 2.-Escriba una funcion variadica que reciba n numeros enteros y obtenga su promedio.
func obtenerPromedio(n...int)float64{
	var (
		//Inicializamos las variables
		suma int
		promedio float64
		cantidad int
	)
	for _,item := range n{
		suma += item
	}
	cantidad = len(n) //len() -> Cantidad de elementos en el slice
	promedio = float64(suma) / float64(cantidad) //Convertimos al mismo tipo de dato para dividir
	return promedio
}

func main() {
	var numeros [] int = [] int {10,9,8,7,6,5} //Numeros ingresados

	promedio := obtenerPromedio(numeros...)

	fmt.Printf("El promedio de los numeros %v, es: %.2f",numeros, promedio) //Imprimimos el promedio
}