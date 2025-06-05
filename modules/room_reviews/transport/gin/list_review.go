package ginReview

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	storageRent "main.go/modules/rent/storage"
	"main.go/modules/room_reviews/biz"
	"main.go/modules/room_reviews/storage"
	"net/http"
	"strconv"
)

func ListReview(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		rentId, err := strconv.Atoi(c.Query("rent_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		store := storage.NewSqlModel(db)
		storeRent := storageRent.NewSqlModel(db)
		business := biz.NewUserReviewsCommon(store, storeRent)
		data, err := business.NewListReview(c.Request.Context(), rentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
