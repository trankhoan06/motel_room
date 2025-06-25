package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/notify/biz"
	"main.go/modules/notify/model"
	"main.go/modules/notify/storage"
	"main.go/worker"
	"net/http"
)

func CreateNotify(db *gorm.DB, taskDistributor worker.TaskDistributor) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateNotify
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		data.UserId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		business := biz.NewNotifyTaskCommon(store, taskDistributor)
		if err := business.NewCreateNotify(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
