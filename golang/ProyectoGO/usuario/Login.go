package usuario
import (
		"fmt"
				"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB
var err error

//type Usuario struct {
//correo   string
//nombre   string
//password string
//}
//var u Usuario


func Login() string {
		var usuario string
		var contrasenna string


		fmt.Println("Introduce el nombre de usuario:")
		fmt.Scanln(&usuario)
		u.nombre = usuario
		fmt.Println("Ahora Introduce la contraseña:")
		fmt.Scanln(&contrasenna)
		u.password = contrasenna

		resultado := IngresarCredenciales(u)
		fmt.Println(resultado)
		return resultado
}


func autenticar(db *sql.DB, nombre string) (string, error){

	var pass string
	stmt, err := db.Prepare("SELECT contraseña FROM usuario WHERE nombre = ?")
	if err != nil {
		return "NOK", err
	}
	defer stmt.Close()

	rows, err := stmt.Query(nombre)
	if err != nil {
		return "NOK", err
	}
	defer rows.Close()
	for rows.Next() {
        if err := rows.Scan(&pass); err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%s is %d\n", pass, nombre)
	}
	if err := rows.Err(); err != nil {
        log.Fatal(err)
}

//	return res.LastInsertId()
	return pass, err
}

//func IngresarCredenciales(user string, password string) string {
//        usuarioVerificado := "Paco"
//		passwordVerificada := "caca123"


	func IngresarCredenciales(user Usuario) string {

		usuarioConfirmado := "El usuario es correcto, accediendo ..."
		usuarioIncorrecto := "Usuario no valido, vuelva a intentarlo"

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	passwordBBDD, err := autenticar(db, user.nombre)
	if err != nil {
		log.Fatal("Failed to insert into database", err)
	}
	log.Printf("Inserted row with ID of: %d\n", passwordBBDD)



		if  user.password == passwordBBDD {
				return usuarioConfirmado
    } else {
				return usuarioIncorrecto
    }
}