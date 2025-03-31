package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"net/http"
)

func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UpdateUser

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		requester := c.MustGet(common.Current_user).(common.Requester)
		data.Email = requester.GetEmail()
		business := biz.NewUserCommonBiz(store)
		if err := business.NewUpdateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
