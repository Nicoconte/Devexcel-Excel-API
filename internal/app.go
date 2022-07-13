package internal

import (
	routes "devexcel-excel-api/internal/routes/v1"
	"fmt"
	"net/http"
)

func Run() {
	http.ListenAndServe(fmt.Sprintf(":%s", Config.Port), routes.RoutesHandler())
}
