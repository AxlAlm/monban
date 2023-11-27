package main

import (
	"context"
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
	r := webapp.SetupWebApp(&db)

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
		panic("oh no")
	}

	rows, err := tx.Exec(context.Background(), "select * from api_keys")
	fmt.Println(rows)

	http.ListenAndServe(":3000", r)
}
