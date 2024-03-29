package Filter

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func RequestTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 在请求处理之前做一些事情
		log.Println("开始时间:", start)

		// 处理请求
		c.Next()

		// 请求处理完成后做一些事情
		end := time.Now()
		log.Println("结束时间:", end)
		log.Println("请求处理时间:", end.Sub(start))
	}
}
