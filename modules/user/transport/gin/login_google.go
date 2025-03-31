package ginUser

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	oauth3 "google.golang.org/api/oauth2/v2"
	"gorm.io/gorm"
	"main.go/modules/user/model"
	"net/http"
)

func LoginGoogle(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.LoginMedia
		// Lấy access token từ yêu cầu JSON
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Kiểm tra nếu access token không có trong yêu cầu
		if data.YourAccessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Access token is required"})
			return
		}

		// Tạo OAuth client sử dụng access token
		client := oauth2.NewClient(c.Request.Context(), oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: data.YourAccessToken},
		))

		// Khởi tạo dịch vụ OAuth2 từ client
		service, err := oauth3.New(client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create OAuth client"})
			return
		}

		// Gọi Google API để lấy thông tin người dùng (bao gồm cả email)
		userInfo, err := service.Userinfo.Get().Do()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info"})
			return
		}

		// Trả về thông tin người dùng (bao gồm cả email) dưới dạng JSON
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id":             userInfo.Id,            // ID người dùng
				"email":          userInfo.Email,         // Email người dùng
				"name":           userInfo.Name,          // Tên người dùng
				"given_name":     userInfo.GivenName,     // Tên đầu tiên
				"family_name":    userInfo.FamilyName,    // Họ người dùng
				"picture":        userInfo.Picture,       // URL ảnh đại diện
				"locale":         userInfo.Locale,        // Ngôn ngữ người dùng
				"verified_email": userInfo.VerifiedEmail, // Kiểm tra email đã xác minh chưa
			},
		})
	}
}
