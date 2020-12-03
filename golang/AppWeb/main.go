package main

import (
    "html/template"
    "net/http"
    "fmt"
    "log"
)

type Usuario struct {
    nombre   string
    password string
}


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

        user := Usuario{
            nombre:   r.FormValue("nombre"),
            password: r.FormValue("password"),
        }
        fmt.Println(user.nombre)
        fmt.Println(user.password)      

    }
}

func main() {
    
    http.HandleFunc("/", pantallaInicio)
    http.HandleFunc("/login", login)

    err := http.ListenAndServe(":8080", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}