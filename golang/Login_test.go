package main

import "testing"

func TestGuardaValores(t *testing.T) {
    expected := ""
    got := IngresarCredenciales()

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}