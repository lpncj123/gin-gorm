package Filter

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		if method == "GET" {
			// GET 请求
			getParams := context.Request.URL.Query()
			log.Println("GET 请求参数有：", getParams)
		} else if method == "POST" {
			// POST 请求
			err := context.Request.ParseForm()
			if err != nil {
				log.Println("解析表单参数出错：", err)
				context.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			postParams := context.Request.PostForm
			log.Println("POST 请求参数有：", postParams)
		}
		// 继续处理请求
		context.Next()
	}
}
