package main

import (
	"fmt"
	"todo-clist/pkg"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dbPath, err := pkg.GetDefaultDBPath()
	if err != nil {
		panic(fmt.Sprintf("can't access database: %s", err))
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("can't create database: %s", err))
	}

	app := NewApp(db)
	if err := app.Execute(); err != nil {
		panic(err)
	}
}
