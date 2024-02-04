package mappings

import (
	"golang/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		v1.GET("/users/:id", controllers.GetUserDetail)
		v1.GET("/users/", controllers.GetUser)
		v1.POST("/login/", controllers.Login)
		v1.PUT("/users/:id", controllers.UpdateUser)
		v1.POST("/users", controllers.PostUser)
	}
}

func GetUser(c *gin.Context) {
	var user []models.User
	_, err := dbmap.Select(&user, "select * from custumers")
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
