package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"log"
	"main.go/component/middleware"
	jwt2 "main.go/component/tokenprovider/jwt"
	"main.go/config"
	"main.go/cronjob"
	emailSend "main.go/email"
	"main.go/modules/email/biz"
	StorageEmail "main.go/modules/email/storage"
	ginMail "main.go/modules/email/transport/gin"
	ginFollow "main.go/modules/follower/transport/gin"
	ginRent "main.go/modules/rent/transport/gin"
	ginReview "main.go/modules/room_reviews/transport/gin"
	"main.go/modules/upload"
	storageUser "main.go/modules/user/storage"
	ginUser "main.go/modules/user/transport/gin"
	ginUserLikeRoom "main.go/modules/user_like_room/transport/gin"
	"main.go/permission"
	ProviderMysql "main.go/provider/mysql"
	"main.go/worker"
	"sync"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := ProviderMysql.NewMysql(cfg)
	if err != nil {
		log.Fatal(err)
	}

	jwtPrefix := jwt2.NewJwtProvider(cfg.App.Prefix, cfg.App.Secret)
	auth := storageUser.NewSqlModel(db)
	middle := middleware.NewModelMiddleware(auth, jwtPrefix)
	fmt.Println(cfg.Redis.Host + ":" + cfg.Redis.Port)
	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
	}
	taskDistributor := worker.NewRedisTaskDistributor(&redisOpt)
	r := gin.Default()
	r.Use(middle.Recover())
	configCORS := setupCors()
	r.Use(cors.New(configCORS))
	r.Static("/static", "./static")
	u := r.Group("/user")
	{
		u.POST("/register", ginUser.Register(db, taskDistributor))
		u.GET("/login", ginUser.Login(db, jwtPrefix, cfg))
		u.GET("/login_facebook", ginUser.LoginFacebook(db))
		u.POST("/login_google", ginUser.LoginGoogle(db, jwtPrefix, cfg))
		u.GET("/get_profile", middle.RequestAuthorize(), ginUser.GetProfile(db))
		u.POST("/forgot_password", ginUser.ForgotPassword(db, taskDistributor))
		u.PATCH("/change_password", middle.RequestAuthorize(), ginUser.ChangePassword(db, taskDistributor))
		u.PATCH("/update_user", middle.RequestAuthorize(), ginUser.UpdateUser(db))
		u.DELETE("/deleted_account", ginUser.DeleteAccount(db))
		u.PATCH("/change_forgot_password", ginUser.ChangeForgotPassword(db, taskDistributor))
	}
	mail := r.Group("/mail")
	mail.PATCH("/verify_code_email", ginMail.VerifyCodeEmail(db, jwtPrefix))
	mail.PATCH("/verify_forgot_password", ginMail.VerifyForgotPassword(db))
	mail.POST("/create_verify_code_email", ginMail.CreateVerifyCodeEmail(db))
	{

	}
	f := r.Group("/follow")
	{
		f.POST("/create", middle.RequestAuthorize(), ginFollow.CreateFollow(db))
		f.POST("/cancel", middle.RequestAuthorize(), ginFollow.CancelFollow(db))
		f.GET("/list_follower", ginFollow.ListFollower(db))
		f.GET("/list_following", ginFollow.ListFollowing(db))
	}
	rent := r.Group("/rent")
	{
		rent.GET("/list_rent", ginRent.ListRent(db))
		rent.GET("/list_rent_the_best_amount_review", ginRent.ListRentTheBestAmountReview(db))
	}
	rent.Use(middle.RequestAuthorize())
	rent.Use(permission.RoleHost())
	{
		rent.POST("/create", ginRent.CreateRent(db))
		rent.DELETE("/deleted", ginRent.DeletedRent(db))
		rent.PATCH("/update", ginRent.DeletedRent(db))
	}
	image := r.Group("/image")
	{
		image.POST("/upload", upload.UploadImage(db))
	}

	rate := r.Group("/rate")
	{
		rate.GET("/list_review", ginReview.ListReview(db))

	}
	rate.Use(middle.RequestAuthorize())
	{
		rate.POST("/create_review", ginReview.CreateReview(db))
		rate.PATCH("/update_review", ginReview.UpdateReview(db))
		rate.DELETE("/deleted_review", ginReview.DeteledReview(db))
	}

	like := r.Group("/like")
	like.Use(middle.RequestAuthorize())
	{
		like.POST("/create", ginUserLikeRoom.CreateUserLikeRoom(db))
		like.DELETE("/deleted", ginUserLikeRoom.DeletedUserLikeRoom(db))
		like.GET("/list", ginUserLikeRoom.ListUserLikeRoom(db))
	}
	//task
	accountSto := storageUser.NewSqlModel(db)
	emailSto := StorageEmail.NewSqlModel(db)
	emailCase := biz.NewSendEmailBiz(emailSto, accountSto)
	NewEmail := emailSend.NewGmailSender(cfg.Email.EmailSenderName, cfg.Email.EmailSenderAddress, cfg.Email.EmailSenderPassword)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		processor := worker.NewRedisTaskProcessor(&redisOpt,
			accountSto,
			NewEmail,
			emailCase,
		)
		err1 := processor.Start()
		if err1 != nil {
			log.Fatal(err1)

		}
	}()
	//cronjob 24h daily
	cronjob.Cronjob(db)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//defer c.Stop()
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
