package main

import (
    "html/template"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/securecookie"
    "log"
    "github.com/MarioJunque/RecetApp/tree/master/golang/AppWeb/funciones"
)


var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
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
        "name": userName,
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

func pantallaInicio(w http.ResponseWriter, r *http.Request) {

    fs := http.FileServer(http.Dir("publico"))
    http.Handle("/publico/", http.StripPrefix("/publico/", fs))
    tmpl := template.Must(template.ParseFiles("publico/index.html"))   
    tmpl.Execute(w, nil)   
    r.ParseForm() 

}    


func internal(w http.ResponseWriter, r *http.Request) {

//    fs := http.FileServer(http.Dir("publico"))
//    http.Handle("/publico/", http.StripPrefix("/publico/", fs))
    tmpl := template.Must(template.ParseFiles("publico/internal.html"))   
    tmpl.Execute(w, nil)

    r.ParseForm() 

} 


func main() {
    
    http.HandleFunc("/recetapp", pantallaInicio)
    http.HandleFunc("/login", funciones.Login)
    http.HandleFunc("/internal", internal)
    http.HandleFunc("/registro", funciones.Registro)

    err := http.ListenAndServe(":8080", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}