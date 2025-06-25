package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	storageRent "main.go/modules/rent/storage"
	"main.go/modules/user_like_room/biz"
	"main.go/modules/user_like_room/model"
	"main.go/modules/user_like_room/storage"
	"net/http"
)

func CreateUserLikeRoom(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.CreateUserLikeRoom
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.UserId = ctx.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlMode(db)
		storeRent := storageRent.NewSqlModel(db)
		business := biz.NewUserLikeRoomCommon(store, storeRent)
		if err := business.NewCreateUserLikeRoom(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
