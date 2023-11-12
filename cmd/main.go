package main

import (
	"fmt"
	"monban/config"
	"monban/database"
	"monban/webapp"
	"net/http"
)

func main() {
	fmt.Println(config.Config)
	db := database.NewDB(
		config.Config.DBUsername,
		config.Config.DBPassword,
		config.Config.DBHost,
		config.Config.DBName,
		config.Config.DBPort,
		config.Config.DBMaxPoolConns,
	)
	defer db.Close()
	r := webapp.SetupWebApp()
	http.ListenAndServe(":3000", r)
}
