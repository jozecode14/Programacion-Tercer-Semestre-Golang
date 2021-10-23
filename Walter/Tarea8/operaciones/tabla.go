package operaciones

import (
	"fmt"
	"strconv"
)

func MediaPorDia(data [][]string){
	diaMedia := make ([]float64,len(data[0]))

	for i:=0; i<len(data); i++{
		for u:=0; u<len(data[0]); u++{
			value,_:=strconv.ParseFloat(data[i][u],64)
			diaMedia[u] += value /float64(len(data))
		}
	}

	for i:=0; i<len(data[0]); i++{
		fmt.Println("Dia",i+1,"promedio:",diaMedia[i])
		fmt.Printf("Menores\tMayores\n")

		for u:=0;u<len(data);u++{
			tempTemp,_:=strconv.ParseFloat(data[u][i],64)
			if tempTemp >= diaMedia[i]{
				fmt.Printf("\t%.2f",tempTemp)
			}
			if tempTemp < diaMedia[i]{
				fmt.Printf("%.2f\t",tempTemp)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
