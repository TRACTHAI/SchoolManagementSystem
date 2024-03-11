package main

import (
	"exampleAPIs/database"
	"exampleAPIs/handler"
	"exampleAPIs/repository"
	"exampleAPIs/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//db := database.Mariadb() // not connected
	db := database.Postgresql()
	defer db.Close()
	r := repository.NewRepositoryAdapter(db)
	s := service.NewServiceAdapter(r)
	h := handler.NewHanerhandlerAdapter(s)

	router := gin.Default()
	router.POST("/post", h.PostHandlers)
	router.PATCH("/patch", h.PatchHandlers)
	router.GET("/get", h.GetHandlers)
	router.DELETE("/delete", h.DeleteHandlers)

	err := router.Run(":8888")
	if err != nil {
		panic(err.Error())
	}
}
