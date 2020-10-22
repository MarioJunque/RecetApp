package recetas
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/MarioJunque/RecetApp/tree/master/golang/ProyectoGO/usuario"
)

var db *sql.DB
var err error

type Ingediente struct {
	id_ingrediente int
	nombre string
}

func buscarIngrediente(user Usuario) {

	var ing Ingrediente

	fmt.Println("Introduce el nombre del ingrediente que desee incluir en su inventario:")
	fmt.Scanln(&ing.nombre)
	nombreIngrediente := ing.nombre

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	ing.id_ingrediente, err := comprobarIngredienteBBDD(db, nombreIngrediente)

	if ing.id_ingrediente != nil {

	id, err := insert(db, ing, user)

		if id != nil {

			fmt.Println("Se ha añadido correctamente su ingrediente")
		} else {

			fmt.Println("Ocurrió un error al añadir el ingredite a su inventario")
		}
	} else {

		fmt.Println("Lo sentimos, su ingrediente no se encuentra en nuestra base de datos")
	}

}

func comprobarIngredienteBBDD(db *sql.DB, nombreIngrediente string) (int, error){

	var id int

	stmt, err := db.Prepare("SELECT id_ingrediente FROM ingredientes WHERE nombre = ?")
	if err != nil {
		return "NOK", err
	}
	defer stmt.Close()

	rows, err := stmt.Query(nombreIngrediente)
	if err != nil {
		return "NOK", err
	}
	defer rows.Close()
	for rows.Next() {
        if err := rows.Scan(&id); err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%s is %d\n", id, nombreIngrediente)
	}
	if err := rows.Err(); err != nil {
        log.Fatal(err)
}

//	return res.LastInsertId()
	return id, err
}

func insert(db *sql.DB, ingrediente Ingrediente, user Usuario) (int, error) {

	stmt, err := db.Prepare("INSERT INTO ingrediente_usuario (id_ingrediente, id_usuario) VALUES(?,?)")
	if err != nil {
		return "NOK", err
	}
	defer stmt.Close()

	res, err = stmt.Exec(ingrediente.id_ingrediente, user.id_usuario)
	if err != nil {
		return "NOK", err
	}

	return res.LastInsertId()
}