package main

import (
	"exampleAPIs/database"
	"exampleAPIs/handler"
	"exampleAPIs/repository"
	"exampleAPIs/service"
	"exampleAPIs/utility"

	"github.com/gin-gonic/gin"
)

func main() {
	//db := database.Mariadb() // not connected
	db := database.Postgresql()
	defer db.Close()
	utility.CountTables(db)
	r := repository.NewRepositoryAdapter(db)
	// fmt.Println(db)
	s := service.NewServiceAdapter(r)
	h := handler.NewHanerhandlerAdapter(s)

	router := gin.Default()
	router.POST("/addStudent", h.PostHandlers)
	router.PATCH("/patchStudent", h.PatchHandlers)
	router.GET("/getStudent", h.GetHandlers)
	router.DELETE("/deleteStudent", h.DeleteHandlers)
	router.GET("/getStudentAll", h.GetAllHandlers)

	err := router.Run(":8888")
	if err != nil {
		panic(err.Error())
	}
}
