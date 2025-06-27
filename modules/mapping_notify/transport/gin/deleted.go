package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/modules/mapping_notify/biz"
	"main.go/modules/mapping_notify/model"
	storageMappingNotify "main.go/modules/mapping_notify/storage"
	"net/http"
)

func DeletedMappingNotify(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateMappingNotify
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storageMappingNotify.NewSqlModel(db)
		business := biz.NewMappingNotifyCommon(store)
		if err := business.NewDeletedMappingNotify(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
