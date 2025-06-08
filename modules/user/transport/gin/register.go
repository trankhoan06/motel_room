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

func Register(db *gorm.DB, taskDistributor worker.TaskDistributor) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.Register
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		storeEmail := StorageEmail.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewRegisterUserBiz(store, hash, storeEmail, taskDistributor)
		err := business.NewRegister(ctx.Request.Context(), &data, 60)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
