package main

import (
	operaciones2 "Universidad/Walter/Tarea8/operaciones"
	"encoding/csv"
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("./data2.csv")
	if err != nil{
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	records, _:=reader.ReadAll()

	operaciones2.ObtenerPromedios(records)
	operaciones2.EvaluarTemperaturas(records)
	operaciones2.MediaPorDia(records)
}
