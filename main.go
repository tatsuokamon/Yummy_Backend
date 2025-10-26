package main

import (
	"os"

	"local.package.backend"
)

func getKey() string {
	return os.Getenv("YUMKEY")
}

func getPort() int {
	return 3000
}

func main() {
	app := backend.NewApp()
	app.SetKey(getKey())
	app.SetRoute()
	app.Run(getPort())
}
