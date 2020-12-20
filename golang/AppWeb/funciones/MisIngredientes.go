package funciones

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func MisIngredientes(db *sql.DB, id_usuario int) ([]Ingrediente, error) {

	var ingrediente Ingrediente
	var miIngrediente []Ingrediente

	stmt := "select id_ingredientes, nombre from ingrediente_usuario, ingredientes where ingrediente_usuario.id_ingredientes=ingredientes.id_ingrediente AND id_usuarios = ?;"

	rows, err := db.Query(stmt, id_usuario)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&ingrediente.Id_ingrediente, &ingrediente.Nombre)
		if err != nil {
			log.Fatal(err)
		}

		miIngrediente = append(miIngrediente, ingrediente)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return miIngrediente, err
}

func ObtenerMiIngrediente(id_usuario int) []Ingrediente {

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	miIngrediente, _ := MisIngredientes(db, id_usuario)

	var ingredientes []Ingrediente
	for _, a1 := range miIngrediente {
		ingredientes = append(ingredientes, a1)
	}
	return ingredientes
}

func MostrarMiIngrediente(w http.ResponseWriter, r *http.Request) {
	id_usuario := GetUserID(r)
	Ingredientes := ObtenerMiIngrediente(id_usuario)
	fmt.Println(Ingredientes)
	tmpl := template.Must(template.ParseFiles("publico/mostrarMisIngredientes.html"))
	tmpl.Execute(w, Ingredientes)

}
