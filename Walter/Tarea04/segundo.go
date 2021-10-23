package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main(){
	var (
		n int
		li int
		ls int
	)
	fmt.Println("Ingresa la cantidad de numeros pseudoaleatorios a generar:")
	fmt.Print("->")
	fmt.Scanln(&n)

	if n < 100 {
		fmt.Println("La cantidad debe ser mayor a 100")
	} else{
		fmt.Println("-Ingresa limite inferior: ")
		fmt.Print("->")
		fmt.Scanln(&li)
		if li < 1 {
			fmt.Println("*Limite inferior debe ser mayor a cero*")
		}else{
			fmt.Println("-Ingresa limite superior: ")
			fmt.Print("->")
			fmt.Scanln(&ls)
			if ls < li {
				fmt.Println("*Ls debe ser mayor a li*")
			}else{
				var v int
				fmt.Println("Ingresa el valor de v en un rango de [0,",n,"]:")
				fmt.Scanln(&v)
				if v <0 && v > n{
					fmt.Println("Debes de ingresar un numero que este en el rango de [0,",n,"]")
				}else{
					log.Println("\t<---LOS NUMEROS PSEUDOALEATORIOS SON--->")
					var numvIguales int
					var multiplos int
					for i := 0; i < n; i++{
						rand.Seed(time.Now().UnixNano())
						randomNumbers := li + rand.Intn(ls-li)
						time.Sleep(time.Millisecond * 100) //Hacer más randoms los numeros
						fmt.Print("|", randomNumbers)
						if v == randomNumbers{
							numvIguales ++
						}
						if randomNumbers % v == 0 {
							multiplos ++
						}
					}
					fmt.Println("")
					log.Println("-\t--RESULTADOS---")
					fmt.Println("\nLos numeros repetidos de ",v," son:",numvIguales)
					fmt.Println("Los multiplos de ",v," son:",multiplos)
				}
			}
		}
	}
}