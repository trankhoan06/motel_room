package ginUser

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
	"main.go/modules/user/model"
	"net/http"
	"os"
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
func LoginGoogle(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.LoginMedia
		// Lấy ID token từ yêu cầu JSON
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Kiểm tra nếu ID token không có trong yêu cầu
		if data.YourAccessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID token is required"})
			return
		}

		// Kiểm tra tính hợp lệ của ID token với Google
		payload, err := idtoken.Validate(c.Request.Context(), data.YourAccessToken, os.Getenv("CLIENTID")) // "YOUR_CLIENT_ID" là Client ID của ứng dụng
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to validate ID token", "details": err.Error()})
			return
		}

		// Trả về thông tin người dùng từ payload
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id":             payload.Claims["sub"],            // ID người dùng
				"email":          payload.Claims["email"],          // Email người dùng
				"name":           payload.Claims["name"],           // Tên người dùng
				"given_name":     payload.Claims["given_name"],     // Tên đầu tiên
				"family_name":    payload.Claims["family_name"],    // Họ người dùng
				"picture":        payload.Claims["picture"],        // URL ảnh đại diện
				"locale":         payload.Claims["locale"],         // Ngôn ngữ người dùng
				"verified_email": payload.Claims["email_verified"], // Kiểm tra email đã xác minh chưa
			},
		})
	}
}
