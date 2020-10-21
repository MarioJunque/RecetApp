package main

import (
	"fmt"
	"github.com/MarioJunque/RecetApp/tree/master/golang/ProyectoGO/usuario"
)

func main() {

		var opcion int
		for{
			fmt.Println("Ecriba\n 1 para registrarse como usuario\n 2 para hacer login")
			fmt.Scanln(&opcion)

			if opcion == 1 {

				usuario.Registro()

			} else if opcion == 2 {

				usuario.Login()	
			} else {
				break
			}
		}
}
