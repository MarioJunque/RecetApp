package funciones

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func RecetasExistentes(db *sql.DB) ([]Receta, error) {

	var receta Receta
	var recetaExistente []Receta

	stmt := "select id_receta,nombre_receta,instrucciones from recetas;"

	rows, err := db.Query(stmt)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&receta.Id_receta, &receta.Nombre_receta, &receta.Instrucciones)
		if err != nil {
			log.Fatal(err)
		}

		recetaExistente = append(recetaExistente, receta)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return recetaExistente, err
}

func MostrarRecetasExistentes() []Receta {

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	ingredientesRecetaExistente, _ := RecetasExistentes(db)

	var recetas []Receta
	for _, a1 := range ingredientesRecetaExistente {
		recetas = append(recetas, a1)
	}
	return recetas
}

func MostrarReceta(w http.ResponseWriter, r *http.Request) {
	Recetas := MostrarRecetasExistentes()
	fmt.Println(Recetas)
	tmpl := template.Must(template.ParseFiles("publico/mostrarRecetasExistentes.html"))
	tmpl.Execute(w, Recetas)

}
