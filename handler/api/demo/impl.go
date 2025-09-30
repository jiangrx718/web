package demo

import (
	"web/gopkg/gins"
	"web/internal/service"
	"web/internal/service/demo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g           *gin.RouterGroup
	demoService service.Demo
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:           g,
		demoService: demo.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/demo")
	g.POST("/create", h.CreateDemo)
	g.GET("/get", h.GetDemo)
	g.GET("/list", h.PagingDemo)
	g.POST("/update", h.UpdateDemo)
	g.POST("/delete", h.DeleteDemo)
}
