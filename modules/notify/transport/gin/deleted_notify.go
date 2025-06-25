package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	storageMappingNotify "main.go/modules/mapping_notify/storage"
	"main.go/modules/notify/biz"
	"main.go/modules/notify/model"
	"main.go/modules/notify/storage"
	"net/http"
)

func DeletedNotify(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.DeletedNotify
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		storeMapping := storageMappingNotify.NewSqlModel(db)
		data.UserId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		business := biz.NewNotifyCommon(store, storeMapping)
		if err := business.NewDeletedNotify(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
