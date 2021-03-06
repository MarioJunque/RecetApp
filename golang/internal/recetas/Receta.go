package recetas

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
//	"github.com/MarioJunque/RecetApp/tree/master/golang/internal/usuario"
)

//var db *sql.DB
//var err error

type Receta struct {
	id_receta int
	nombre    string
	numeroComensales int
	instrucciones string
	nombre_receta string
	numeroIngredientes int
}


func MostrarMisRecetas(id_usuario int) {

	misRecetas := ObtenerMisRecetas(id_usuario)
    fmt.Println(misRecetas)

}

func IngredientesUsuarioReceta(db *sql.DB, id_usuario int) ([]Receta, error) {

	var receta Receta
	var ingredientesUsuario []Receta


	stmt := "select id_receta, count(ingrediente_receta.id_ingredientes) from recetas left outer join ingrediente_receta on recetas.id_receta = ingrediente_receta.id_recetas join ingrediente_usuario on ingrediente_receta.id_ingredientes = ingrediente_usuario.id_ingredientes and ingrediente_usuario.id_usuarios = ? group by recetas.id_receta;"

	rows, err := db.Query(stmt, id_usuario)

    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&receta.id_receta,&receta.numeroIngredientes)
        if err != nil {
            log.Fatal(err)
        }

	ingredientesUsuario = append(ingredientesUsuario, receta)

    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    } 
    return ingredientesUsuario, err

}

func NumeroIngredientesReceta(db *sql.DB) ([]Receta, error) {

	var receta Receta
	var ingredientesReceta []Receta

	stmt := "select id_receta, count(id_ingredientes) from recetas left outer join ingrediente_receta on recetas.id_receta = ingrediente_receta.id_recetas group by id_receta;"

	rows, err := db.Query(stmt)

    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&receta.id_receta,&receta.numeroIngredientes)
        if err != nil {
            log.Fatal(err)
        }

	ingredientesReceta = append(ingredientesReceta, receta)

    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }

    return ingredientesReceta, err
}

func ObtenerMisRecetas(id_usuario int) ([]Receta){


    db, err := sql.Open("mysql", "root:root@/recetapp")
    if err != nil {
        log.Fatal("Cannot open DB connection", err)
    }
    defer db.Close()

	ingredientesUsuario, _ := IngredientesUsuarioReceta(db, id_usuario)
	ingredientesReceta, _ := NumeroIngredientesReceta(db)
    var misRecetas []Receta

    for _, a1 := range ingredientesReceta {
        for _, a2 := range ingredientesUsuario {
            if a1 == a2 {
                misRecetas = append(misRecetas, a2)
            }
        }
    }

	return misRecetas

    }
