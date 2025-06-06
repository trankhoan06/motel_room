package ginRent

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/rent/biz"
	"main.go/modules/rent/storage"
	"net/http"
)

func DeletedRent(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var RentId common.IdCommon
		if err := c.ShouldBindJSON(&RentId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewRentCommon(store)
		request := c.MustGet(common.Current_user).(common.Requester)
		RentId.UserId = request.GetUserId()
		if err := business.NewDeletedRent(c.Request.Context(), &RentId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "This room will be deleted after 7 days."})
	}
}
