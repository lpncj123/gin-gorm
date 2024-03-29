package main

import (
	models "gogingorm/modals"
	"gogingorm/router"
)

func main() {
	models.NewGormDB()
	r := router.App()
	r.Run(":8098")
}
