package main

import (
	"testing"
	"github.com/MarioJunque/RecetApp/tree/master/golang/internal/usuario"
)

func TestElUsuarioAccede(t *testing.T) {
    expected := "El usuario es correcto, accediendo ..."
    got := usuario.IngresarCredenciales("Paco", "caca123")

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}

func TestElUsuarioNOAccede(t *testing.T) {
    expected := "Usuario no valido, vuelva a intentarlo"
    got := usuario.IngresarCredenciales("Pepe", "1234")

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}
