package usuario
import (
		"fmt"
)

func registro() {
		var usuario string
		var contrasenna string
		var email string

		fmt.Println("Introduzca su correo electrónico:")
		fmt.Scanln(&email)
		fmt.Println("Introduce el nombre de usuario:")
		fmt.Scanln(&usuario)
		fmt.Println("Introduce la contraseña:")
		fmt.Scanln(&contrasenna)

		resultado := RegistrarUsuario(email, usuario, contrasenna)
		fmt.Println(resultado)
}

func RegistrarUsuario(mail string, user string, pass string) string {
	return "Sus datos se han registrado con éxito"
}