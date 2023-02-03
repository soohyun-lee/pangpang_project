package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soohyun-lee/side_project_overseas/initializers"
	"github.com/soohyun-lee/side_project_overseas/models"
)

func UserLogin(c *gin.Context) {

	var body struct {
		Name string
	}

	c.Bind(&body)

	info := models.User{
		Name: body.Name,
	}

	result := initializers.DB.Find(&info, "name = ?", body.Name)

	if info.ID == 0 {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "No Users",
		})
		return
	}

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": result.Error,
		})
		return
	}

	c.AbortWithStatusJSON(200, gin.H{
		"result": "OK",
	})
}
