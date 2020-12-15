package funciones

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
    "net/http"
    "html/template"
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

	redirectTarget := "/recetapp"
	ingrediente:= Ingrediente {
            Nombre:   r.FormValue("ingrediente"),
        }    
        
    ingrediente.Id_ingrediente, err = ComprobarIngredienteBBDD(db, ingrediente.Nombre)

    switch err {
	case sql.ErrNoRows:
		fmt.Println("No está disponible este ingrediente en la base de datos")
	case nil:
		id_usuario := GetUserID(r)
		id, err := Insert(db, ingrediente, id_usuario)

		switch err {
		case sql.ErrNoRows:
			log.Fatal(id, err)
		case nil:
			redirectTarget = "/internal"
		default:
			panic(err)
		}

	default:
		panic(err)
	}

	http.Redirect(w, r, redirectTarget, 302)

	}

}


//func BuscarIngrediente(id_usuario int) {

func ComprobarIngredienteBBDD(db *sql.DB, nombreIngrediente string) (int, error) {

	var ingrediente Ingrediente
	stmt := "SELECT id_ingrediente FROM ingredientes WHERE nombre = ?"
	fmt.Println(nombreIngrediente)
	fmt.Println(stmt)
	row := db.QueryRow(stmt,nombreIngrediente)
	fmt.Println(row)
	err := row.Scan(&ingrediente.Id_ingrediente)
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

func Insert(db *sql.DB, ingrediente Ingrediente, id_usuario int) (int64, error) {

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
        err := rows.Scan(&ingrediente.Id_ingrediente,&ingrediente.Nombre)
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

func ComprobarIngrediente(db *sql.DB, id_ingrediente int, nombreIngrediente string) int {

	return 1
}

