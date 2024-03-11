package handler

import (
	"exampleAPIs/model"
	"exampleAPIs/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPort interface {
	PostHandlers(c *gin.Context)
	PatchHandlers(c *gin.Context)
	GetHandlers(c *gin.Context)
	DeleteHandlers(c *gin.Context)
}

type handlerAdapter struct {
	s service.ServicePort
}

func NewHanerhandlerAdapter(s service.ServicePort) HandlerPort {
	return &handlerAdapter{s: s}
}

func (h *handlerAdapter) PostHandlers(c *gin.Context) {
	var parametersInput model.ParametersInput
	if err := c.ShouldBindJSON(&parametersInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	err := h.s.PostServices(parametersInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Post method sent parameters auccessfully."})
}

func (h *handlerAdapter) PatchHandlers(c *gin.Context) {
	var parametersUpdate model.ParametersUpdate
	if err := c.ShouldBindJSON(&parametersUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
	}
	err := h.s.PatchServices(parametersUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Update successfully"})
}

func (h *handlerAdapter) GetHandlers(c *gin.Context) {
	parameter1 := c.Query("parameter1")
	if parameter1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "parameter1(get) is required"})
		return
	}
	dataResponse, err := h.s.GetServices(parameter1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "parameter1Info": dataResponse})
}

func (h *handlerAdapter) DeleteHandlers(c *gin.Context) {
	parameter1 := c.Query("parameter1")
	if parameter1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "parameter1(delete) is required"})
		return
	}
	err := h.s.DeleteServices(parameter1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Delete method successfully"})
}
