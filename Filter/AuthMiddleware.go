package Filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求路径是否需要拦截
		if needAuth(c.Request.URL.Path) {
			// 这里可以添加验证逻辑，例如检查请求头中是否包含有效的身份验证信息
			// 如果验证失败，可以中止请求并返回相应的错误信息
			// 如果验证成功，可以继续处理请求
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// 如果请求路径不需要拦截，则继续处理请求
		c.Next()
	}
}

// 判断请求路径是否需要进行身份验证
func needAuth(path string) bool {
	// 这里可以根据实际需求判断是否需要拦截
	// 可以使用正则表达式或其他方式来匹配需要拦截的路径
	// 这里只是一个简单示例，实际应用中需要根据具体情况进行修改
	return path == "/user"
}
