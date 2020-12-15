package funciones

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
    "net/http"
    "html/template"
)

type Receta struct {
	Id_receta int
    NumeroComensales int
	Instrucciones string
}