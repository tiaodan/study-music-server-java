package mapper

import (
	"study-api-autotest-go/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	cfg := config.AppConfig

	DB, err = gorm.Open(mysql.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			},
		),
	})
	if err != nil {
		return err
	}

	// Note: Database tables already exist, skipping AutoMigrate
	// If you need to create new tables, uncomment the following:
	// err = DB.AutoMigrate(
	//     &models.Consumer{},
	//     &models.Singer{},
	//     &models.Song{},
	//     &models.SongList{},
	//     &models.Collect{},
	//     &models.Comment{},
	//     &models.RankList{},
	//     &models.Banner{},
	//     &models.Admin{},
	//     &models.ListSong{},
	//     &models.UserSupport{},
	// )

	return nil
}
