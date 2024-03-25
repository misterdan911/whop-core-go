package main

import (
	"os"
	"whop-core-go/app"
)

func main() {

	app.Init()

	// listen and serve

	if os.Getenv("APP_ENV") == "local" {
		// supaya gak keluar notifikasi firewall di windows
		app.GinEngine.Run("127.0.0.1:" + os.Getenv("APP_PORT"))
	}else {
		app.GinEngine.Run(":" + os.Getenv("APP_PORT"))
	}
}
