clear
echo "Actualizando la base de datos"
DATABASE=../BBDD/dump_v02.sql
echo "Indroduzca la contraseña para acceso root de MariaDB"
mysql -u root -p recetapp < $DATABASE
echo "Comprobando los test"
go test 
echo "Compilando aplicación..."
echo "Generando ejectuable para Windows OS..."
GOOS=windows GOARCH=amd64 go build -o bin/recetApp.exe
echo "Generando ejecutable para Mac OS..."
GOOS=darwin GOARCH=amd64 go build -o bin/recetApp
echo "Generando ejecutable para Linux OS..."
GOOS=linux GOARCH=amd64 go build -o bin/recetAppLinux
echo "Compilación finalizada"
