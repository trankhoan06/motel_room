package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	storageRent "main.go/modules/rent/storage"
	"main.go/modules/search/biz"
	"main.go/modules/search/model"
	"main.go/modules/search/storage"
	"net/http"
)

func SearchRent(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.SearchRent
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		storeRent := storageRent.NewSqlModel(db)
		business := biz.NewSearchRentCommon(store, storeRent)
		res, err := business.NewSearchRent(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
