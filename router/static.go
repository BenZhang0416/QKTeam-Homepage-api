package router

import (
	"api/boot/http"
)

// Change the relativePath according to your demand
func AddStaticRoute() {
	http.Router.Static("/static", "./public")
}
