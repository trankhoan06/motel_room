package ginUser

import (
	"github.com/gin-gonic/gin"
	fb "github.com/huandu/facebook/v2"
	"gorm.io/gorm"
	"main.go/modules/user/model"
	"net/http"
)

func LoginFacebook(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.LoginMedia
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := fb.Get("/me", fb.Params{
			"fields":       "first_name,id",
			"access_token": data.YourAccessToken,
		})
		c.JSON(http.StatusOK, gin.H{"data": res["first_name"], "email": res["id"]})

	}
}
