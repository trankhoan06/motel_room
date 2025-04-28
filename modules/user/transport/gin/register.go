package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/config"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	storage2 "main.go/modules/user/storage"
	"net/http"
)

func Register(db *gorm.DB, cfg *config.Config) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.Register
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage2.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewRegisterUserBiz(store, hash, cfg)
		verify, err := business.NewRegister(ctx.Request.Context(), &data, 60)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"verify": verify, "data": true})
	}
}
