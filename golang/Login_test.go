package main

import "testing"

func TestGuardaNombre(t *testing.T) {
    expected := "enrique"
    got,_ := IngresarCredenciales()

    if expected != got {
        t.Errorf("got '%s' expected '%s'", got, expected)
    }
}