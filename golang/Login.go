

package main

import "fmt"

func main() {
 
IngresarCredenciales()

}

func IngresarCredenciales() (usuario string, contrasenna string) {

    usuarioVerificado := "Paco"
    password := "caca123"

    fmt.Println("Introduce el nombre de usuario:")
    fmt.Scanln(&usuario)

    fmt.Println("Ahora Introduce la contraseña:")
    fmt.Scanln(&contrasenna)

    if usuario == usuarioVerificado && password == contrasenna {
    fmt.Println("El usuario es correcto, accediendo ...")
    } else{
    fmt.Println("Usuario no valido, vuelva a intentarlo")
    }

    return 
}
