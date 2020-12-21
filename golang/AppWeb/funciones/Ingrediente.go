package funciones

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Ingrediente struct {
	Id_ingrediente int
	Nombre         string
}

func AnnadirIngredienteAMiLista(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("publico/annadirIngrediente.html"))
		tmpl.Execute(w, nil)
	} else {

		db, err := sql.Open("mysql", "root:root@/recetapp")
		if err != nil {
			log.Fatal("Cannot open DB connection", err)
		}
		defer db.Close()

		redirectTarget := "/ingredienteInvalido"
		ingrediente := Ingrediente{
			Nombre: r.FormValue("ingrediente"),
		}
		
		ingrediente.Id_ingrediente, err = ComprobarIngredienteBBDD(db, ingrediente.Nombre)

		switch err {
		case sql.ErrNoRows:
			fmt.Println("No está disponible este ingrediente en la base de datos")
		case nil:
			id_usuario := GetUserID(r)
			tengoIngrediennte := ComprobarIngredienteUsuario(db,id_usuario,ingrediente.Id_ingrediente)

			if tengoIngrediennte == true {
			id, err := Insert(db, ingrediente, id_usuario)

			switch err {
			case sql.ErrNoRows:
				log.Fatal(id, err)
			case nil:
				redirectTarget = "/internal"
			default:
				panic(err)
				}
			} else { 
				redirectTarget = "/ingredienteRepetido"
			}
		default:
			panic(err)
		}

		http.Redirect(w, r, redirectTarget, 302)

	}

}


func ComprobarIngredienteBBDD(db *sql.DB, nombreIngrediente string) (int, error) {

	var ingrediente Ingrediente
	row := db.QueryRow("SELECT id_ingrediente FROM ingredientes WHERE nombre = ?", nombreIngrediente)
	err = row.Scan(&ingrediente.Id_ingrediente)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No está disponible este ingrediente en la base de datos")
		return -1, err
	case nil:
		return ingrediente.Id_ingrediente, err
	default:
		panic(err)
	}
}

func ComprobarIngredienteUsuario(db *sql.DB, id_usuario int, id_ingrediente int) (bool) {

	var ingrediente Ingrediente
	row := db.QueryRow("SELECT id_ingredientes FROM ingrediente_usuario WHERE id_usuarios = ? AND id_ingredientes = ?",id_usuario,id_ingrediente)
	err = row.Scan(&ingrediente.Id_ingrediente)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No está disponible este ingrediente en la base de datos")
		return true
	case nil:
		return false
	default:
		panic(err)
	}
}

func Insert(db *sql.DB, ingrediente Ingrediente, id_usuario int) (int64, error) {

	fmt.Println(id_usuario)
	stmt, err := db.Prepare("INSERT INTO ingrediente_usuario (id_ingredientes, id_usuarios) VALUES(?,?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(ingrediente.Id_ingrediente, id_usuario)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

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

func MostrarMisIngredientes(w http.ResponseWriter, r *http.Request) {
	id_usuario := GetUserID(r)
	Ingredientes := ObtenerMiIngrediente(id_usuario)
	fmt.Println(Ingredientes)
	tmpl := template.Must(template.ParseFiles("publico/mostrarMisIngredientes.html"))
	tmpl.Execute(w, Ingredientes)

}

func BorrarIngrediente(db *sql.DB, id_ingrediente int, id_usuario int) {

	stmt, err := db.Prepare("DELETE FROM ingrediente_usuario WHERE id_ingredientes = ? AND id_usuarios = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id_ingrediente, id_usuario)
	if err != nil {
		log.Fatal(err)
	}

}
