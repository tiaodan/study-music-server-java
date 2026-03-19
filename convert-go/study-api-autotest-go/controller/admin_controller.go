package controller

import (
	"net/http"
	"study-api-autotest-go/models"
	"study-api-autotest-go/service"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *service.AdminService
}

func NewAdminController() *AdminController {
	return &AdminController{
		adminService: service.NewAdminService(),
	}
}

func (c *AdminController) Login(ctx *gin.Context) {
	var req models.AdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := c.adminService.Login(req.Username, req.Password)
	ctx.JSON(http.StatusOK, resp)
}
