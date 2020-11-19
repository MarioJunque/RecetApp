package main

import (
	"testing"
	"database/sql"
	"github.com/MarioJunque/RecetApp/tree/master/golang/internal/recetas"
	"log"
)

var db *sql.DB

func TestIngredienteBorradoConExito(t *testing.T) {

	id_ingrediente := 1
	id_usuario := 1

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	stmt := "INSERT INTO usuarios (id_usuario, nombre, contraseña, email) VALUES(1,'kike',1234,'e@e.com');"
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	stmt2 := "INSERT INTO ingredientes (id_ingrediente, nombre) VALUES(1, 'guisantes);"
	_, err = db.Exec(stmt2)
	if err != nil {
		log.Fatal(err)
	}

	stmt3 := "INSERT INTO ingrediente_usuario (id_ingredientes, id_usuarios) VALUES(1,1)"
	_, err = db.Exec(stmt3)
	if err != nil {
		log.Fatal(err)
	}

	recetas.BorrarIngrediente(db, id_ingrediente, id_usuario)
	
	stmt4 := "SELECT id_ingrediente FROM ingrediente_usuario WHERE id_ingrediente = 1"
	row := db.QueryRow(stmt4)

	var resultado int

	if row == nil {

		resultado = 0
	}

	expected := 0
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}
