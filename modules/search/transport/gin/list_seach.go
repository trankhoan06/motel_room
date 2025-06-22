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

func ListSearch(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ListSearch
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.UserId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewSearchCommon(store)
		res, err := business.NewListSearch(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": res})
	}
}
func ListSearchCustom(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ListSearch
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.UserId = 0
		store := storage.NewSqlModel(db)
		business := biz.NewSearchCommon(store)
		res, err := business.NewListSearch(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": res})
	}
}
