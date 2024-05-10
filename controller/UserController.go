package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	models "gogingorm/modals"
	"gorm.io/gorm"
	"log"
	"net/http"
	strings "strings"
)

func GetUser(c *gin.Context) {
	ids := c.Query("ids")
	idSlice := strings.Split(ids, ",")
	// 获取用户信息的逻辑
	// 例如，从数据库中查询用户信息
	users := make([]models.User, 0)
	models.DB.Debug().Table("user").Find(&users, idSlice)

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"errCode": "000000000",
		"errMsg":  "成功",
	})
}

func InsertUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 例如，保存用户信息
	models.DB.Debug().Table("user").Create(&user)

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"errCode": "000000000",
		"errMsg":  "成功",
	})
}

func UserTransactional(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if user.UserAge < 18 {
			log.Println("UserAge:", user.UserAge)
			return errors.New("未成年不能调用")
		}
		if err := tx.Debug().Table("user").Create(&user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"errCode": "000000000",
		"errMsg":  "成功",
	})
}
