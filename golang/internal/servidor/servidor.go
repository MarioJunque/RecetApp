package servidor


import (
	
    "net/http"

	"github.com/labstack/echo/v4"
)


func iniciar () error {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Iniciando aplicación ...")
	})
	e.Logger.Fatal(e.Start(":8080"))

}


