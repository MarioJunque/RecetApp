package funciones

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

)

type Usuario struct {
    Id_usuario int
    Nombre   string
    Password string
    Correo string
}

var db *sql.DB
var err error

func Autenticar(db *sql.DB, nombre string) (string, error){

	var user Usuario
	stmt := "SELECT id_usuario, nombre, contraseña, email FROM usuarios WHERE nombre = ?"
	row := db.QueryRow(stmt, nombre)
	err := row.Scan(&user.Id_usuario, &user.Nombre, &user.Password, &user.Correo)
	switch err {
	case sql.ErrNoRows:
  	fmt.Println("Ningún resultado de la base de datos")
  	return "NOK", err
	case nil:
  	return user.Password, err
	default:
  		panic(err)
	}
}

func ComprobarCredenciales(user Usuario) (bool) {

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	passwordBBDD, err := Autenticar(db, user.Nombre)
	if err != nil {
		log.Fatal("Failed to insert into database", err)
	}
	log.Printf("Inserted row with ID of: %d\n", passwordBBDD)



		if  user.Password == passwordBBDD {
				return true
    } else {
				return false
    }
}