package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pathak107/cloudesk/pkg/handler"
	"github.com/pathak107/cloudesk/pkg/middleware"
)

func main() {
	h, err := handler.NewCloudHandler()
	if err != nil {
		log.Fatalf("Failed to create handler")
	}
	// Routes setup
	r := gin.Default()
	r.Static("/static", "./public")
	r.Use(middleware.ErrorHandler())
	v1 := r.Group("/api/v1")
	{
		vm := v1.Group("/vm")
		{
			vm.GET("/", func(ctx *gin.Context) {})      //List of all the Vms
			vm.GET("/:vmID", func(ctx *gin.Context) {}) //Information about a single VM
			vm.POST("/", h.LaunchVM)                    //Launch a new VM
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
			auth.POST("/login", func(ctx *gin.Context) {})
			auth.POST("/register", func(ctx *gin.Context) {})
		}
	}

	r.Run(":3000")
}
