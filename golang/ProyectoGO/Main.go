package main

import (
	"fmt"

	"github.com/MarioJunque/RecetApp/tree/master/golang/ProyectoGO/recetas"
	"github.com/MarioJunque/RecetApp/tree/master/golang/ProyectoGO/usuario"
)

func main() {

	var opcion int
	for {
		fmt.Println("Ecriba\n 1 para registrarse como usuario\n 2 para hacer login\n 3 para salir de la aplicación")
		fmt.Scanln(&opcion)

		if opcion == 1 {

			usuario.Registro()

		} else if opcion == 2 {

			id_usuario, resultado := usuario.Login()

			if resultado == "El usuario es correcto, accediendo ..." {
				for {
					var opcion2 int
					fmt.Println("Ecriba\n 1 para añadir un ingrediente a su inventario\n 2 para generar una receta\n 3 para mostrar todas las refcetas \n 3 para cerrar sesión")
					fmt.Scanln(&opcion2)

					if opcion2 == 1 {

						recetas.BuscarIngrediente(id_usuario)
					}
					if opcion == 3 {

					}
					if opcion2 == 3 {
						break
					}
				}
			} else {
				continue
			}

		} else if opcion == 3 {

			break
		}
	}
}
