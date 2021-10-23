package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)
//FUNCION QUE VALIDA LAS ENTRADAS, RETORNA UN MENSAJE Y UN BOOLEANO
func validarEntradas1(rangoMin float64,rangoMax float64,datos int,diaInicio int,diaFinal int) (string, bool){
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
	tempMin,_ := strconv.Atoi(rangoMin)
	tempMax,_ := strconv.Atoi(rangoMax)

	//MANDAMOS A LLAMAR LA FUNCION QUE VALIDA LAS ENTRADAS
	mensaje, err := validarEntradas1(rangoMinF, rangoMaxF, datosI, diaInicioI, diaFinalI)

	if err != false{
		fmt.Println(mensaje)
		fmt.Printf("\t\tDATOS INGRESADOS:\n\t\tRango minimo: \t%.0f°c\n\t\tRango maximo: \t%.0f°c\n\t\tTotal de datos: %d\n\t\tMes inicial: \t%d\n\t\tMes final: \t%d\n\n",rangoMinF,rangoMaxF,datosI,diaInicioI,diaFinalI)
		fmt.Println("\t---------------------------------\n\t+  INGRESANDO LOS DATOS AL CSV  +\n\t---------------------------------")
		rand.Seed(time.Now().UnixNano())

		file,err := os.Create("./data2.csv")
		if err != nil{
			fmt.Println(err)
		}

		writer := csv.NewWriter(file)
		data := make([][]string, datosI)/**/
		for i := 0; i < datosI; i++{
			data[i] = make([]string, diaFinalI-diaInicioI+1)
			for j := 0; j <=(diaFinalI - diaInicioI); j++{
				data[i][j] = fmt.Sprintf("%.2f", float64(rand.Intn(tempMax - tempMin +1)+tempMin)+rand.Float64())
			}
		}
		//Writer de datos en csv
		err = writer.WriteAll(data)
		if err != nil{
			fmt.Println(err)
		}

		//Imprimir
		fmt.Printf("#")
		for i := 0; i < (diaFinalI-diaInicioI+1); i++{
			fmt.Printf("\t%d", i+1)
		}
		fmt.Println()
		for i2 := 0; i2 < datosI; i2 ++{
			fmt.Printf("%d\t", i2+1)
			for u := 0; u <=(diaFinalI - diaInicioI); u++{
				fmt.Printf("%v\t", data[i2][u])
			}
			fmt.Println()
		}

	} else{
		fmt.Println(mensaje)
	}
}