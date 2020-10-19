package main

import "testing"

func TestGuardaCredenciales(t *testing.T) {
    expected := "enrique"
    got := IngresarCredenciales()

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}