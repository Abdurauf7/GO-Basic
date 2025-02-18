package routes

import (
	"practice/Rest/controllers/events"
	"practice/Rest/controllers/users"

	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {
	server.GET("/events", events.GetEvents)
	server.GET("/events/:id", events.GetEvent)
	server.POST("/events", events.CreateEvent)
	server.PUT("/events/:id", events.UpdateEvent)
	server.DELETE("/events/:id", events.DeleteEvent)
}

func UsersRoutes(server *gin.Engine) {
	server.POST("/signup", users.Signup)
}
