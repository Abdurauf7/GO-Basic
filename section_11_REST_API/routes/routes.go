package routes

import (
	"practice/Rest/controllers/events"
	"practice/Rest/controllers/registration"
	"practice/Rest/controllers/users"
	"practice/Rest/middlewares"

	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {
	autGroup := server.Group("/")
	autGroup.Use(middlewares.Auth)
	autGroup.GET("/events", events.GetEvents)
	autGroup.GET("/events/:id", events.GetEvent)
	autGroup.POST("/events", events.CreateEvent)
	autGroup.PUT("/events/:id", events.UpdateEvent)
	autGroup.DELETE("/events/:id", events.DeleteEvent)
	autGroup.POST("/events/:id/register", registration.RegisterForEvent)
	autGroup.DELETE("/events/:id/register", registration.CancelRegistration)
}

func UsersRoutes(server *gin.Engine) {
	server.POST("/signup", users.Signup)
	server.POST("/login", users.Login)
}
