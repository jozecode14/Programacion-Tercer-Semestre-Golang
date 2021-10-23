package operaciones

import (
	"fmt"
	"strconv"
)

func ObtenerPromedios(data [][]string){

	diaPromedio := make([]float64, len(data[0]))
	for i := 0; i < len(data); i++{
		for j := 0; j < len(data[0]); j++{ //0 hasta la cantidad de datos totales
			value, _ := strconv.ParseFloat(data[i][j],64) //Pasamos los datos a float64
			diaPromedio[j] += value /float64(len(data))   //Dividimos el valor entre la cantidad total de data para sacar el promedio
		}
	}

	calor_,frio_ := diaPromedio[0],diaPromedio[0]

	var diaCalor_, diaFrio_ int

	for i:= 0; i < len(diaPromedio); i++{
		if calor_ <= diaPromedio[i] {
			calor_ = diaPromedio[i]
			diaCalor_ = i + 1
		}
		if frio_ > diaPromedio[i]{
			frio_ = diaPromedio[i]
			diaFrio_ = i + 1
		}
	}
	fmt.Println("")
	fmt.Print("2.1¿CUAL FUE EL DIA MAS CALUROSO Y CUAL EL MAS FRIO?\n")
	fmt.Print("\tPROMEDIO CALOR -> ")
	fmt.Println("-El promedio con mas calor fue de:",calor_,"el dia:",diaCalor_)
	fmt.Print("\tPROMEDIO FRIO  -> ")
	fmt.Println("-El promedio con mas frio fue de:",frio_,"el dia:",diaFrio_)
}
