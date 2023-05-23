package main

import (
	"skyshi/config"
	"skyshi/routes"
)

func main() {
	config.LoadEnvVariable()
	db := config.Config()

	r := routes.Route(db)
	r.Run(":3030")
}
