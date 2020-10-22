package usuario
import (
		"fmt"
		"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

		type Usuario struct {
     	correo   string
      	nombre   string
      	password string
   		}

//   		var usuarios[1]Usuario
   		var u Usuario

func Registro() {
		var usuario string
		var contrasenna string
		var email string

		fmt.Println("Introduzca su correo electrónico:")
		fmt.Scanln(&email)
		u.correo = email
//		usuarios[0].correo = email
//		fmt.Println(usuarios[0].correo)
//		fmt.Scanln(&email)
		fmt.Println("Introduce el nombre de usuario:")
		fmt.Scanln(&usuario)
		u.nombre = usuario
//		usuario = usuarios[0].nombre
//		fmt.Scanln(&usuario)
		fmt.Println("Introduce la contraseña:")
		fmt.Scanln(&contrasenna)
		u.password = contrasenna
//		contrasenna = usuarios[0].password
//		fmt.Scanln(&contrasenna)

//		resultado := RegistrarUsuario(email, usuario, contrasenna)
        resultado := RegistrarUsuario(u)
		fmt.Println(resultado)

}

//func RegistrarUsuario(mail string, user string, pass string) string {

//		mailVerificado := "x@x.com"
//	    usuarioVerificado := "Paco"
//		passwordVerificada := "caca123"

//		usuarios[0].correo = "x@x.com"
//  		usuarios[0].nombre = "Paco"
// 		usuarios[0].password = "caca123"

func RegistrarUsuario(user Usuario) string {

	// Create DB pool (doesn't actually connect or test connection)
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

//		if mail == mailVerificado && user == usuarioVerificado && pass == passwordVerificada {
	if id == "OK" {
				return RegistroOK
    } else {
				return RegistroNOK
    }
}

func insert(db *sql.DB, user Usuario) (string, error) {

	stmt, err := db.Prepare("INSERT INTO usuario VALUES(?,?,?)")
	if err != nil {
		return "NOK", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.nombre, user.correo, user.password)
	if err != nil {
		return "NOK", err
	}

//	return res.LastInsertId()
	return "OK", err
}

