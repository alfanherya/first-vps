package main

import (
	"first-app/config"
	"fmt"
)

func main() {
	// initialize environment variable
	config.InitEnv()
	config.InitDb()
	defer config.DB.Close()

	// setup router
	r := config.SetupRouter()

	// run the server
	port := config.GetEnv("PORT", "8090")
	r.Run(fmt.Sprintf(":%s", port))

	fmt.Println("Hello World")
}
