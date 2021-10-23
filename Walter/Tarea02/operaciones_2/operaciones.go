package operaciones_2 //Nombre de la carpeta donde se ubica el paquete

import (
	"fmt"
	"os"
	"strconv"
)
func Sumar()  {
	A, _ := strconv.Atoi(os.Args[1])
	B, _ := strconv.Atoi(os.Args[2])

	fmt.Println(A+B)
}

