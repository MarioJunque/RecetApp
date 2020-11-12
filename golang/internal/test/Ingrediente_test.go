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
	nombreIngrediente := "guisantes"

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	recetas.BorrarIngrediente(db, id_ingrediente, nombreIngrediente)
	ingredienteEliminado := recetas.ComprobarIngrediente(db, id_ingrediente, nombreIngrediente )

	expected := 1
    got := ingredienteEliminado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}
