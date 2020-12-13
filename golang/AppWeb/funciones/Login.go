package funciones

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
//	"github.com/gorilla/mux"
    "github.com/gorilla/securecookie"
	"log"
//	"html/template"
    "net/http"

)

type Usuario struct {
    Id_usuario int
    Nombre   string
    Password string
    Correo string
}

var db *sql.DB
var err error

var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))

func GetUserName(request *http.Request) (userName string) {
    if cookie, err := request.Cookie("session"); err == nil {
        cookieValue := make(map[string]string)
        if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
            userName = cookieValue["name"]
        }
    }
    return userName
}

func setSession(userName string, response http.ResponseWriter) {
    value := map[string]string{
        "Nombre": userName,
    }
    if encoded, err := cookieHandler.Encode("session", value); err == nil {
        cookie := &http.Cookie{
            Name:  "session",
            Value: encoded,
            Path:  "/",
        }
        http.SetCookie(response, cookie)
    }
}

func clearSession(response http.ResponseWriter) {
    cookie := &http.Cookie{
        Name:   "session",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(response, cookie)
}

func Login(response http.ResponseWriter, request *http.Request) {
	user:= Usuario {
            Nombre:   request.FormValue("nombre"),
            Password: request.FormValue("password"),
        } 
	redirectTarget := "/recetapp"
	if user.Nombre != "" && user.Password != "" {
		// .. check credentials ..
		usuarioValido := ComprobarCredenciales(user)
		
		if usuarioValido == true {
			setSession(user.Nombre, response)
			redirectTarget = "/internal"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

/*func Login(w http.ResponseWriter, r *http.Request) {

        fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("index.html")
        t.Execute(w, nil)
    } else {

//        user := funciones.Usuario {
		user:= Usuario {
            Nombre:   r.FormValue("nombre"),
            Password: r.FormValue("password"),
        }    
        
        usuarioValido := ComprobarCredenciales(user)
//        usuarioValido := funciones.ComprobarCredenciales(user)
        fmt.Println(usuarioValido)

        if usuarioValido == true {

            redirectTarget := "/internal"
            http.Redirect(w, r, redirectTarget, 302)
        } else {
        	t, _ := template.ParseFiles("index.html")
        	t.Execute(w, nil)
        }    
        }    

    }*/

func Autenticar(db *sql.DB, nombre string) (string, error){

	var user Usuario
	stmt := "SELECT id_usuario, nombre, contraseña, email FROM usuarios WHERE nombre = ?"
	row := db.QueryRow(stmt, nombre)
	err := row.Scan(&user.Id_usuario, &user.Nombre, &user.Password, &user.Correo)
	switch err {
	case sql.ErrNoRows:
  	fmt.Println("Ningún resultado de la base de datos")
  	return "NOK", err
	case nil:
  	return user.Password, err
	default:
  		panic(err)
	}
}

func ComprobarCredenciales(user Usuario) (bool) {

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	passwordBBDD, err := Autenticar(db, user.Nombre)
	if err != nil {
		log.Fatal("Failed to insert into database", err)
	}
	log.Printf("Inserted row with ID of: %d\n", passwordBBDD)



		if  user.Password == passwordBBDD {
				return true
    } else {
				return false
    }
}
