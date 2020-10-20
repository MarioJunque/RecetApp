clear
echo "Comprobando los test"
go test 
echo "Compilando aplicación..."
echo "Generando ejectuable para Windows OS..."
GOOS=windows GOARCH=amd64 go build -o recetApp.exe
echo "Generando ejecutable para Mac OS..."
GOOS=darwin GOARCH=amd64 go build -o recetApp
echo "Generando ejecutable para Linux OS..."
GOOS=linux GOARCH=amd64 go build -o recetAppLinux
echo "Compilación finalizada"
