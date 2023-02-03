package main

import (
	"github.com/soohyun-lee/side_project_overseas/initializers"
	"github.com/soohyun-lee/side_project_overseas/models"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Board{})
}

// package routes

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/soohyun-lee/side_project_overseas/initializers"
// 	"github.com/soohyun-lee/side_project_overseas/models"
// )

// type BoardData struct {
// 	Content string `json:"content"`
// }

// func BoardSave(c *gin.Context) {
// 	var data BoardData

// 	err := json.NewDecoder(c.Reqeust.Body).Decode(&data)

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"message": "Error",
// 		})
// 	}
// }

// func UserLogin(c *gin.Context) {
// 	var body struct {
// 		Email string
// 		Name  string
// 	}

// 	c.Bind(&body)

// 	info := models.User{
// 		Email: body.Email,
// 		Name:  body.Name,
// 	}

// 	result := initializers.DB.Create(&info)

// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(400, gin.H{
// 			"error": result.Error,
// 		})
// 		return
// 	}

// 	c.AbortWithStatusJSON(200, gin.H{
// 		"user": info,
// 	})
// }
