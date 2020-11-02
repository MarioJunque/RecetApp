package recetas

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
//	"github.com/MarioJunque/RecetApp/tree/master/golang/usuario"
)

var db *sql.DB
var err error

type Ingrediente struct {
	id_ingrediente int
	nombre         string
}

func BuscarIngrediente(id_usuario int) {

	var ing Ingrediente

	fmt.Println("Introduce el nombre del ingrediente que desee incluir en su inventario:")
	fmt.Scanln(&ing.nombre)
	nombreIngrediente := ing.nombre

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	ing.id_ingrediente, err = ComprobarIngredienteBBDD(db, nombreIngrediente)

	if ing.id_ingrediente != -1 {

		id, err := Insert(db, ing, id_usuario)
		id_int := int(id)

		if id_int != -1 {

			fmt.Println("Se ha añadido correctamente su ingrediente")
		} else {
			log.Fatal(err)
			fmt.Println("Ocurrió un error al añadir el ingredite a su inventario")
		}
	} else {

		fmt.Println("Lo sentimos, su ingrediente no se encuentra en nuestra base de datos")
	}

}

func ComprobarIngredienteBBDD(db *sql.DB, nombreIngrediente string) (int, error) {

	var ingrediente Ingrediente
	stmt := "SELECT id_ingrediente FROM ingredientes WHERE nombre = ?"
	row := db.QueryRow(stmt, nombreIngrediente)
	err := row.Scan(&ingrediente.id_ingrediente)
	switch err {
	case sql.ErrNoRows:
		//  	fmt.Println("No está disponible este ingrediente en la base de datos")
		return -1, err
	case nil:
		return ingrediente.id_ingrediente, err
	default:
		panic(err)
	}
}

func Insert(db *sql.DB, ingrediente Ingrediente, id_usuario int) (int64, error) {

	stmt, err := db.Prepare("INSERT INTO ingrediente_usuario (id_ingredientes, id_usuarios) VALUES(?,?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(ingrediente.id_ingrediente, id_usuario)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func MostrarMisIngredientes(id_usuario int) {

//	var ingrediente Ingrediente

		db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	ingredientes, err := ObtenerMisIngredientes(db, id_usuario)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(ingredientes)
//		fmt.Println("%d\n",ingrediente.id_ingrediente)
//    	fmt.Println("%s\n",ingrediente.nombre)
	}
}

func ObtenerMisIngredientes(db *sql.DB, id_usuario int) ([]Ingrediente, error) {

	var ingrediente Ingrediente
    var ingredientes []Ingrediente
    stmt := "select id_ingredientes, nombre from ingrediente_usuario, ingredientes where ingrediente_usuario.id_ingredientes=ingredientes.id_ingrediente AND id_usuarios = ?"

	rows, err := db.Query(stmt, id_usuario)

//    rows, err := db.Query("select id_ingredientes, nombre from ingrediente_usuario, ingredientes where ingrediente_usuario.id_ingredientes=ingredientes.id_ingrediente AND id_usuarios=?;", id_usuario)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&ingrediente.id_ingrediente,&ingrediente.nombre)
        if err != nil {
            log.Fatal(err)
        }

//    ingredientes = append(ingredientes, Ingrediente{id_ingrediente: ingrediente.id_ingrediente, nombre: ingrediente.nombre})
	ingredientes = append(ingredientes, ingrediente)

    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }

    return ingredientes, err
}

