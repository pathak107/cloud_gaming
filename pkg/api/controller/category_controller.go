package controller

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pathak107/cloudesk/pkg/api/dto"
	"github.com/pathak107/cloudesk/pkg/api/helpers"
	"github.com/pathak107/cloudesk/pkg/api/service"
)

type CategoryController struct {
	catSvc *service.CategoryService
}

func (c *CategoryController) Create(ctx *gin.Context) {
	var category dto.CategoryCreate
	if err := ctx.ShouldBind(&category); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//Image File
	file, _ := ctx.FormFile("image")
	if file != nil {
		filename := fmt.Sprintf("category_image_%s%s", fmt.Sprint(time.Now().UnixMilli()), filepath.Ext(file.Filename))
		log.Println(filename)
		if err := ctx.SaveUploadedFile(file, filename); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		category.ImageUrl = helpers.StringPtr(filename)
	}
	if err := c.catSvc.Create(&category); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "created successfullt",
	})
}

func (c *CategoryController) FindOne(ctx *gin.Context) {
	catID := ctx.Param("cat_id")
	category, err := c.catSvc.FindOne(catID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func (c *CategoryController) FindAll(ctx *gin.Context) {
	categories, err := c.catSvc.FindAll()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func (c *CategoryController) Update(ctx *gin.Context) {
	var category dto.CategoryUpdate
	catID := ctx.Param("cat_id")
	if err := ctx.ShouldBind(&category); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//Image File
	file, _ := ctx.FormFile("image")
	if file != nil {
		filename := fmt.Sprintf("category_image_%s%s", fmt.Sprint(time.Now().UnixMilli()), filepath.Ext(file.Filename))
		log.Println(filename)
		if err := ctx.SaveUploadedFile(file, filename); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		category.ImageUrl = helpers.StringPtr(filename)
	}
	if err := c.catSvc.Update(&category, catID); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "created successfully",
	})
}

func (c *CategoryController) Delete(ctx *gin.Context) {
	catID := ctx.Param("cat_id")
	if err := c.catSvc.Delete(catID); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})
}

func NewCategoryController(catSvc *service.CategoryService) *CategoryController {
	return &CategoryController{
		catSvc: catSvc,
	}
}
