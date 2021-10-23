package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)
/*
4 Ahora haga que la página web reciba un número entero y obtenga su cuadrado, de tal forma que si ejecutamos el
	servidor como go run ici_stats.go 4, el resultado de la página web será algo como: El cuadrado de 4 es 16.
 */
func main() {
	num := os.Args[1]
	http.HandleFunc("/hello",func(rw http.ResponseWriter,req *http.Request){
		num1, _ := strconv.Atoi(num)
		rw.Write([]byte(fmt.Sprintf("El cuadrado de %d es: %d",num1,num1*num1)))
	})
	http.ListenAndServe(":8080",nil)
}
