package main

import (
	"testing"
	"database/sql"
	"github.com/MarioJunque/RecetApp/tree/master/golang/internal/recetas"
)

var db *sql.DB

func TestIngredienteBorradoConExito(t *testing.T) {

	expected := "Su ingrediente se ha eliminado con éxito"
    got := recetas.BorrarIngrediente(db, 4)

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }

	}