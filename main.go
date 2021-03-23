package main

import (
	"api/boot/http"
	"api/boot/logger"
	mw "api/boot/middleware"
	"api/boot/orm"
	_ "api/config"
	_ "api/router"
)

func _init() {
	http.DefaultMiddleWares = mw.DefaultMiddleWares
	logger.InitLog()
	orm.InitOrm()
	http.InitHttp()
}

func _end() {
	orm.EndOrm()
}

func main() {
	_init()
	http.Run()
	defer _end()
}
