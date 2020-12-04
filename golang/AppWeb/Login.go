package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

)

var db *sql.DB
var err error

func Autenticar(db *sql.DB, nombre string) (int, string, error){

	var user Usuario
	stmt := "SELECT id_usuario, nombre, contraseña, email FROM usuarios WHERE nombre = ?"
	row := db.QueryRow(stmt, nombre)
	err := row.Scan(&user.id_usuario, &user.nombre, &user.password, &user.correo)
	switch err {
	case sql.ErrNoRows:
  	fmt.Println("Ningún resultado de la base de datos")
  	return -1, "NOK", err
	case nil:
  	return user.id_usuario, user.password, err
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

	id_usuario, passwordBBDD, err := autenticar(db, user.nombre)
	if err != nil {
		log.Fatal("Failed to insert into database", err)
	}
	log.Printf("Inserted row with ID of: %d\n", passwordBBDD)



		if  user.password == passwordBBDD {
				return true
    } else {
				return false
    }
}