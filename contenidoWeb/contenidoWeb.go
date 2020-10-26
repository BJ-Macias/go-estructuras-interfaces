package contenidoWeb


import (
	"fmt"
)

type ContenidoWeb struct {
	Interfaces []Multimedia
}

func (cw *ContenidoWeb) Mostrar() {
	for _, f := range cw.Interfaces {
		f.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales int
}

func (i *Imagen) Mostrar() {
	fmt.Println("Titulo Imagen: " + i.Titulo + " Formato Imagen: " + i.Formato + " Numero de canales de imagen: ", i.Canales)
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int
}

func (a *Audio) Mostrar() {
	fmt.Println("Titulo Audio: " + a.Titulo + " Formato Audio: " + a.Formato + " Duracion en seg de Audio: ", a.Duracion)
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int
}

func (v *Video) Mostrar() {
	fmt.Println("Titulo Video: " + v.Titulo + " Formato Video: " + v.Formato + " Numero de Frames de Video: ", v.Frames)
}
