package funciones

import (
	"database/sql"
	"fmt"
	"html/template"

	_ "github.com/go-sql-driver/mysql"

	//	"github.com/gorilla/mux"
	"log"
	"github.com/gorilla/securecookie"
	"net/http"
	"strconv"
)

type Usuario struct {
	Id_usuario int
	Nombre     string
	Password   string
	Correo     string
}

var db *sql.DB
var err error

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))


func GetUserID(request *http.Request) (userID int) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			ID_string := cookieValue["Name"]
			userID, err = strconv.Atoi(ID_string)
		}
	}
	return userID
}


func setSession(userID string, response http.ResponseWriter) {
	value := map[string]string{
		//        "Id_usuario": userID,
		"Name": userID,
	}
	fmt.Println(value)
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
	user := Usuario{
		Nombre:   request.FormValue("nombre"),
		Password: request.FormValue("password"),
	}
	redirectTarget := "/usuarioInvalido"
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
}


func Logout(response http.ResponseWriter, request *http.Request) {

	clearSession(response)
	tmpl := template.Must(template.ParseFiles("publico/logout.html"))
    tmpl.Execute(response, nil)

    request.ParseForm()
}

func Autenticar(db *sql.DB, nombre string) (string, int, error) {

	var user Usuario
	stmt := "SELECT id_usuario, nombre, contraseña, email FROM usuarios WHERE nombre = ?"
	row := db.QueryRow(stmt, nombre)
	err := row.Scan(&user.Id_usuario, &user.Nombre, &user.Password, &user.Correo)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Ningún resultado de la base de datos")
		return "NOK", -1, err
	case nil:
		return user.Password, user.Id_usuario, err
	default:
		panic(err)
	}
}

func ComprobarCredenciales(user Usuario) (int, bool) {

	db, err := sql.Open("mysql", "root:root@/recetapp")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	passwordBBDD, id_usuario, err := Autenticar(db, user.Nombre)
	if err != nil {
		log.Fatal("Failed to insert into database", err)
	}
	log.Printf("Inserted row with ID of: %d\n", passwordBBDD)

	if user.Password == passwordBBDD {
		return id_usuario, true
	} else {
		return -1, false
	}
}
