package controller

import (
	"net/http"
	"strconv"
	"study-api-autotest-go/models"
	"study-api-autotest-go/service"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	songService *service.SongService
}

func NewSongController() *SongController {
	return &SongController{
		songService: service.NewSongService(),
	}
}

func (c *SongController) AddSong(ctx *gin.Context) {
	var req models.SongRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := c.songService.AddSong(&req)
	ctx.JSON(http.StatusOK, resp)
}

func (c *SongController) UpdateSong(ctx *gin.Context) {
	var req models.SongRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := c.songService.UpdateSong(&req)
	ctx.JSON(http.StatusOK, resp)
}

func (c *SongController) DeleteSong(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	resp := c.songService.DeleteSong(uint(id))
	ctx.JSON(http.StatusOK, resp)
}

func (c *SongController) SongOfId(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	resp := c.songService.SongOfId(uint(id))
	ctx.JSON(http.StatusOK, resp)
}

func (c *SongController) SongOfSingerId(ctx *gin.Context) {
	singerIdStr := ctx.Query("singerId")
	singerId, err := strconv.ParseUint(singerIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid singerId"})
		return
	}
	resp := c.songService.SongOfSingerId(uint(singerId))
	ctx.JSON(http.StatusOK, resp)
}

func (c *SongController) SongOfName(ctx *gin.Context) {
	name := ctx.Query("name")
	resp := c.songService.SongOfName(name)
	ctx.JSON(http.StatusOK, resp)
}

func (c *SongController) AllSong(ctx *gin.Context) {
	resp := c.songService.AllSong()
	ctx.JSON(http.StatusOK, resp)
}
