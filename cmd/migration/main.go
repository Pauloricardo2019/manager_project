package main

import (
	"gerenciador/adapter/database/sql"
	"gerenciador/cmd/migration/migrations"
	"github.com/go-gormigrate/gormigrate/v2"
	"log"
)

func main() {
	db, err := sql.GetGormDB()
	if err != nil {
		log.Fatal(err)
	}

	migrationsToExec := migrations.GetMigrationsToExec()
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrationsToExec)

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")

}
