package main

import (
	"github.com/spf13/viper"
	"log"
	"net/http"

	api "ps.kz/go_search/controller"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	log.Println("Server started")

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":"+viper.GetString("server.port"), router))

}
