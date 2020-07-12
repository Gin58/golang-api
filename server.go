package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	controller "./controllers/controller"
)

func main() {
	serve()
}

func serve() {
	router := gin.Default()
	router.Static("/views", "./views")
	router.StaticFS("/golang-api", http.Dir("./views/static"))
	router.GET("fetchAllProducts", controller.FetchAllProducts)
	router.GET("/fetchProduct", controller.FindProduct)
	router.POST("/addProduct", controller.AddProduct)
	router.POST("/changeStatusProduct", controller.ChangeStateProduct)
	router.POST("/deleteProduct", controller.DeleteProduct)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}