package main

import (
	"fmt"
	"net/http"
	"os"
)
/*
3.- Hacer la página web básica que reciba el nombre desde la línea de comando de la terminal,
	es decir que se ejecute como: go run ici_stats.go Alex y al cargar la página aparezca: Hola, Alex
 */
func main(){
	palabra := os.Args[1] //Ingresamos el nombre en la consola
	http.HandleFunc("/hello",func(rw http.ResponseWriter,req *http.Request){
		palabra = string(palabra)
		rw.Write([]byte(fmt.Sprintf("Hola, %s",palabra)))
	})
	http.ListenAndServe(":8080",nil)
}
