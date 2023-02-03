package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soohyun-lee/side_project_overseas/initializers"
	"github.com/soohyun-lee/side_project_overseas/models"
)

type BoardData struct {
	Text string `json:"text"`
}

func BoardCreate(c *gin.Context) {
	var data BoardData

	err := json.NewDecoder(c.Request.Body).Decode(&data)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
	} else {
		info := models.Board{
			Text: data.Text,
		}
		result := initializers.DB.Create(&info)
		fmt.Println("resutl", result)

		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"Response": "OK",
		})
	}
}

func BoardList(c *gin.Context) {
	boardDate := c.Query("boarddate")
	var allboard []models.Board

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, boardDate)
	endDate := startDate.Add(time.Hour*23 + time.Minute*59)

	initializers.DB.Where(
		"created_at BETWEEN ? AND ?", startDate, endDate).Find(&allboard)

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data": allboard,
	})
}
