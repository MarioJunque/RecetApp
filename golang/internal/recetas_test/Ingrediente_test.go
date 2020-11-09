package main

import (
	"testing"
	"database/sql"
	"github.com/MarioJunque/RecetApp/tree/master/golang/internal/recetas"
)

var db *sql.DB

func TestIngredienteBorradoConExito(t *testing.T) {

	id_ingrediente := 1
	nombreIngrediente := "guisantes"

	recetas.BorrarIngrediente(db, id_ingrediente, nombreIngrediente)
	ingredienteEliminado := recetas.ComprobarIngrediente(db, id_ingrediente, nombreIngrediente )

	expected := 1
    got := ingredienteEliminado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func TestIngredienteBorradoConExito2(t *testing.T) {

	id_ingrediente := 2
	nombreIngrediente := "salchichón"

	recetas.BorrarIngrediente(db, id_ingrediente, nombreIngrediente)
	ingredienteEliminado := recetas.ComprobarIngrediente(db, id_ingrediente, nombreIngrediente )

	expected := 1
    got := ingredienteEliminado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}


//func TestIngredienteBorradoConExito(t *testing.T) {

//	var listaIngredientes []recetas.Ingrediente

//	resultado,_:= recetas.BorrarIngrediente(listaIngredientes, 4)

//	expected := "Su ingrediente se ha eliminado con éxito"
//    got := resultado

//    if expected != got {
//        t.Errorf("got '%s' expected '%s'", got, expected)
//    }

//}

//func TestBorroElIngredienteQueQuiero (t *testing.T) {

//	var ingredientes []recetas.Ingrediente

//	ingredientes = append(ingredientes, recetas.Ingrediente{Id_ingrediente: 4, Nombre: "pepino"})

//	_, ingredientesBorrado := recetas.BorrarIngrediente(ingredientes, 4)

//	ingredientesArray := len(ingredientesBorrado)

//	expected := 0
  //  got := ingredientesArray

//    if expected != got {
//        t.Errorf("got '%v' expected '%v'", got, expected)
//    }

//}

//func TestBorroElIngredienteQueQuiero2 (t *testing.T) {

//	var ingredientes []recetas.Ingrediente

//	ingredientes = append(ingredientes, recetas.Ingrediente{Id_ingrediente: 4, Nombre: "pepino"})
//	ingredientes = append(ingredientes, recetas.Ingrediente{Id_ingrediente: 5, Nombre: "pollo"})

//	_, ingredientesBorrado := recetas.BorrarIngrediente(ingredientes, 4)

//	ingredientesArray := len(ingredientesBorrado)

//	expected := 1
//    got := ingredientesArray

//    if expected != got {
//        t.Errorf("got '%v' expected '%v'", got, expected)
//    }

//}