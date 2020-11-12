package main
 
import (
	"fmt"
	"os"
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


var db *sql.DB
 
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
 
func setup() {

	db, err := sql.Open("mysql", "root:root@/recetapp_tests")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

_, err = db.Exec("CREATE TABLE ingredientes (id_ingrediente int(11) NOT NULL AUTO_INCREMENT,nombre varchar(50) DEFAULT NULL,PRIMARY KEY (id_ingrediente));")
_, err = db.Exec("CREATE TABLE recetas (id_receta int(11) NOT NULL AUTO_INCREMENT,numeroComensales int(11) DEFAULT NULL,instrucciones varchar(800) NOT NULL,nombre_receta varchar(200) NOT NULL, PRIMARY KEY (id_receta));")
_, err = db.Exec("CREATE TABLE usuarios (id_usuario int(11) NOT NULL AUTO_INCREMENT,nombre varchar(50) DEFAULT NULL,contraseña varchar(50) DEFAULT NULL,email varchar(50) DEFAULT NULL,gluten tinyint(1) DEFAULT 0,lactosa tinyint(1) DEFAULT 0,histamin tinyint(1) DEFAULT 0,PRIMARY KEY (id_usuario));") 
_, err = db.Exec("CREATE TABLE ingrediente_receta (id_ingrediente_receta int(11) NOT NULL AUTO_INCREMENT,id_ingredientes int(11) DEFAULT NULL,id_recetas int(11) DEFAULT NULL,PRIMARY KEY (id_ingrediente_receta), KEY id_ingrediente (id_ingredientes),KEY id_receta (id_recetas),CONSTRAINT ir1 FOREIGN KEY (id_ingredientes) REFERENCES ingredientes (id_ingrediente) ON DELETE CASCADE ON UPDATE NO ACTION,CONSTRAINT ir2 FOREIGN KEY (id_recetas) REFERENCES recetas (id_receta) ON DELETE CASCADE ON UPDATE NO ACTION);")
_, err = db.Exec("CREATE TABLE ingrediente_usuario (id_ingrediente_usuario int(11) NOT NULL AUTO_INCREMENT,id_ingredientes int(11) DEFAULT NULL,id_usuarios int(11) DEFAULT NULL,PRIMARY KEY (id_ingrediente_usuario),KEY id_ingrediente (id_ingredientes),KEY id_usuario (id_usuarios),CONSTRAINT c1 FOREIGN KEY (id_ingredientes) REFERENCES ingredientes (id_ingrediente) ON DELETE CASCADE ON UPDATE NO ACTION,CONSTRAINT c2 FOREIGN KEY (id_usuarios) REFERENCES usuarios (id_usuario) ON DELETE CASCADE ON UPDATE NO ACTION);") 
	
	if err != nil {
        log.Fatal(err)
    }

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}
 
func teardown() {

	db, err := sql.Open("mysql", "root:root@/recetapp_tests")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE IF EXISTS ingrediente_usuario, ingrediente_receta, ingredientes, recetas, usuarios;")
 
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}