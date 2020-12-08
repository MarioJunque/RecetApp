package main

import (
    "html/template"
    "net/http"
//    "fmt"
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


func internal(w http.ResponseWriter, r *http.Request) {

//    fs := http.FileServer(http.Dir("publico"))
//    http.Handle("/publico/", http.StripPrefix("/publico/", fs))
    tmpl := template.Must(template.ParseFiles("publico/internal.html"))   
    tmpl.Execute(w, nil)

    r.ParseForm() 

} 

func main() {
    
    http.HandleFunc("/", pantallaInicio)
    http.HandleFunc("/login", funciones.Login)

    http.HandleFunc("/internal", internal)

    err := http.ListenAndServe(":8080", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}