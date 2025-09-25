package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EdgeHandler struct{}
type EdgeData struct {
	Key   string `json:"key"`
	Count int    `json:"count"`
}

func NewEdgeHandler(router *gin.Engine) {
	handler := &EdgeHandler{}

	v1 := router.Group("/api/v1")
	{
		v1.POST("/save", handler.SaveEdgeData) // Example endpoint to save edge data
	}
}

func (h *EdgeHandler) SaveEdgeData(c *gin.Context) {
	var data EdgeData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	// Save the data to the database
	// For now, just log it
	log.Printf("edge data: %+v", data)
	c.JSON(http.StatusOK, gin.H{"message": "edge data saved successfully"})
}
