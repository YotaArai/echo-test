package database

import (
  "os"
  "net/url"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

// DBを使い回すことで、DBへのConnectとCloseを毎回しないようにする
var DB *gorm.DB

func Connect() {
  dsn := ""
  if(os.Getenv("APP_ENV") == "local"){
    user := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    database := os.Getenv("DB_DATABASE")

    dsn = "host=" + host + " user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=Asia/Tokyo"
    
  }else{
    mydb_dsn := os.Getenv("DATABASE_URL")
    parsed_dsn, _ := url.Parse(mydb_dsn)
    dsn = parsed_dsn.String()
    dsn += " sslmode=require"
  }
  DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
