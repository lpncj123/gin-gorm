package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	models "gogingorm/modals"
	"log"
	"net/http"
)

func InsertUserNameByRedis(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn := models.REDIS.Get() // 从连接池获取一个连接
	defer conn.Close()         // 使用完毕后记得关闭连接
	//选择数据库
	if _, err := conn.Do("SELECT", 1); err != nil { // 注意：Redis数据库索引是从0开始的
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select database"})
		return
	}
	// 例如，从数据库中查询用户信息
	_, err := conn.Do("SET", "name", user.UserName)
	if err != nil {
		log.Printf("Failed to set user name in Redis: %v", err)
		// 可能需要向客户端返回错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user name"})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"errCode": "000000000",
		"errMsg":  "成功",
	})
}

func GetUserNameByRedis(c *gin.Context) {
	// 从Redis查询key为name的值
	conn := models.REDIS.Get() // 从连接池获取一个连接
	defer conn.Close()         // 使用完毕后记得关闭连接
	//选择数据库
	if _, err := conn.Do("SELECT", 1); err != nil { // 注意：Redis数据库索引是从0开始的
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select database"})
		return
		// 可能需要向客户端返回错误
	}
	valueBytes, err := redis.Bytes(conn.Do("GET", "name"))
	if err != nil {
		log.Printf("Failed to get value from Redis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get value"})
		return
	}

	username := string(valueBytes)
	log.Printf("User name: %s", username)

	c.JSON(http.StatusOK, gin.H{
		"errCode": "000000000",
		"errMsg":  "成功",
		"data":    username,
	})

}
