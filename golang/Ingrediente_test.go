package main

import (
	"testing"
	"github.com/MarioJunque/RecetApp/tree/master/golang/recetas"
	"database/sql"
)

var db *sql.DB

func TestComprobarIngredienteBBDD(t *testing.T) {

	expected := ""
    got := usuario.ComprobarIngredienteBBDD(db, "tomate")

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }

	}