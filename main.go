package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"tva_api/api"
	"tva_api/config"
	"tva_api/db/mgo"
)

func initServer() http.Handler {
	var server = http.NewServeMux()
	server.Handle("/api/", http.StripPrefix("/api", api.NewApiServer()))
	return server
}
func init() {
	var _, err = mgo.NewDB(config.DBHost, config.DBName)
	if err != nil {
		fmt.Println("Cannot connect to db at ", config.DBHost)
	}
	os.Setenv(config.KeyDB, config.DBName)
}

func main() {

	s := &http.Server{
		Addr:           ":8080",
		Handler:        initServer(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server ready........")
	log.Fatal(s.ListenAndServe())

}
