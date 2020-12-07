package main

import (
    "html/template"
    "net/http"
    "fmt"
    "log"
    "github.com/MarioJunque/RecetApp/tree/master/golang/AppWeb/funciones"
)


func pantallaInicio(w http.ResponseWriter, r *http.Request) {

    fs := http.FileServer(http.Dir("publico"))
    http.Handle("/publico/", http.StripPrefix("/publico/", fs))
    tmpl := template.Must(template.ParseFiles("publico/index.html"))   
    tmpl.Execute(w, nil)

    r.ParseForm() 

}    

func login(w http.ResponseWriter, r *http.Request) {

        fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("index.html")
        t.Execute(w, nil)
    } else {

        user := funciones.Usuario {
            Nombre:   r.FormValue("nombre"),
            Password: r.FormValue("password"),
        }    
        
        usuarioValido := funciones.ComprobarCredenciales(user)
        fmt.Println(usuarioValido)

        if usuarioValido == true {

            redirectTarget := "/internal"
            http.Redirect(w, r, redirectTarget, 302)
        } else {
            http.HandleFunc("/", pantallaInicio)
        }    
        }    

    }

func internal(w http.ResponseWriter, r *http.Request) {

//    fs := http.FileServer(http.Dir("publico"))
//    http.Handle("/publico/", http.StripPrefix("/publico/", fs))
    tmpl := template.Must(template.ParseFiles("publico/internal.html"))   
    tmpl.Execute(w, nil)

    r.ParseForm() 

} 

func main() {
    
    http.HandleFunc("/", pantallaInicio)
    http.HandleFunc("/login", login)

    http.HandleFunc("/internal", internal)

    err := http.ListenAndServe(":8080", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}