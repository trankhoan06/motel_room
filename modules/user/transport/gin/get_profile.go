package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"net/http"
)

func GetProfile(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.Current_user).(common.Requester)
		c.JSON(http.StatusOK, gin.H{"data": u})
	}
}
