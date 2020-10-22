package usuario
import (
		"fmt"
)

		type Usuario struct {
     	correo   string
      	nombre   string
      	password string
   		}

//   		var usuarios[1]Usuario
   		var u Usuario

func Registro() Usuario {
		var usuario string
		var contrasenna string
		var email string

		fmt.Println("Introduzca su correo electrónico:")
		fmt.Scanln(&email)
		u.correo := email
//		usuarios[0].correo = email
//		fmt.Println(usuarios[0].correo)
//		fmt.Scanln(&email)
		fmt.Println("Introduce el nombre de usuario:")
		fmt.Scanln(&usuario)
		u.nombre := usuario
//		usuario = usuarios[0].nombre
//		fmt.Scanln(&usuario)
		fmt.Println("Introduce la contraseña:")
		fmt.Scanln(&contrasenna)
		u.password := contrasenna
//		contrasenna = usuarios[0].password
//		fmt.Scanln(&contrasenna)

		resultado := RegistrarUsuario(email, usuario, contrasenna)
		fmt.Println(resultado)
		return usuarios[0]
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
	if mail == u.correo && user == u.nombre && pass == u.password {
				return RegistroOK
    } else {
				return RegistroNOK
    }
}