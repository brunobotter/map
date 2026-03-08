package main

import (
	"github.com/brunobotter/map/main/app"
	"github.com/brunobotter/map/main/providers"
)

func main() {
	app.NewApplication(providers.List()).Bootstrap()
}
