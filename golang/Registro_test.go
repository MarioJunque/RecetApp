package main

import (
	"testing"
	"github.com/MarioJunque/RecetApp/tree/master/golang/usuario"
)

func TestGuardaEnBBDD(t *testing.T) {
    expected := "Sus datos se han registrado con éxito"
    got := usuario.RegistrarUsuario("x@x.com", "Paco", "caca123")

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}