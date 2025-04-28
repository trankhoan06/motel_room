package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/config"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"net/http"
)

func ChangePassword(db *gorm.DB, cfg *config.Config) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ChangePassword
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		hash := common.NewSha256Hash()
		request := c.MustGet(common.Current_user).(common.Requester)
		data.Id = request.GetUserId()
		data.Email = request.GetEmail()
		business := biz.NewRegisterUserBiz(store, hash, cfg)
		if err := business.NewChangePassword(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})

	}
}
