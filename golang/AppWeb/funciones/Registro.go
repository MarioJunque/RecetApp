package funciones
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"html/template"
    "net/http"
)

//type Usuario struct {
//	id_usuario int
//	correo   string
//    nombre   string
//    password string
//}

type Intolerancias struct {
	gluten int
    lactosa int
    histamina int
}

//var u Usuario

func Registro(w http.ResponseWriter, r *http.Request) {

        fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("index.html")
        t.Execute(w, nil)
    } else {

		user:= Usuario {
            Nombre:   r.FormValue("nombre"),
            Password: r.FormValue("password"),
            Correo:   r.FormValue("email"),
        }    
        
        resultado := RegistrarUsuario(user)

        if resultado == "RegistroOK" {

            redirectTarget := "/internal.html"
            http.Redirect(w, r, redirectTarget, 302)
        } else {
        	t, _ := template.ParseFiles("index.html")
        	t.Execute(w, nil)
        }    
        }    

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
	stmt, err := db.Prepare("INSERT INTO usuarios (nombre, contraseña, email) VALUES(?,?,?);")
	if err != nil {
		return "NOK", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Nombre, user.Password, user.Correo)
	if err != nil {
		return "NOK", err
	}
	return "OK", err
}

func AnnadirIntoleranciasAlimenticias(id_usuario int) {
	
	var int Intolerancias
	var int_gluten string
	var int_lactosa string
	var int_histamina string

	fmt.Println("¿Eres intolerante al gluten? [s|n]")
	fmt.Scanln(&int_gluten)
	if int_gluten == "s" {
        int.gluten = 1
    } else if int_gluten == "n" {
        int.gluten = 0
    } else {
    	fmt.Println("Error. La respuesta dada es incorrecta")
    }
	fmt.Println("¿Eres intolerante a la lactosa? [s|n]")
	fmt.Scanln(&int_lactosa)
	if int_lactosa == "s" {
        int.lactosa = 1
    } else if int_lactosa == "n" {
        int.lactosa = 0
    } else {
    	fmt.Println("Error. La respuesta dada es incorrecta")
    }
    fmt.Println("¿Eres intolerante a la histamina? [s|n]")
	fmt.Scanln(&int_histamina)
	if int_histamina == "s" {
        int.histamina = 1
    } else if int_histamina == "n" {
        int.histamina = 0
    } else {
    	fmt.Println("Error. La respuesta dada es incorrecta")
    }

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	_, err = InsertIntolerancias(db, int, id_usuario)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Intolerancias añadidas correctamente")
	}
}

func InsertIntolerancias(db *sql.DB, intolerancia Intolerancias, id_usuario int) (string, error) {

	stmt, err := db.Prepare("UPDATE usuarios SET gluten = ?, lactosa = ?, histamina= ? WHERE id_usuario = ?;")
	if err != nil {
		return "NOK", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(intolerancia.gluten, intolerancia.lactosa, intolerancia.histamina, id_usuario)
	if err != nil {
		return "NOK", err
	}
	return "OK", err
}








