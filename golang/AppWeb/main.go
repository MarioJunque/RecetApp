package main

import (
    "html/template"
    "net/http"
)

type Usuario struct {
    nombre   string
    password string
}

func main() {
    
    fs := http.FileServer(http.Dir("publico"))
    http.Handle("/publico/", http.StripPrefix("/publico/", fs))
    tmpl := template.Must(template.ParseFiles("publico/index.html"))


    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := Usuario{
            nombre:   r.FormValue("nombre"),
            password: r.FormValue("password"),
        }

        // do something with details
        _ = details

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}