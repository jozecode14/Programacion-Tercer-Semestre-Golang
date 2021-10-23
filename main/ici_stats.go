package main

import (
  "errors"
  f "file/file"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main(){
  if len(os.Args) < 2{ //Verificamos la longitud de nuestros argumentos
    err := errors.New("ERROR")
    fmt.Println(err,"Te hacen falta mas parámetros, la estructura es la siguiente: >main g n 2 10 datos.dat ")
  }
  if len(os.Args) >= 5{
    flagSelected := strings.ToLower(os.Args[1]) //Convertimos a minuscula
    numbersQuantity, _ := strconv.Atoi(os.Args[2]) //Convertimos a entero
    buttonLimit, _ := strconv.Atoi(os.Args[3])
    maxLimit, _ := strconv.Atoi(os.Args[4])
    file := os.Args[5] //El argumento será tipo string

    //VALIDAMOS ENTRADAS
    if numbersQuantity < 1{
      err := errors.New("ERROR > La cantidad de numeros a generar debe de ser mayor a 0")
      fmt.Println(err)
    }
    /////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    if buttonLimit > maxLimit {
      err := errors.New("ERROR > El limite inferior debe ser menor que limite superior Li < Ls")
      fmt.Println(err)
    }
    /////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    if flagSelected != "g" {
      err := errors.New("ERROR > La bandera seleccionada no existe")
      fmt.Println(err)
    }
    fmt.Print("\t+----------------------------------+\n\t| ENTRADAS VALIDADAS CORRECTAMENTE |\n\t+---------------" +
      "-------------------+\n")

    f.CreatingDat(numbersQuantity, buttonLimit, maxLimit, file) //Con esta funcion crearemos el archivo".dat"

  } else{
    flagSelected := os.Args[1]
    newFile := os.Args[2]
    f.MoreOptions(flagSelected,newFile)
  }
}
