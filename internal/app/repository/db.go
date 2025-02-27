package repository

import (
	"fmt"
	"log"
	"os"
	"sakata-server-rest/internal/model" // Game 構造体をインポート

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// .env の読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 環境変数から取得
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	// DSN を組み立てる
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	// DB に接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// 自動マイグレーション
	if err := db.AutoMigrate(&model.Game{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	return db
}
