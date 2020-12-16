package funciones

/*import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
    "net/http"
    "html/template"
)

type Receta struct {
	Id_receta int
    NumeroComensales int
	Instrucciones string
}


func MostrarReceta(w http.ResponseWriter, r *http.Request){

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
}*/