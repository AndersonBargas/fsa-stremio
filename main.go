package main

import (
	"FSA/api"
	"FSA/db"
	"FSA/utils"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AndersonBargas/rainstorm/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	webServerPort := flag.String("port", "8080", "The port which the WebServer will listen to")
	databasePath := flag.String("db", "./fsa.db", "The path where the database will be stored")

	utils.ValidarPortaTCP(*webServerPort)

	// basic configs
	// c, err := configs.NewBasicConfigsEnv()
	// if err != nil {
	// 	fmt.Println("Error obtaining the basic configs from env: %w", err)
	// 	os.Exit(1)
	// }

	// rainstorm - database
	rainstormInstance, err := rainstorm.Open(*databasePath)
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
	e.Logger.Fatal(e.Start(":" + *webServerPort))
}
