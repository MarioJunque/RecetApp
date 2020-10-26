package usuario
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

		type Usuario struct {
		id_usuario int	
     	correo   string
      	nombre   string
      	password string
   		}

   		var u Usuario

func Registro() {
		var usuario string
		var contrasenna string
		var email string

		fmt.Println("Introduzca su correo electrónico:")
		fmt.Scanln(&email)
		u.correo = email
		fmt.Println("Introduce el nombre de usuario:")
		fmt.Scanln(&usuario)
		u.nombre = usuario
		fmt.Println("Introduce la contraseña:")
		fmt.Scanln(&contrasenna)
		u.password = contrasenna
        resultado := RegistrarUsuario(u)
		fmt.Println(resultado)

}

func RegistrarUsuario(user Usuario) string {

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	id, err := insert(db, user)
	if err != nil {
		log.Fatal("Failed to insert into database", err)
	}
	log.Printf("Inserted row with ID of: %d\n", id)

		RegistroOK := "Sus datos se han registrado con éxito"
		RegistroNOK := "Vaya, parece que hubo un problema. Vuelva a intentarlo"

	if id == "OK" {
				return RegistroOK
    } else {
				return RegistroNOK
    }
}

func insert(db *sql.DB, user Usuario) (string, error) {

	stmt, err := db.Prepare("INSERT INTO usuarios (nombre, contraseña, email) VALUES(?,?,?)")
	if err != nil {
		return "NOK", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.nombre, user.password, user.correo)
	if err != nil {
		return "NOK", err
	}
	return "OK", err
}

