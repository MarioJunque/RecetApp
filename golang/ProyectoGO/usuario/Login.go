package usuario
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

)

var db *sql.DB
var err error

func Login() (int, string) {
		var usuario string
		var contrasenna string

		fmt.Println("Introduce el nombre de usuario:")
		fmt.Scanln(&usuario)
		u.nombre = usuario
		fmt.Println("Ahora Introduce la contraseña:")
		fmt.Scanln(&contrasenna)
		u.password = contrasenna

		id_usuario,resultado := IngresarCredenciales(u)
		fmt.Println(resultado)
		return id_usuario, resultado
}


func autenticar(db *sql.DB, nombre string) (int, string, error){

	var user Usuario
	stmt := "SELECT * FROM usuarios WHERE nombre = ?"
	row := db.QueryRow(stmt, nombre)
	err := row.Scan(&user.id_usuario, &user.nombre, &user.password, &user.correo)
	switch err {
	case sql.ErrNoRows:
  	fmt.Println("Ningún resultado de la base de datos")
  	return -1, "NOK", err
	case nil:
  	return user.id_usuario, user.password, err
	default:
  		panic(err)
	}
}

	func IngresarCredenciales(user Usuario) (int, string) {

		usuarioConfirmado := "El usuario es correcto, accediendo ..."
		usuarioIncorrecto := "Usuario no valido, vuelva a intentarlo"	

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	id_usuario, passwordBBDD, err := autenticar(db, user.nombre)
	if err != nil {
		log.Fatal("Failed to insert into database", err)
	}
	log.Printf("Inserted row with ID of: %d\n", passwordBBDD)



		if  user.password == passwordBBDD {
				return id_usuario, usuarioConfirmado
    } else {
				return -1, usuarioIncorrecto
    }
}