package main

import (
	"net/http"

	"github.com/jw81/go-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
