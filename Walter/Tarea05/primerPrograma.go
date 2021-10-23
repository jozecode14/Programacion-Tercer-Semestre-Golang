package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main(){
	var(
		n int
		min int
		max int
	)
	fmt.Print("Arranque -> ")
	log.Println("")
	fmt.Println("Ingresa cuantos numeros pseudoaleatorios en un rango de [10,10,000] quiere generar:")
	fmt.Print("->")
	fmt.Scanln(&n)

	fmt.Println("Ingresa el rango minimo:")
	fmt.Print("->")
	fmt.Scanln(&min)

	fmt.Println("Ingresa el rango mayor :")
	fmt.Print("->")
	fmt.Scanln(&max)

	generarNumerosAleatorios(n,min,max)
}


func generarNumerosAleatorios(n,min,max int){
	if n < 10 || n > 10000{
		fmt.Println("El numero ingresado no esta en el rango de [10,10,000]")
	} else{
		if min < 1 {
			fmt.Println("ERROR *El rango minimo debe ser mayor a 0'Intenta de nuevo'*")
		}else{
			if max < min {
				fmt.Println("ERROR *El rango maximo debe ser mayor que el rango minimo 'Intenta de nuevo'*")
			} else {
				slice := make([]int,0)
				fmt.Print("Tiempo en imprimir los numeros ")
				log.Println("\n---NUMEROS PSEUDOALEATORIOS---")
				for i:= 1; i <= n; i++{
				rand.Seed(time.Now().UnixNano())
				randomNumbers:=min + rand.Intn(max-min)
				time.Sleep(time.Millisecond * 50) //Hacer más randoms los numeros
				fmt.Print(randomNumbers,"|")
				slice = append(slice,randomNumbers)
				}
				var(
					pares,nones,sumaPar,sumaNone int
				)
				start := time.Now()
				end1 := time.Duration(1)
				end2:= time.Duration(1)
				end3:= time.Duration(1)
				end4:= time.Duration(1)
				fmt.Print("\nTiempo de operaciones -> ")
				log.Println("\n\t---TIEMPO DE LAS OPERACIONES---")
				time.Sleep(time.Millisecond*1)
				for _, slice := range slice {
					if slice %2 == 0{
						pares ++
						end1 = time.Since(start)
						fmt.Println(pares,"+ 1 tomo:",end1)
						sumaPar += slice
						end2 = time.Since(start)
						fmt.Println( sumaPar,"+",sumaPar,"tomo:",end2)


					}else{
						nones ++
						end3 = time.Since(start)
						fmt.Println(nones,"+ 1 tomo:",end3)
						sumaNone += slice
						end4 = time.Since(start)
						fmt.Println(sumaNone,"+",sumaNone,"tomo:",end4)


					}
				}

				sumaParF := float64(sumaPar)
				sumaNoneF := float64(sumaNone)
				nF:= float64(n)

				promedioPar := sumaParF / nF
				promedioNone := sumaNoneF / nF
				fmt.Println("")
				fmt.Print("Tiempo de resultados -> ")
				log.Println("\n\t---RESULTADOS---")
				fmt.Println("")
				fmt.Println("El total de numeros pares son:",pares,"\t\t{tiempo en hacer las operaciones ->",end1,"}")
				fmt.Println("El total de numeros impares son:",nones,"\t\t{tiempo en hacer las operaciones ->",end2,"}")
				fmt.Println("El promedio de los numeros pares es de:",promedioPar,"\t{tiempo en hacer las operaciones ->",end3,"}")
				fmt.Println("El promedio de los numeros impares es de:",promedioNone,"\t{tiempo en hacer las operaciones ->",end4,"}")
				fmt.Println("")
				fmt.Print("Tiempo de salida -> ")
				log.Println("\n---PROGRAMA FINALIZADO DE MANERA EXITOSA---")
			}

		}
	}

}