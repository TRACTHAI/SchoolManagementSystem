package handler

import (
	"exampleAPIs/model"
	"exampleAPIs/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPort interface {
	PostHandlers(c *gin.Context)
	PatchHandlers(c *gin.Context)
	GetHandlers(c *gin.Context)
	DeleteHandlers(c *gin.Context)
	GetAllHandlers(c *gin.Context)
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
	fmt.Println(parametersInput)
	err := h.s.PostServices(parametersInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Post student auccessfully."})
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
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Update student successfully"})
}

func (h *handlerAdapter) GetHandlers(c *gin.Context) {
	parameter1 := c.Query("studentID")
	if parameter1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "studentID(get) is required"})
		return
	}
	dataResponse, err := h.s.GetServices(parameter1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "studentInfo": dataResponse})
}

func (h *handlerAdapter) DeleteHandlers(c *gin.Context) {
	parameter1 := c.Query("studentID")
	if parameter1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "studentID(delete) is required"})
		return
	}
	err := h.s.DeleteServices(parameter1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Delete student successfully"})
}

func (h *handlerAdapter) GetAllHandlers(c *gin.Context) {
	dataResponse, err := h.s.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "studentInfo": dataResponse})
}
