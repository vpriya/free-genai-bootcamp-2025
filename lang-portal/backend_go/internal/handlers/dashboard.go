package handlers

import (
	"lang-portal/backend_go/internal/services"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service *services.DashboardService
}

func NewDashboardHandler(service *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

func (h *DashboardHandler) GetLastStudySession(c *gin.Context) {
	result, err := h.service.GetLastStudySession()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func (h *DashboardHandler) GetStudyProgress(c *gin.Context) {
	result, err := h.service.GetStudyProgress()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func (h *DashboardHandler) GetQuickStats(c *gin.Context) {
	result, err := h.service.GetQuickStats()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}
