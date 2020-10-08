

package main

import "fmt"




func main() {

     usuarioVerificado := "Paco"
	 password := "caca123"
	 var usuario string
	 var contraseña string

    fmt.Println("Introduce el nombre de usuario:")
    fmt.Scanln(&usuario)

    fmt.Println("Ahora Introduce la contraseña:")
    fmt.Scanln(&contraseña)

    if usuario == usuarioVerificado && password == contraseña {
    	fmt.Println("El usuario es correcto, accediendo ...")
    } else{
    	fmt.Println("Usuario no valido, vuelva a intentarlo")
    }
}