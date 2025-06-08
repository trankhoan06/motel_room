package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	StorageEmail "main.go/modules/email/storage"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"main.go/worker"
	"net/http"
)

func ChangeForgotPassword(db *gorm.DB, taskDistributor worker.TaskDistributor) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.NewPasswordForgot
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		storeEmail := StorageEmail.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewRegisterUserBiz(store, hash, storeEmail, taskDistributor)
		if err := business.NewChangePasswordForgot(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})

	}
}
