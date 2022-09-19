package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pathak107/cloudesk/pkg/api/controller"
	"github.com/pathak107/cloudesk/pkg/api/database"
	"github.com/pathak107/cloudesk/pkg/api/middleware"
	"github.com/pathak107/cloudesk/pkg/api/service"
)

func main() {
	//Database setup
	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("Database setup failed: %v", err)
	}

	//Service Setup
	userSvc := service.NewUserService(db.Conn)
	authSvc := service.NewAuthService(db.Conn)
	catSvc := service.NewCategoryService(db.Conn)

	//Controllers setup
	userController := controller.NewUserController(userSvc, authSvc)
	categoryController := controller.NewCategoryController(catSvc)

	// Routes setup
	r := gin.Default()
	r.Static("/static", "./public")
	r.Use(middleware.ErrorHandler())
	v1 := r.Group("/api/v1")
	{
		vm := v1.Group("/vm")
		{
			vm.GET("/", func(ctx *gin.Context) {})      //List of all the Vms under one acc
			vm.GET("/:vmID", func(ctx *gin.Context) {}) //Information about a single VM
			vm.POST("/", func(ctx *gin.Context) {})     //Launch a new VM
			vm.PUT("/", func(ctx *gin.Context) {})      //action= stop, createImage, take snapshot
			vm.DELETE("/", func(ctx *gin.Context) {})   //Delete a VM

			vm.GET("/connect", func(ctx *gin.Context) {}) //Returns a token
			vm.GET("/status", func(ctx *gin.Context) {})  // Polls or SSE for status
			vm.GET("/ping", func(ctx *gin.Context) {})    //Ping Test to find nearest data center
		}

		admin := v1.Group("/admin")
		{
			admin.GET("/vm", func(ctx *gin.Context) {}) //List all the VMS
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/login", userController.Login)
			auth.POST("/register", userController.Register)
		}
	}

	r.Run(":3000")
}
