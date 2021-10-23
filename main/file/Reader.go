package file

import (
  "encoding/csv"
  "fmt"
  "os"
  "strconv"
)

func MoreOptions(newFlag string, fileName string){
  //Abrimos el archivo
  file, err := os.Open(fileName)
  if err != nil{ //Validamos el error
      fmt.Println(err)
  }

  //Leemos el archivo
  reader := csv.NewReader(file)
  data, err := reader.ReadAll()
  if err != nil{
      fmt.Println(err)
  }

  if newFlag == "M" {  //M = Número mayor
    numeroMayor, _ := strconv.Atoi(data[0][0])
    for i:=0; i<len(data[0]); i++{
      value, _ := strconv.Atoi(data[0][i])
      if numeroMayor < value {
         numeroMayor = value
      }
    }
    fmt.Print("\t+--------------+\n\t| NUMERO MAYOR |\n\t+--------------+\n")
    fmt.Println("\tEl numero mayor es:",numeroMayor)

  }
  if newFlag == "m" {
    numeroMenor, _ := strconv.Atoi(data[0][0])
    for i:=0; i<len(data[0]); i++{
      value, _ := strconv.Atoi(data[0][i])
      if numeroMenor > value {
          numeroMenor = value
      }
    }
    fmt.Print("\t+--------------+\n\t| NUMERO MENOR |\n\t+--------------+\n")
    fmt.Println("\tEl numero menor es:",numeroMenor)
  }
  if newFlag == "p" {
    var promedio float64
    for i:=0; i<len(data[0]); i++{
      value, _ := strconv.Atoi(data[0][i])
      promedio += float64(value)/float64(len(data[0]))
    }
    fmt.Print("\t+-----------+\n\t| PROMEDIO |\n\t+-----------+\n")
    fmt.Printf("\tPromedio: %.2f",promedio)
  }
  if newFlag == "a" {
    fmt.Print("\t+-------+\n\t| DATOS |\n\t+-------+\n")
    fmt.Print("\t")
    for i:=0; i<len(data[0]); i++{
      fmt.Printf("%s ",data[0][i])
    }
  }
}
