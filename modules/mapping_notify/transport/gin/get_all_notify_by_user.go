package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/mapping_notify/biz"
	"main.go/modules/mapping_notify/storage"
	"net/http"
)

func GetAllNotifyByUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewMappingNotifyCommon(store)
		data, err := business.NewGetAllNotifyByUser(c.Request.Context(), userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
