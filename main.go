package main

import (
	"fmt"
)

type Imagen struct {
	titulo  string
	formato string
	canales int
}

type Audio struct {
	titulo   string
	formato  string
	duracion int
}

type Video struct {
	titulo  string
	formato string
	frames  int
}

func (img *Imagen) mostrar() {
	fmt.Println("\n * IMAGEN *")
	fmt.Println("Titulo: ", img.titulo)
	fmt.Println("Formato: ", img.formato)
	fmt.Println("Canales: ", img.canales)
}

func (aud *Audio) mostrar() {
	fmt.Println("\n * AUDIO *")
	fmt.Println("Titulo: ", aud.titulo)
	fmt.Println("Formato: ", aud.formato)
	fmt.Println("Duracion: ", aud.duracion)
}

func (vid *Video) mostrar() {
	fmt.Println("\n * VIDEO *")
	fmt.Println("Titulo: ", vid.titulo)
	fmt.Println("Formato: ", vid.formato)
	fmt.Println("Frames: ", vid.frames)
}

type Multimedia interface {
	mostrar()
}

type ContenidoWeb struct {
	contenido []Multimedia
}

func (c *ContenidoWeb) mostrar() {
	for _, mult := range c.contenido {
		mult.mostrar()
	}
}

func main() {
	c := ContenidoWeb{contenido: []Multimedia{}}
  var t, f string
  var n int
  for opc := 0 ; opc != 5 ; opc += 0 {
    fmt.Println("_________________________________________")
    fmt.Println("\nMENU MULTIMEDIA - ¿Qué desea guardar?")
		fmt.Println("  1 - Imagen")
		fmt.Println("  2 - Audio")
		fmt.Println("  3 - Video")
		fmt.Println("  4 - Mostrar")
		fmt.Println("  5 - Salir")
		fmt.Print("\nOPC >> ")
		fmt.Scan(&opc)
    fmt.Println("_________________________________________")
    switch opc {
      case 1:
        fmt.Println("\nNueva Imagen")
        fmt.Print("Titulo: ")
        fmt.Scan(&t)
        fmt.Print("Formato: ")
        fmt.Scan(&f)
        fmt.Print("Canales: ")
        fmt.Scan(&n)
        nImg := Imagen{t, f, n}
        c.contenido = append(c.contenido, &nImg)
        fmt.Print("\033[H\033[2J")
      case 2:
        fmt.Println("\nNuevo Audio")
        fmt.Print("Titulo: ")
        fmt.Scan(&t)
        fmt.Print("Formato: ")
        fmt.Scan(&f)
        fmt.Print("Duración: ")
        fmt.Scan(&n)
        nAud := Audio{t, f, n}
        c.contenido = append(c.contenido, &nAud)
        fmt.Print("\033[H\033[2J")
      case 3:
        fmt.Println("\nNuevo Video")
        fmt.Print("Titulo: ")
        fmt.Scan(&t)
        fmt.Print("Formato: ")
        fmt.Scan(&f)
        fmt.Print("Frames: ")
        fmt.Scan(&n)
        nVid := Video{t, f, n}
        c.contenido = append(c.contenido, &nVid)
        fmt.Print("\033[H\033[2J")
      case 4:
        fmt.Println("\n Mostrar todo")
        c.mostrar()
        fmt.Print("\n\nPulse cualquier tecla para continuar...")
        fmt.Scanln()
        fmt.Print("\033[H\033[2J")
      case 5:
        fmt.Println("\n* Has salido *")
      default:
        fmt.Print("\033[H\033[2J")
    }
  }
}
