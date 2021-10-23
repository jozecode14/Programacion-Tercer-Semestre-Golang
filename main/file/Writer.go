package file

import (
  "encoding/csv"
  "errors"
  "fmt"
  "math/rand"
  "os"
  "time"
)

func CreatingDat(quantityData int, firstLimit int, lastLimit int, fileName string){

    data := make([][]string, 1)
    rand.Seed(time.Now().UnixNano())

    file, err := os.Create(fileName+".dat")
      if err != nil{
        fmt.Println(err)
      }

      data[0] = make([]string, quantityData)
      for i:=0; i<quantityData; i++{
        value := fmt.Sprintf("%d",rand.Intn(lastLimit - firstLimit) + firstLimit)
        data[0][i] = value
      }

      writer := csv.NewWriter(file)

      err = writer.WriteAll(data)
      if err != nil{
        err = errors.New("ERROR > el archivo no se pudo crear")
        fmt.Println(err)
      }
  fmt.Printf("\t> Teclea el comando 'type %s.dat' para abrir el archivo",fileName)

}
