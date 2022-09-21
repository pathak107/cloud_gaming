package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pathak107/cloudesk/pkg/cloud"
	"github.com/pathak107/cloudesk/pkg/dto"
	"github.com/pathak107/cloudesk/pkg/graphql"
)

func (h *Handler) LaunchVM(ctx *gin.Context) {
	var createVmDTO dto.CreateVmDTO
	if err := ctx.ShouldBind(&createVmDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	vm, err := h.cloudSvc.LaunchVM(context.Background(), &cloud.CreateInstanceParams{
		Name:     createVmDTO.Name,
		Image:    createVmDTO.Image,
		Hardware: createVmDTO.Hardware,
		Storage:  createVmDTO.Storage,
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	//Store information in database about this VM
	graphql.AddVmInfo(&vm)

	ctx.JSON(http.StatusOK, gin.H{
		"data": "vm creation started",
	})
}
