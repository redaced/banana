package main

import (
	"github.com/redaced/banana/config"
	// "github.com/redaced/banana/controller"
	// "banana/database"
	"github.com/redaced/banana/route"

	// "fmt"
	"log"
	"net/http"
	// "reflect"
)

func main() {
	// Set Configuration
	config := config.SetConfig()

	// Database Connection
	// db := database.Opendb(config.Db_user, config.Db_password, config.Db_name)

	// Routing
	route.Resource("/test")
	// route.Get("/welcome", controller.Get)
	r := route.Invoker()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	// On service
	log.Fatal(http.ListenAndServe(config.Port, r))
}
