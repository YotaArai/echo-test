package main

import (
	"echo-test/database"

  "net/http"

  "github.com/labstack/echo/v4"
)

type Post struct {
  Id        int    `json:"id"`
  Content   string `json:"content"`
}


func getPosts(c echo.Context) error {
  posts := []Post{}
  database.DB.Find(&posts)
  return c.JSON(http.StatusOK, posts)
}

func main() {
  e := echo.New()
  database.Connect()
  sqlDB, _ := database.DB.DB()
  defer sqlDB.Close()

  e.GET("/posts", getPosts)

  e.Logger.Fatal(e.Start(":3000"))
}
