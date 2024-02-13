package http

import (
	"github.com/Akshayvij07/thambola-generator/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

// @title Clinic Application API
// @title Clean- Application API
// @version 1.0.0
// @description  Backend API Golang using Clean Code architecture
// @contact.name API Support
// @contact.email				akshayvijay2000@gmail.com
// @BasePath /
// @query.collection.format multi
func NewServerHTTP(ticketHandler *handler.TicketHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	// engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Auth middleware
	api := engine.Group("/api")

	api.POST("create-ticket", ticketHandler.CreateTicket)
	api.GET("all-ticket", ticketHandler.AllTicket)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8080")
}
