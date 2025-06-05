package ginReview

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	storageRent "main.go/modules/rent/storage"
	"main.go/modules/room_reviews/biz"
	"main.go/modules/room_reviews/model"
	"main.go/modules/room_reviews/storage"
	"net/http"
)

func DeteledReview(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var data model.DeletedReviews
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		data.UserId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		storeRent := storageRent.NewSqlModel(db)
		business := biz.NewUserReviewsCommon(store, storeRent)
		if err := business.NewDeleteReview(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(200, gin.H{"data": data})
	}
}
