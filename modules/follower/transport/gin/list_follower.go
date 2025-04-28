package ginFollow

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/modules/follower/biz"
	"main.go/modules/follower/model"
	"main.go/modules/follower/storage"
	storageUser "main.go/modules/user/storage"
	"net/http"
)

func ListFollower(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateFollower
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		store := storage.NewSql(db)
		storeUser := storageUser.NewSqlModel(db)
		business := biz.NewFollowUserCommon(store, storeUser)
		follows, err := business.NewListFollower(c.Request.Context(), data.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": follows})
	}
}
