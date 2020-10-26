package main

import (
	"fmt"

	"./contenidoWeb"
)

func main() {

	var op, canales, duracion, frames int
	var tituloI, formatoI, tituloA, formatoA, tituloV, formatoV string
	fmt.Println(op)

	i01 := contenidoWeb.Imagen{}
	a01 := contenidoWeb.Audio{}
	v01 := contenidoWeb.Video{}
	cw := contenidoWeb.ContenidoWeb{}

	for op != 5 {

		fmt.Println("Elija una opcion")
		fmt.Println("1. Capturar Imagen")
		fmt.Println("2. Capturar Audio")
		fmt.Println("3. Capturar Video")
		fmt.Println("4. Mostrar")
		fmt.Println("5. Salir")
		fmt.Scan(&op)

		if op == 1 {
			fmt.Print("Ingrese Titulo de imagen: ")
			fmt.Scan(&tituloI)
			fmt.Print("Ingrese Formato de imagen: ")
			fmt.Scan(&formatoI)
			fmt.Print("Ingrese num de Canales de imagen: ")
			fmt.Scan(&canales)
			i01 = contenidoWeb.Imagen{Titulo: tituloI, Formato: formatoI, Canales: canales}
		}
		if op == 2 {
			fmt.Print("Ingrese Titulo de audio: ")
			fmt.Scan(&tituloA)
			fmt.Print("Ingrese Formato de audio: ")
			fmt.Scan(&formatoA)
			fmt.Print("Ingrese duracion en seg de audio: ")
			fmt.Scan(&duracion)
			a01 = contenidoWeb.Audio{Titulo: tituloA, Formato: formatoA, Duracion: duracion}
		}
		if op == 3 {
			fmt.Print("Ingrese Titulo de video: ")
			fmt.Scan(&tituloV)
			fmt.Print("Ingrese Formato de video: ")
			fmt.Scan(&formatoV)
			fmt.Print("Ingrese num de frames de video: ")
			fmt.Scan(&frames)
			v01 = contenidoWeb.Video{Titulo: tituloV, Formato: formatoV, Frames: frames}
		}
		if op == 4 {
			cw = contenidoWeb.ContenidoWeb{Interfaces: []contenidoWeb.Multimedia{&i01, &a01, &v01},}
			cw.Mostrar()
		}
		if op == 5 {
			fmt.Println("Saliendo")
		}
	}
}
