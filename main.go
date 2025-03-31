package main

import (
	"github.com/gin-contrib/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main.go/component/middleware"
	jwt2 "main.go/component/tokenprovider/jwt"
	"main.go/modules/upload"
	storage2 "main.go/modules/user/storage"
	ginUser "main.go/modules/user/transport/gin"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := os.Getenv("DOMAIN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	jwtPrefix := jwt2.NewJwtProvider(os.Getenv("PREFIX"), os.Getenv("SECRET"))
	auth := storage2.NewSqlModel(db)
	middle := middleware.NewModelMiddleware(auth, jwtPrefix)
	r := gin.Default()
	r.Use(middle.Recover())
	configCORS := setupCors()
	r.Use(cors.New(configCORS))
	r.Static("/static", "./static")
	u := r.Group("/user")
	{
		u.GET("/register", ginUser.Register(db))
		u.GET("/login", ginUser.Login(db, jwtPrefix))
		u.GET("/login_facebook", ginUser.LoginFacebook(db))
		u.POST("/login_google", ginUser.LoginGoogle(db))
		u.GET("/get_profile", middle.RequestAuthorize(), ginUser.GetProfile(db))
		u.GET("/forgot_password", ginUser.ForgotPassword(db))
		u.PATCH("/verify_code_email", ginUser.VerifyCodeEmail(db, jwtPrefix))
		u.PATCH("/verify_forgot_password", ginUser.VerifyForgotPassword(db))
		u.PATCH("/change_password", middle.RequestAuthorize(), ginUser.ChangePassword(db))
		u.PATCH("/update_user", middle.RequestAuthorize(), ginUser.UpdateUser(db))
		u.PATCH("/change_forgot_password", ginUser.ChangeForgotPassword(db))
		u.POST("/create_verify_code_email", ginUser.CreateVerifyCodeEmail(db))
	}
	image := r.Group("/image")
	{
		image.POST("/upload", upload.UploadImage(db))
	}
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func setupCors() cors.Config {
	configCORS := cors.DefaultConfig()
	configCORS.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	configCORS.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"}
	configCORS.AllowCredentials = true
	//configCORS.AllowOrigins = []string{"http://localhost:3000"}
	configCORS.AllowAllOrigins = true

	return configCORS
}
