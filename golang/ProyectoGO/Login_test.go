package main

import (
	"testing"
	"github.com/MarioJunque/RecetApp/tree/master/golang/ProyectoGO/login"
)

func TestElUsuarioAccede(t *testing.T) {
    expected := "El usuario es correcto, accediendo ..."
    got := login.IngresarCredenciales("Paco", "caca123")

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}

func TestElUsuarioNOAccede(t *testing.T) {
    expected := "Usuario no valido, vuelva a intentarlo"
    got := login.IngresarCredenciales("Pepe", "1234")

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}
