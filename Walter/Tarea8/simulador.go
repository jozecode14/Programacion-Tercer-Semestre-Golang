package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"text/template/parse"
	"time"
)
//FUNCION QUE VALIDA LAS ENTRADAS, RETORNA UN MENSAJE Y UN BOOLEANO
func validarEntradas(rangoMin float64,rangoMax float64,datos int,diaInicio int,diaFinal int) (string, bool){
	var mensaje string
	var err bool

	if rangoMin < rangoMax {
		if datos >=1{
			if diaInicio < diaFinal{
				mensaje = "\t----------------------------------------\n\t+  ENTRADAS REGISTRADAS CORRECTAMENTE  +\n\t----------------------------------------"
				err = true
				return mensaje, err
			} else{
				mensaje = "ERROR - El dia de inicio debe ser menor a el dia final."
				err = false
				return mensaje, err
			}
		}else{
			mensaje = "ERROR - Se debe de registrar al menos un dato."
			err = false
			return mensaje, err
		}
	}else{
		mensaje = "ERROR - El rango maximo debe ser menor a el rango minimo."
		err = false
		return mensaje, err
	}
}

func main() {
	fmt.Println(variable)
	//ENTRADAS EN LA TERMINAL
	rangoMin := os.Args[1]  //Temperatura minima
	rangoMax := os.Args[2]  //Temperatura maxima
	datos := os.Args[3]     //Cantidad de datos
	diaInicio := os.Args[4] //Dia de inicio - empieza a leer datos
	diaFinal := os.Args[5]  //Dia final - termina de leer datos

	//CONVERSION A TIPO FLOTANTE
	rangoMinF, _ := strconv.ParseFloat(rangoMin, 64)
	rangoMaxF, _ := strconv.ParseFloat(rangoMax, 64)

	//CONVERSION A TIPO DE DATO ENTERO
	datosI, _ := strconv.Atoi(datos)
	diaInicioI, _ := strconv.Atoi(diaInicio)
	diaFinalI, _ := strconv.Atoi(diaFinal)

	//MANDAMOS A LLAMAR LA FUNCION QUE VALIDA LAS ENTRADAS
	mensaje, err := validarEntradas(rangoMinF, rangoMaxF, datosI, diaInicioI, diaFinalI)

	//PARA VALIDAR LAS ENTRADAS LA SALIDA BOOLEANA DEBE SER TRUE PARA ENTRAR DESARROLLAR LAS TEMPERATURAS
	if err != false{
		file, err := os.Create("./data.csv") //Creamos el archivo csv.

		if err != nil{ //Validamos la entrada de el file
			fmt.Println(err)
		}

		var slice [] string //Creamos un slice de tipo string

		fmt.Println(mensaje)
		fmt.Printf("\t\tDATOS INGRESADOS:\n\t\tRango minimo: \t%.0f°c\n\t\tRango maximo: \t%.0f°c\n\t\tTotal de datos: %d\n\t\tDia de inicio: \t%d\n\t\tDia final: \t%d\n\n",rangoMinF,rangoMaxF,datosI,diaInicioI,diaFinalI)
		fmt.Print("\t#\t")
		rand.Seed(time.Now().UnixNano()) //Creamos la semilla
		for i := diaInicioI; i <= diaFinalI; i++ {
			fmt.Printf("%d\t",i) //Creamos los numeros de los dias ingresados
		}
		fmt.Println()
		for i := 1; i <= datosI; i++ {
			fmt.Print("\t",i,"\t")
			for j := 0; j <= diaFinalI - diaInicioI; j++ {
				tempRandom := rangoMinF + rand.NormFloat64() + (rangoMaxF-rangoMinF) - 0.5 //Creamos el numeros pseudoaleatorio entre las dos cantidades
				time.Sleep(time.Millisecond * 100) //Para numeros mas randoms creamos un sleep de 100 milisegundos
				fmt.Printf("%.2f\t", tempRandom) //Imprimimos los numeros randoms
				dato:= fmt.Sprintf("%.2f",tempRandom) //Convertimos el numero float64 a string
				slice = append(slice,dato) //El numero lo ingresamos a el slice con un append.
			}
			fmt.Println()

		}
		slice2 := strings.Join(slice, "\n") //El slice lo pasamos a tipo string, " "- significa que va hacer un espacio en blanco

		writer := csv.NewWriter(file) //Creamos un writer y ingresamos el archivo creado anterior mente llamado "file"

		var data = [][]string{ //Esta es la sintaxis para ingresar datos a un archivo .csv
			{slice2},
		}

		err = writer.WriteAll(data) //Validamos el error del writer.
		if err != nil {
			fmt.Println(err)
		}
	}else{
		fmt.Println(mensaje) //Si err == false, imprime el error de la entrada.
	}
}