package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/search/biz"
	"main.go/modules/search/model"
	"main.go/modules/search/storage"
	"net/http"
)

func DeletedSearch(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.UpdateSearch

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.UserId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewSearchCommon(store)
		if err := business.NewDeletedSearch(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
