package main

import (
	file2"Universidad/Walter/examen/file"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//CONDICIONAMOS LOS ARGS
	if len(os.Args) < 2{
		err := errors.New("ERROR Falta de argumentos")
		fmt.Println(err)
		os.Exit(2)
	}

	if len(os.Args) > 3 {
		//Declaramos las variables que irian si el usuario ingresa más de 3 datos
		var selectFlag = strings.ToUpper(os.Args[1])
		var numbersQuantuty,_= strconv.Atoi(os.Args[2])
		var buttonLimit,_ = strconv.Atoi(os.Args[3])
		var topLimit,_ = strconv.Atoi(os.Args[4])
		var file = os.Args[5]

		if numbersQuantuty < 1 {
			error := errors.New("ERROR La cantidad de numeros no puede ser menor a 1")
			fmt.Println(error)
		}
		if buttonLimit > topLimit {
			error := errors.New("ERROR El limite inferior debe ser mas grande que el limite superior")
			fmt.Println(error)
		}
		if selectFlag != "G"{
			error := errors.New("ERROR La opción es incorrecta")
			fmt.Println(error)
		}
		//TERMINAMOS DE VALIDAR ENTRADAS
		fmt.Println("Entraste al main")
		file2.WritingData(numbersQuantuty,buttonLimit,topLimit,file,selectFlag) //Mandamos a llamar la funcion WritingData() Para crear el archivo .dat

	}
		/*if len(os.Args) == 3{
			var optionFlag = os.Args[2]
			var file = os.Args[3]
		}*/

	if len(os.Args) < 2{
		err := errors.New("Insuficientes argumentos")
		fmt.Printf("%v\nLos argumentos deben ser n g li ls file\n",err)
		os.Exit(2)
	}

}