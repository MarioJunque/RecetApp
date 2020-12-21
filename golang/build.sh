clear
echo "Actualizando la base de datos ..."
DATABASE="../../BBDD/dump_v02.sql"
echo $DATABASE
echo "Eliminando base de datos ..."
mysqladmin -u root -p drop recetapp
echo "Generando nueva base de datos"
mariadb -u root -p recetapp < $DATABASE
echo "Iniciando página web ..."
go run main.go
open -a safari