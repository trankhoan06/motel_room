package ginFollow

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/follower/biz"
	"main.go/modules/follower/model"
	"main.go/modules/follower/storage"
	storageUser "main.go/modules/user/storage"
	"net/http"
)

func CreateFollow(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateFollower
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		request := c.MustGet(common.Current_user).(common.Requester)
		data.FollowerId = request.GetUserId()
		store := storage.NewSql(db)
		storeUser := storageUser.NewSqlModel(db)
		business := biz.NewFollowUserCommon(store, storeUser)
		if err := business.NewCreateFollow(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
