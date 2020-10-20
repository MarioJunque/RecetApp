package main

import "testing"

func TestGuardaNombre(t *testing.T) {
    expected := "El usuario es correcto, accediendo ..."
    got := IngresarCredenciales("Paco", "caca123")

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}
