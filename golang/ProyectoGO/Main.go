package main

import (
	"fmt"
	"github.com/MarioJunque/RecetApp/tree/master/golang/ProyectoGO/usuario"
	"github.com/MarioJunque/RecetApp/tree/master/golang/ProyectoGO/recetas"
)


func main() {

		var opcion int
		for{
			fmt.Println("Ecriba\n 1 para registrarse como usuario\n 2 para hacer login\n 3 para salir de la aplicación")
			fmt.Scanln(&opcion)

			if opcion == 1 {

				usuario.Registro()

			} else if opcion == 2 {

				id_usuario, resultado := usuario.Login()	

					if resultado == "El usuario es correcto, accediendo ..." {

					var opcion2 int
					fmt.Println("Ecriba\n 1 para añadir los ingredientes que tiene a su disposión\n 2 para generar una receta\n 3 para cerrar sesión")
					fmt.Scanln(&opcion2)

					if opcion2 == 1 {

						recetas.BuscarIngrediente(id_usuario)
					}	
				}

			} else if opcion == 3{
				
				break
			}
		}
}
