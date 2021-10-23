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
				log.Println("\t<---LOS NUMEROS PSEUDOALEATORIOS SON--->")
				var (
					pares int
					impares int
				)
				for i := 0; i < n; i++{
					rand.Seed(time.Now().UnixNano())
					randomNumbers:=li + rand.Intn(ls-li)
					time.Sleep(time.Millisecond * 100) //Hacer más randoms los numeros
					fmt.Print("|",randomNumbers)
					if randomNumbers % 2  == 0{
						pares ++
					}else {
						impares ++
					}
				}
				fmt.Println("")
				log.Printf("\t<--RESULTADOS-->\n")
				fmt.Println("\tSe genero un total de",n,"numeros pseudoaleatorios")
				fmt.Println("\tPares:",pares)
				fmt.Println("\tImpares:",impares)
				log.Println("\t¡LA SALIDA FUE EXITOSA!")
			}
		}
	}

}