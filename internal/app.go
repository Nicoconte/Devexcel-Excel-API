package internal

import (
	routes "devexcel-excel-api/internal/routes/v1"
	"devexcel-excel-api/internal/utils"
	"fmt"
	"net/http"
)

func Run() {
	http.ListenAndServe(fmt.Sprintf("%s:%s", utils.Config.Host, utils.Config.Port), routes.RoutesHandler())
}
