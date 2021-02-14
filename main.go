package main

import (
	"github.com/ashish493/go-to-rest/app"
	"github.com/ashish493/go-to-rest"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}