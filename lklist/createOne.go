package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var i int = 1

	for true {
		fmt.Printf("\tMuestra: %d.\n",i)
		i++
		rand.Seed(time.Now().UnixNano())
		numRand1 := 1 + rand.Intn(5 - 1) //El valor aleatorio para agregar el nuevo número
		numRand2 := 1 + rand.Intn(10 - 1)//El valor aleatorio para eliminar la cola

		lista := linkedList{}   //Se crea una nueva lista
		nodo1 := &node{data: 1} //Se crean los nodos "&" toma la dirección del nodo
		nodo2 := &node{data: 2}
		nodo3 := &node{data: 3}
		nodo4 := &node{data: 4}

		lista.anteponer(nodo1)
		lista.anteponer(nodo2)
		lista.anteponer(nodo3)

		lista.imprimirLista()//Lista con números del 1 al 3
		log.Print()
		fmt.Println()
		/////////////////////////////////////////////////////////////////
		time.Sleep(time.Duration(numRand1)*time.Second)
		lista.anteponer(nodo4)
		lista.imprimirLista()//Lista con números del 1 al 4
		log.Print()
		fmt.Println()
		/////////////////////////////////////////////////////////////////
		time.Sleep(time.Duration(numRand2)*time.Second)
		lista.eliminarValor(4) //Se elimina el primer valor ingresado
		lista.imprimirLista()//Lista con números del 2 al 4
		log.Print()
		fmt.Println()
		/////////////////////////////////////////////////////////////////
		fmt.Printf("El tiempo transcurrido de atención fue de %d segundos\n\n",numRand1 + numRand2)
	}

}

type node struct { //Inicialización del nodo
	data int
	//El asterisco indica el puntero de la dirección de memoria
	siguiente *node //"siguiente" es la dirección del siguiente nodo
}

type linkedList struct { //Inicialización de la lista vinculada.
	head   *node //Se guarda la dirección la pila. La lista enlazada no guarda toda la lista, solamente la pila
	longitud int //Longitud de la lista
}

func (l *linkedList) anteponer (n *node) { //"l" es el recibidor de la dirección de la lista para poder modificar la lista | "n" guardará la dirección para agregar un nodo hasta el final
	shiftHead := l.head //recorre la pila a la segunda posición
	l.head = n //Poner el nuevo nodo en la pila (n es el nuevo nodo)
	l.head.siguiente = shiftHead //La pila apunta al antiguo nodo que ahora es el segundo y sigue el orden de la fila
	l.longitud++ //Se incrementa la longitud porque se agregó un nuevo nodo
}

func (l linkedList) imprimirLista(){ //Imprimir los datos de cada nodo
	ImprimirPila := l.head //Se guarda la pila en una variable para despues imprimirla
	for l.longitud  != 0{ //Se utiliza un ciclo para imprimir descendientemente la longitud de la lista hasta que sea 0
		fmt.Printf("[%d] ", ImprimirPila.data) //Se imprimen los datos de la pila
		ImprimirPila = ImprimirPila.siguiente //Se actualiza la pila con el siguiente nodo para imprimirlo
		l.longitud-- //la longitud se va reduciendo en 1 para que pueda llegar a 0
	}
	fmt.Println(" ")//Salto de linea después de que se imprima todo
}

func (l *linkedList) eliminarValor (value int){ //"value" guarda el número que se va a eliminar
	//Condiciones especiales
	if l.longitud == 0 { //Si la longitud de la lista es 0 sale de la función
		return
	}

	if l.head.data == value { //Si se elimina la pila, el segundo dato de la lista, pasará a ser la nueva pila
		l.head = l.head.siguiente
		l.longitud--
		return
	}

	previoAEliminar := l.head //guardamos la pila en una variable
	for previoAEliminar.siguiente.data != value {
		if previoAEliminar.siguiente.siguiente == nil { //Si solo queda un elemento en la lista que está apuntando a nil, entonces no hay otro valor que aparezca en la lista y sale de la función
			return
		}
		//El ciclo se ejecuta hasta que se encuentre un nodo que corresponde con el número que se quiere eliminar
		previoAEliminar = previoAEliminar.siguiente //Se actualiza la lista
	}
	previoAEliminar.siguiente = previoAEliminar.siguiente.siguiente //Se actualiza la lista para que la pila salte el número que ase acaba de eliminar
	l.longitud-- //Se reduce la lista en uno porque acabamos de eliminar un elemento
}