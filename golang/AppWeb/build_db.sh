clear
DATABASEROUTE=../../BBDD/dump_v03.sql
echo "Eliminando base de datos ..."
mysqladmin -u root -p drop recetapp
echo "Generando nueva base de datos"
mysqladmin -u root -p create recetapp
mysqldump -u root -p recetapp < $DATABASEROUTE
echo "Iniciando página web ..."
go run main.go
open -a safari localhost:8080/recetapp 
