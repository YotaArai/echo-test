package main

import (
	"echo-test-ya/database"

  "net/http"
  "os"

  "github.com/labstack/echo/v4"
  "github.com/joho/godotenv"
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

func getPost(c echo.Context) error {
  id := c.Param("id")

  post := Post{}
  if err := c.Bind(&post); err != nil {
    return err
  }
  database.DB.Find(&post, id)
  return c.JSON(http.StatusOK, post)
}

func updatePost(c echo.Context) error {
  id := c.Param("id")

  post := Post{}
  if err := c.Bind(&post); err != nil {
    return err
  }
  database.DB.Find(&post, id)
  post.Content = c.FormValue("content")
  database.DB.Save(&post)
  return c.JSON(http.StatusOK, post)
}

func createPost(c echo.Context) error {
  post := Post{
		Content: c.FormValue("content"),
	}
  database.DB.Create(&post)
  return c.JSON(http.StatusOK, post)
}

func deletePost(c echo.Context) error {
  id := c.Param("id")
  database.DB.Delete(&Post{}, id)
  return c.NoContent(http.StatusNoContent)
}

func main() {
  err := godotenv.Load()
  if err != nil {
    panic(err.Error())
  }

  e := echo.New()
  database.Connect()
  sqlDB, _ := database.DB.DB()
  defer sqlDB.Close()

  e.GET("/posts", getPosts)
  e.GET("/posts/:id", getPost)
  e.PUT("/posts/:id", updatePost)
  e.POST("/posts", createPost)
  e.DELETE("/posts/:id", deletePost)

  e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
