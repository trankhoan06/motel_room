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

func ChangeStatusAccount(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var status model.StatusUserValue
		err := c.ShouldBindJSON(&status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		request := c.MustGet(common.Current_user).(common.Requester)
		store := storage.NewSqlModel(db)
		business := biz.NewUserCommonBiz(store)
		if err := business.NewChangeStatusAccount(c.Request.Context(), map[string]interface{}{"id": request.GetUserId()}, status.Status); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Account deleted"})
	}
}
