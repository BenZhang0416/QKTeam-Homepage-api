package Route

import (
	"api/boot/http"
	"api/controller/version"
)

func init()  {
	AddStaticRoute()
	AddRouter()
}

// Add your route here
func AddRouter() {
	http.Router.GET("/", version.Version)
}