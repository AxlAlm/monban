package main

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./migrate <up|down>")
	}
	migrationDirection := os.Args[1]

	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)

	switch migrationDirection {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		log.Fatalf("Invalid migration direction: %s", migrationDirection)
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Printf("Migrations %s completed successfully", migrationDirection)
}
