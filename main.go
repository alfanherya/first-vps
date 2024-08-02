package main

import (
	"first-app/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	controller.HelloWorld(r)
	r.Run(":8090")
	fmt.Println("Hello World")
}
