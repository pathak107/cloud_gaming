package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FindOne(ctx *gin.Context) {
	catID := ctx.Param("cat_id")
	category, err := h.catSvc.FindOne(catID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}
