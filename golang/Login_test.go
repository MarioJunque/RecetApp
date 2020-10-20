package main

import (
	"testing"
	"github.com/MarioJunque/RecetApp/tree/master/golang/login"
)

func TestGuardaNombre(t *testing.T) {
    expected := "El usuario es correcto, accediendo ..."
    got := login.IngresarCredenciales("Paco", "caca123")

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}
