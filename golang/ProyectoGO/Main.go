package main

import (
	"fmt"
	"github.com/MarioJunque/RecetApp/tree/master/golang/ProyectoGO/usuario"
)

func main() {

		var opcion int
		for{
			fmt.Println("Ecriba\n 1 para registrarse como usuario\n 2 para hacer login\n 3 para salir de la aplicación")
			fmt.Scanln(&opcion)

			if opcion == 1 {

				resultado := usuario.Registro()

				if resultado == "RegistroOK" {

				}

			} else if opcion == 2 {

				usuario.Login()	
			} else if opcion == 3{
				
				break
			}
		}
}
