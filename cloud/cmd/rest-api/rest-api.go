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
		content := v1.Group("/content")
		{
			content.GET("/apps", func(ctx *gin.Context) {})
			content.GET("/apps/:app_id", func(ctx *gin.Context) {})
			content.POST("/apps", func(ctx *gin.Context) {})
			content.PATCH("/apps/:app_id", func(ctx *gin.Context) {})
			content.DELETE("/apps/:app_id", func(ctx *gin.Context) {})

			content.GET("/apps_pack", func(ctx *gin.Context) {})
			content.GET("/apps_pack/:apps_pack_id", func(ctx *gin.Context) {})
			content.POST("/apps_pack", func(ctx *gin.Context) {})
			content.PATCH("/apps_pack/:apps_pack_id", func(ctx *gin.Context) {})
			content.DELETE("/apps_pack/:apps_pack_id", func(ctx *gin.Context) {})

			content.GET("/categories", middleware.AuthorizeUser(), categoryController.FindAll)
			content.GET("/categories/:cat_id", categoryController.FindOne)
			content.POST("/categories", categoryController.Create)
			content.PATCH("/categories/:cat_id", categoryController.Update)
			content.DELETE("/categories/:cat_id", categoryController.Delete)

			content.GET("/rdp", func(ctx *gin.Context) {})
			content.GET("/rdp/:rdp_id", func(ctx *gin.Context) {})
			content.POST("/rdp", func(ctx *gin.Context) {})
			content.PATCH("/rdp/:rdp_id", func(ctx *gin.Context) {})
			content.DELETE("/rdp/:rdp_id", func(ctx *gin.Context) {})

			content.GET("/membership", func(ctx *gin.Context) {})
			content.GET("/membership/:mem_id", func(ctx *gin.Context) {})
			content.POST("/membership", func(ctx *gin.Context) {})
			content.PATCH("/membership/:mem_id", func(ctx *gin.Context) {})
			content.DELETE("/membership/:mem_id", func(ctx *gin.Context) {})
		}

		vm := v1.Group("/vm")
		{
			vm.GET("/", func(ctx *gin.Context) {})      //List of all the Vms under one acc
			vm.GET("/:vmID", func(ctx *gin.Context) {}) //Information about a single VM
			vm.POST("/", func(ctx *gin.Context) {})     //Launch a new VM
			vm.PUT("/", func(ctx *gin.Context) {})      //action= stop, createImage, take snapshot
			vm.DELETE("/", func(ctx *gin.Context) {})   //Delete a VM

			vm.GET("/connect", func(ctx *gin.Context) {}) //Returns a token
			vm.GET("/status", func(ctx *gin.Context) {})  // Polls or SSE for status
		}

		cloud := v1.Group("/cloud") //To get information related to price. hardware specs, instance types etc
		{
			cloud.GET("")
			cloud.GET("/ping", func(ctx *gin.Context) {}) //Ping Test to find nearest data center
		}

		admin := v1.Group("/admin")
		{
			admin.GET("/vm", func(ctx *gin.Context) {}) //List all the VMS
		}

		payment := v1.Group("/payment")
		{
			payment.GET("/", func(ctx *gin.Context) {}) //Get all transactions, related to one account
		}

		user := v1.Group("/user")
		{
			user.GET("/", func(ctx *gin.Context) {})
			user.GET("/:user_id", func(ctx *gin.Context) {})
			user.POST("/", func(ctx *gin.Context) {})
			user.PATCH("/:user_id", func(ctx *gin.Context) {})
			user.DELETE("/:user_id", func(ctx *gin.Context) {})
			user.GET("/transaction", func(ctx *gin.Context) {})
			user.GET("/subscription/", func(ctx *gin.Context) {})
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/login", userController.Login)
			auth.POST("/register", userController.Register)
		}
	}

	r.Run(":3000")
}
