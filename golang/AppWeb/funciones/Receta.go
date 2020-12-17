package funciones


import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"

)

type Receta struct {
	id_receta int
	nombre    string
	numeroComensales int
	instrucciones string
	nombre_receta string
	numeroIngredientes int
}

/*func Login(response http.ResponseWriter, request *http.Request) {
 
	redirectTarget := "/recetapp"
	if user.Nombre != "" && user.Password != "" {
		// .. check credentials ..
		Id_usuario, usuarioValido := ComprobarCredenciales(user)
		ID_string := strconv.Itoa(Id_usuario)
		
		if usuarioValido == true {
			setSession(ID_string, response)
			redirectTarget = "/internal"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}*/

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

/*func MostrarReceta(w http.ResponseWriter, r *http.Request){

   db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	ingredientes, err := LeerReceta(db, id_receta)

	if err != nil {
		log.Fatal(err)
	} else {
		paginaRecetaHandler()



}

func LeerReceta(db *sql.DB,id_receta int) (string,error){

 	stmt = "SELECT instrucciones from recetas WHERE id_receta = ?"
 	rows, err := db.Query(stmt)

 	defer rows.Close()

    var recetas []Receta
for rows.Next() {
    var recep Receta
    err := rows.Scan(&recep.Id_receta, &recep.NumeroComensales, &recep.Instrucciones) // check err
    users = append(recetas, recep)
}
err := rows.Err() // check err

}

func paginaRecetaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.ServeFile(w http.ResponseWriter, r *http.Request, "GenerarReceta.html"))
}*/*/