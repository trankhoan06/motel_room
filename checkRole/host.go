package checkRole

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"main.go/common"
	"main.go/modules/user/model"
)

func RoleHost() func(*gin.Context) {
	return func(c *gin.Context) {
		request := c.MustGet(common.Current_user).(common.Requester)
		if *request.GetRole() != model.RoleUserHost {
			log.Println("Exactly_token err:", "user don't host")
			panic(common.ErrPermissionRole(errors.New("user don't host")))
		}
		c.Next()
	}
}
