package main

import (
	"webapp/model"
	"webapp/routes"
)

func main() {
	model.Config()
	routes.Config()
}
