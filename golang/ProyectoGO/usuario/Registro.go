package usuario
import (
		"fmt"
)

		type Usuario struct {
     	correo   string
      	nombre   string
      	password string
   		}

   		var usuarios[1]Usuario

func Registro() {
		var usuario string
		var contrasenna string
		var email string

		fmt.Println("Introduzca su correo electrónico:")
		fmt.Scanln(&email)
		usuarios[0].correo = email
//		fmt.Println(usuarios[0].correo)
//		fmt.Scanln(&email)
		fmt.Println("Introduce el nombre de usuario:")
		fmt.Scanln(&usuarios[0].nombre)
		usuario = usuarios[0].nombre
//		fmt.Scanln(&usuario)
		fmt.Println("Introduce la contraseña:")
		fmt.Scanln(&usuarios[0].password)
		contrasenna = usuarios[0].password
//		fmt.Scanln(&contrasenna)

		resultado := RegistrarUsuario(email, usuario, contrasenna)
		fmt.Println(resultado)
}

func RegistrarUsuario(mail string, user string, pass string) string {

//		mailVerificado := "x@x.com"
//	    usuarioVerificado := "Paco"
//		passwordVerificada := "caca123"

//		usuarios[0].correo = "x@x.com"
//  		usuarios[0].nombre = "Paco"
// 		usuarios[0].password = "caca123"


		RegistroOK := "Sus datos se han registrado con éxito"
		RegistroNOK := "Vaya, parece que hubo un problema. Vuelva a intentarlo"

//		if mail == mailVerificado && user == usuarioVerificado && pass == passwordVerificada {
	if mail == usuarios[0].correo && user == usuarios[0].nombre && pass == usuarios[0].password {
				return RegistroOK
    } else {
				return RegistroNOK
    }
}