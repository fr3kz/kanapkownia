package main

import (
	"./controllers"
	"./routers"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	controllers.ConnectToDB()
	routers.StartServer()
}
