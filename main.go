package main

import (
	"first-app/config"
	"fmt"
)

func main() {
	// initialize environment variable
	// config.InitEnv()
	// config.InitDb()
	// defer config.DB.Close()

	// setup router
	// r := config.SetupRouter()

	// run the server
	// port := config.GetEnv("PORT", "8090")
	// r.Run(fmt.Sprintf(":%s", port))

	// fmt.Println("Hello World")

	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:     db,
		App:    app,
		Log:    log,
		Config: viperConfig,
	})

	port := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
