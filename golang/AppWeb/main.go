package main

import (
	"html/template"
	"log"
	"net/http"

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

func annadirIngrediente(w http.ResponseWriter, r *http.Request) {

	//    fs := http.FileServer(http.Dir("publico"))
	//    http.Handle("/publico/", http.StripPrefix("/publico/", fs))
	tmpl := template.Must(template.ParseFiles("publico/annadirIngrediente.html"))
	tmpl.Execute(w, nil)

	r.ParseForm()

}

func ingredienteNoValido(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("publico/ingredienteInvalido.html"))
	tmpl.Execute(w, nil)

	r.ParseForm()
}

func ingredienteRepetido(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("publico/ingredienteRepetido.html"))
	tmpl.Execute(w, nil)

	r.ParseForm()
}

func pantallaInicioTrasLogout(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("publico/index.html"))
	tmpl.Execute(w, nil)
	r.ParseForm()

}

func usuarioNoValido(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("publico/usuarioNoValido.html"))
	tmpl.Execute(w, nil)
	r.ParseForm()

}

func registroConExito(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("publico/registroExitoso.html"))
	tmpl.Execute(w, nil)
	r.ParseForm()

}

func ingredienteAñadidoConExito(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("publico/ingredienteAñadido.html"))
	tmpl.Execute(w, nil)
	r.ParseForm()

}

func borrarIngrediente(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("publico/borrarIngrediente.html"))
	tmpl.Execute(w, nil)
	r.ParseForm()

}

func ingredienteBorrado(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("publico/ingredienteBorrado.html"))
	tmpl.Execute(w, nil)
	r.ParseForm()

}

func noTienesIngrediente(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("publico/noIngrediente.html"))
	tmpl.Execute(w, nil)
	r.ParseForm()

}

func main() {

	http.HandleFunc("/recetapp", pantallaInicio)
	http.HandleFunc("/login", funciones.Login)
	http.HandleFunc("/usuarioInvalido", usuarioNoValido)
	http.HandleFunc("/internal", internal)
	http.HandleFunc("/registro", funciones.Registro)
	http.HandleFunc("/registroExitoso", registroConExito)
	http.HandleFunc("/AnnadirIngrediente", annadirIngrediente)
	http.HandleFunc("/ingrediente", funciones.AnnadirIngredienteAMiLista)
	http.HandleFunc("/ingredienteInvalido", ingredienteNoValido)
	http.HandleFunc("/ingredienteRepetido", ingredienteRepetido)
	http.HandleFunc("/logout", funciones.Logout)
	http.HandleFunc("/recetas", funciones.MostrarMisRecetas)
	http.HandleFunc("/recetasExistentes", funciones.MostrarReceta)
	http.HandleFunc("/misIngredientes", funciones.MostrarMisIngredientes)
	http.HandleFunc("/recetApp", pantallaInicioTrasLogout)
	http.HandleFunc("/ingredienteAñadido", ingredienteAñadidoConExito)
	http.HandleFunc("/borrarIngrediente", borrarIngrediente)
	http.HandleFunc("/ingredienteBorrado", funciones.BorrarIngrediente)
	http.HandleFunc("/ingredienteBorradoConExito", ingredienteBorrado)
	http.HandleFunc("/noIngrediente", noTienesIngrediente)
	//    http.HandleFunc("/receta",funciones.Receta)

	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
