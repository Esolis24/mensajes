package mensajes

import "fmt"

func Hola() {
	fmt.Println("Hola mensajes")
}

const mensajes = "Hola desde mi constante"

func funcionPrivada(){
	fmt.Println("Hola desde la funcion privada")
}
func Imprimir(){
	fmt.Println(mensajes)
	funcionPrivada()
}