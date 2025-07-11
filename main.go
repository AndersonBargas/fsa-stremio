package main

import (
	"FSA/api"
	"FSA/db"
	"fmt"
	"log"
	"os"

	"github.com/AndersonBargas/rainstorm/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// rainstorm - database
	rainstormInstance, err := rainstorm.Open("./fsa.db")
	if err != nil {
		fmt.Println("Error opening the database: %w", err)
		os.Exit(2)
	}
	dbInstance, err := db.NewDatabaseRainstorm(rainstormInstance)
	defer dbInstance.Close()

	// data providers
	// providers := infrastructure.InitProviders()

	// job scheduler
	// scheduler := infrastructure.NewScheduler(providers)

	// echo - webserver
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes, err := api.NewRoutes(e)
	if err != nil {
		log.Fatalf("Error getting the router instance: %v", err)
	}
	routes.RegisterEndpoints()
	e.Logger.Fatal(e.StartTLS(":443", "server.crt", "server.key"))
}
