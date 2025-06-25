package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	storageRent "main.go/modules/rent/storage"
	"main.go/modules/user_like_room/biz"
	"main.go/modules/user_like_room/model"
	"main.go/modules/user_like_room/storage"
	"net/http"
)

func ListUserLikeRoom(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		var list model.ListUserLikeRoom
		if err := ctx.ShouldBindJSON(&list); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlMode(db)
		storeRent := storageRent.NewSqlModel(db)
		business := biz.NewUserLikeRoomCommon(store, storeRent)
		data, err := business.NewListUserLikeRoom(ctx.Request.Context(), list.RentId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
