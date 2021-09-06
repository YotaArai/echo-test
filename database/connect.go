package database

import (
  "os"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

// DBを使い回すことで、DBへのConnectとCloseを毎回しないようにする
var DB *gorm.DB

func Connect() {

  user := os.Getenv("DB_USERNAME")
  password := os.Getenv("DB_PASSWORD")
  host := os.Getenv("DB_HOST")
  port := os.Getenv("DB_PORT")
  database := os.Getenv("DB_DATABASE")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=Asia/Tokyo"
  DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
