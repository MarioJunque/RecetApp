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
    gluten int
    lactosa int
    histamina int
}

var u Usuario

func Registro() {
	var usuario string
	var contrasenna string
	var email string
	var int_gluten string
	var int_lactosa string
	var int_histamina string

	fmt.Println("Introduzca su correo electrónico:")
	fmt.Scanln(&email)
	u.correo = email
	fmt.Println("Introduce el nombre de usuario:")
	fmt.Scanln(&usuario)
	u.nombre = usuario
	fmt.Println("Introduce la contraseña:")
	fmt.Scanln(&contrasenna)
	u.password = contrasenna
	fmt.Println("¿Eres intolerante al gluten? [s|n]")
	fmt.Scanln(&int_gluten)
	if int_gluten == "s" {
        u.gluten = 1
    } else if int_gluten == "n" {
        u.gluten = 0
    } else {
    	fmt.Println("Error. La respuesta dada es incorrecta")
    }
    fmt.Println("¿Eres intolerante a la lactosa? [s|n]")
	fmt.Scanln(&int_lactosa)
	if int_lactosa == "s" {
        u.lactosa = 1
    } else if int_lactosa == "n" {
        u.lactosa = 0
    } else {
    	fmt.Println("Error. La respuesta dada es incorrecta")
    }
    fmt.Println("¿Eres intolerante a la histamina? [s|n]")
	fmt.Scanln(&int_histamina)
	if int_histamina == "s" {
        u.histamina = 1
    } else if int_histamina == "n" {
        u.histamina = 0
    } else {
    	fmt.Println("Error. La respuesta dada es incorrecta")
    }
	
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
	insercion := fmt.Sprintf("INSERT INTO usuarios (nombre, contraseña, email, gluten, lactosa, histamina) VALUES(%s,%s,%s,%d,%d,%d);", user.nombre, user.password, user.correo, user.gluten, user.lactosa, user.histamina)
	stmt, err := db.Prepare(insercion)
	if err != nil {
		return "NOK", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.nombre, user.password, user.correo, user.gluten, user.lactosa, user.histamina)
	if err != nil {
		return "NOK", err
	}
	return "OK", err
}