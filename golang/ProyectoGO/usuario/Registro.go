package usuario
import (
		"fmt"
)

func Registro() string {
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
		return resultado
}

func RegistrarUsuario(mail string, user string, pass string) string {

		mailVerificado := "x@x.com"
	    usuarioVerificado := "Paco"
		passwordVerificada := "caca123"

		RegistroOK := "Sus datos se han registrado con éxito"
		RegistroNOK := "Vaya, parece que hubo un problema. Vuelva a intentarlo"

		if mail == mailVerificado && user == usuarioVerificado && pass == passwordVerificada {
				return RegistroOK
    } else {
				return RegistroNOK
    }
}