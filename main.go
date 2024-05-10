package main

import (
	models "gogingorm/modals"
	"gogingorm/router"
)

func main() {
	models.NewGormDB()
	models.NewRedis()
	r := router.App()
	r.Run(":8098")
}
