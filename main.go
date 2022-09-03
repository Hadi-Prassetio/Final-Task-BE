package main

import (
	"_waysbeans_/database"
	"_waysbeans_/pkg/mysql"
	"_waysbeans_/routes"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	// env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	r := mux.NewRouter()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code

	var port = "5000"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe("localhost:5000", r)
}
