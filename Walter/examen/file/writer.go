package file

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func WritingData(totalQuantity int, buttonRange int, topRange int, fileName string,flag string){
	if flag == "G"{
		rand.Seed(time.Now().UnixNano())
		data := make([][]string, 1)

		file, err := os.Create(fileName+".dat")
		if err != nil{
			fmt.Println(err)
		}

		data[0] = make([]string, totalQuantity)
		for i:=0; i<totalQuantity; i++{
			value := fmt.Sprintf("%d",rand.Intn(buttonRange - topRange) + buttonRange)
			data[0][i] = value
		}

		writer := csv.NewWriter(file)

		err = writer.WriteAll(data)
		if err != nil{
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error")
	}
}