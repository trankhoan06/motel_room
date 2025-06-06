package cronjob

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"log"
	"main.go/modules/rent/storage"
)

func Cronjob(db *gorm.DB) {
	fmt.Print("start cronjob")
	c := cron.New(cron.WithSeconds())

	// Job chạy mỗi 0h hằng ngày
	c.AddFunc("0 0 0 * * *", func() {
		store := storage.NewSqlModel(db)
		err := store.DeletedAfter7Day(context.Background())
		if err != nil {
			log.Println(err)
		} else {
			log.Println("success")
		}
	})
	go c.Start()
}
