package database

import (
	"log"
	//"os"
	//"time"

	"github.com/lne-io/feedback/models"
	"github.com/lne-io/feedback/config"

	"gorm.io/gorm"
	//goLogger "gorm.io/gorm/logger"
	"gorm.io/driver/sqlite"
)

func ConnectDB() {
	var err error

	/*newLogger := goLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		goLogger.Config{
			SlowThreshold: time.Second,
			LogLevel: goLogger.Info,
			Colorful: true,
		},
	)*/
	databaseLocation := config.GetEnv("DATABASE_LOCATION", "data/")
	databaseName := config.GetEnv("DATABASE_NAME", "feedback.db")
	
	DB, err = gorm.Open(sqlite.Open(databaseLocation + databaseName), &gorm.Config{})
	//Logger: newLogger,
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection opened to database")
	DB.AutoMigrate(&models.Website{}, &models.Feedback{})
	log.Println("Database Migrated")
}