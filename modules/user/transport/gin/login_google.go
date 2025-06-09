package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/component/tokenprovider"
	"main.go/config"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"net/http"
)

//	func LoginGoogle(db *gorm.DB) func(ctx *gin.Context) {
//		return func(c *gin.Context) {
//			var data model.LoginMedia
//			// Lấy access token từ yêu cầu JSON
//			if err := c.ShouldBindJSON(&data); err != nil {
//				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//				return
//			}
//			oauth2Config := oauth2.Config{
//				ClientID:     "",
//				ClientSecret: "",
//				RedirectURL:  "https://developers.google.com/oauthplayground",
//				Endpoint:     google.Endpoint,
//				Scopes: []string{
//					"openid", // Scope yêu cầu OpenID
//					"https://www.googleapis.com/auth/userinfo.email",   // Để lấy email người dùng
//					"https://www.googleapis.com/auth/userinfo.profile", // Để lấy thông tin hồ sơ người dùng
//				},
//			}
//			// Kiểm tra nếu access token không có trong yêu cầu
//			token, err := oauth2Config.Exchange(c.Request.Context(), data.YourAccessToken)
//			if err != nil {
//				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//				return
//			}
//
//			// Tạo OAuth client sử dụng access token
//			client := oauth2.NewClient(c.Request.Context(), oauth2.StaticTokenSource(
//				&oauth2.Token{AccessToken: token.AccessToken},
//			))
//
//			// Khởi tạo dịch vụ OAuth2 từ client
//			service, err := oauth3.New(client)
//			if err != nil {
//				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create OAuth client"})
//				return
//			}
//
//			// Gọi Google API để lấy thông tin người dùng (bao gồm cả email)
//			userInfo, err := service.Userinfo.Get().Do()
//			if err != nil {
//				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info"})
//				return
//			}
//
//			// Trả về thông tin người dùng (bao gồm cả email) dưới dạng JSON
//			c.JSON(http.StatusOK, gin.H{
//				"data": gin.H{
//					"id":             userInfo.Id,            // ID người dùng
//					"email":          userInfo.Email,         // Email người dùng
//					"name":           userInfo.Name,          // Tên người dùng
//					"given_name":     userInfo.GivenName,     // Tên đầu tiên
//					"family_name":    userInfo.FamilyName,    // Họ người dùng
//					"picture":        userInfo.Picture,       // URL ảnh đại diện
//					"locale":         userInfo.Locale,        // Ngôn ngữ người dùng
//					"verified_email": userInfo.VerifiedEmail, // Kiểm tra email đã xác minh chưa
//				},
//			})
//		}
//	}
//
// https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email
func LoginGoogle(db *gorm.DB, provider tokenprovider.TokenProvider, cfg *config.Config) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.LoginMedia
		// Lấy ID token từ yêu cầu JSON
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewLoginBiz(store, provider, hash, cfg)
		data1, er := business.NewLoginGoogle(c.Request.Context(), &data, 30*24)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
			return
		}
		// Trả về thông tin người dùng từ payload
		c.JSON(http.StatusOK, gin.H{
			"data": data1,
		})
	}
}
