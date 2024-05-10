package router

import (
	"github.com/gin-gonic/gin"
	"gogingorm/Filter"
	"gogingorm/controller"
)

func App() *gin.Engine {
	ginServer := gin.Default()
	// 用户信息相关路由
	userRoutes := ginServer.Group("/user")
	{
		userRoutes.GET("", Filter.UserMiddle(), controller.GetUser)
	}

	// 用户创建和更新操作
	userManagementRoutes := ginServer.Group("/user-management")
	{
		userManagementRoutes.POST("/create", Filter.UserMiddle(), controller.InsertUser)
		userManagementRoutes.POST("/transactional", Filter.UserMiddle(), controller.UserTransactional)
	}

	// 用户名称存储和检索
	userNameRoutes := ginServer.Group("/user-name")
	{
		userNameRoutes.POST("/save-by-redis", Filter.UserMiddle(), controller.InsertUserNameByRedis)
		userNameRoutes.GET("/retrieve-by-redis", Filter.UserMiddle(), controller.GetUserNameByRedis)
	}

	return ginServer
}
