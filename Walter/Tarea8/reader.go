package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)
//Funcion que checa el error
func checkError(msg string, err error){
	if err != nil{
		log.Fatal(msg,err)
	}
}

func main(){
	//Abriremos el archivo data.csv
	Openfile, err := os.Open("./data.csv")
	checkError("Error al abrir el archivo data.csv",err)

	fileData,err := csv.NewReader(Openfile).ReadAll()
	checkError("Error al leer el archivo data.csv",err)

	for i, value := range fileData{
		fmt.Printf(" i = %d value = %s\n",i,value)
	}
}