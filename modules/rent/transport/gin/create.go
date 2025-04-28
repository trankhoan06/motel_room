package ginRent

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/rent/biz"
	"main.go/modules/rent/model"
	"main.go/modules/rent/storage"
	"net/http"
)

func CreateRent(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateRent
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		request := c.MustGet(common.Current_user).(common.Requester)
		data.UserId = request.GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewRentCommon(store)
		if err := business.NewCreateRent(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
