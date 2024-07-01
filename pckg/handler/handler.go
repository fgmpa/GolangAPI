package handler

import (
	"GolangAPI/pckg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		cars := api.Group("/cars")
		{
			cars.POST("/", h.createCar)
			cars.GET("/:id", h.getCar)
			cars.GET("/", h.getAllCars)
			cars.PUT("/:id", h.updateCar)
			cars.DELETE("/:id", h.deleteCar)
		}
	}
	return router
}
