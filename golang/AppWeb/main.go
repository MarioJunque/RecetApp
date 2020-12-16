package main

import (
    "html/template"
    "net/http"
    "log"
    "github.com/MarioJunque/RecetApp/tree/master/golang/AppWeb/funciones"
)



func pantallaInicio(w http.ResponseWriter, r *http.Request) {

//    fs := http.FileServer(http.Dir("publico"))
//    http.Handle("/publico/", http.StripPrefix("/publico/", fs))
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

func annadirIngrediente(w http.ResponseWriter, r *http.Request) {

//    fs := http.FileServer(http.Dir("publico"))
//    http.Handle("/publico/", http.StripPrefix("/publico/", fs))
    tmpl := template.Must(template.ParseFiles("publico/annadirIngrediente.html"))   
    tmpl.Execute(w, nil)

    r.ParseForm() 

}  


func main() {
    
    http.HandleFunc("/recetapp", pantallaInicio)
    http.HandleFunc("/login", funciones.Login)
    http.HandleFunc("/internal", internal)
    http.HandleFunc("/registro", funciones.Registro)
    http.HandleFunc("/AnnadirIngrediente", annadirIngrediente)
    http.HandleFunc("/ingrediente", funciones.AnnadirIngredienteAMiLista)
//    http.HandleFunc("/receta",funciones.Receta)

    err := http.ListenAndServe(":8080", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}