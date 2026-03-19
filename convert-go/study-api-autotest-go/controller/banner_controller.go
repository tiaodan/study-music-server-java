package controller

import (
	"net/http"
	"study-api-autotest-go/service"

	"github.com/gin-gonic/gin"
)

type BannerController struct {
	bannerService *service.BannerService
}

func NewBannerController() *BannerController {
	return &BannerController{
		bannerService: service.NewBannerService(),
	}
}

func (c *BannerController) AllBanner(ctx *gin.Context) {
	resp := c.bannerService.AllBanner()
	ctx.JSON(http.StatusOK, resp)
}
