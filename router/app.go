package router

import (
	"github.com/gin-gonic/gin"
	"gogingorm/Filter"
	"gogingorm/controller"
)
import "github.com/thinkerou/favicon"

func App() *gin.Engine {
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./favicon.ico"))
	ginServer.Use(Filter.RequestTimeMiddleware(), Filter.AuthMiddleware())
	ginServer.GET("/user", Filter.UserMiddle(), controller.GetUser)
	ginServer.POST("/insertUser", Filter.UserMiddle(), controller.InsertUser)
	ginServer.POST("/UserTransactional", Filter.UserMiddle(), controller.UserTransactional)
	return ginServer
}
