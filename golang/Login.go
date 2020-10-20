package main

import (
		"fmt"
		"database/sql"
)

var db *sql.DB
var err error

func main() {
		var usuario string
		var contrasenna string

		fmt.Println("Introduce el nombre de usuario:")
		fmt.Scanln(&usuario)
		fmt.Println("Ahora Introduce la contraseña:")
		fmt.Scanln(&contrasenna)

		resultado := IngresarCredenciales(usuario, contrasenna)
		fmt.Println(resultado)
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
		usuario := "root"
		pass := "password"
		host := "localhost"
		nombreBaseDeDatos := "agenda"
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
		if err != nil {
				return nil, err
		}
		return db, nil
}

func IngresarCredenciales(user string, password string) string {
    usuarioVerificado := "Paco"
		passwordVerificada := "caca123"

		usuarioConfirmado := "El usuario es correcto, accediendo ..."
		usuarioIncorrecto := "Usuario no valido, vuelva a intentarlo"

		if user == usuarioVerificado && password == passwordVerificada {
				return usuarioConfirmado
    } else {
				return usuarioIncorrecto
    }
}
