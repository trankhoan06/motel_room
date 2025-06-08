package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	storageRent "main.go/modules/rent/storage"
	"main.go/modules/user_like_room/biz"
	"main.go/modules/user_like_room/storage"
	"net/http"
)

func ListUserLikeRoom(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		userId := ctx.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlMode(db)
		storeRent := storageRent.NewSqlModel(db)
		business := biz.NewUserLikeRoomCommon(store, storeRent)
		data, err := business.NewListUserLikeRoom(ctx.Request.Context(), userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
