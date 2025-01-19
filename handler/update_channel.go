package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teacat/chaturbate-dvr/chaturbate"
	"github.com/urfave/cli/v2"
)

//=======================================================
// Request & Response
//=======================================================

type UpdateChannelRequest struct {
	Username   string `json:"username"`
	IsFavorite bool   `json:"is_favorite"`
	Framerate  int    `json:"framerate"`
	Resolution int    `json:"resolution"`
	// ResolutionFallback string `json:"resolution_fallback"`
	// SplitDuration      int    `json:"split_duration"`
	// SplitFilesize      int    `json:"split_filesize"`
	// Interval           int    `json:"interval"`
}

type UpdateChannelResponse struct {
}

//=======================================================
// Factory
//=======================================================

type UpdateChannelHandler struct {
	chaturbate *chaturbate.Manager
	cli        *cli.Context
}

func NewUpdateChannelHandler(c *chaturbate.Manager, cli *cli.Context) *UpdateChannelHandler {
	return &UpdateChannelHandler{c, cli}
}

//=======================================================
// Handle
//=======================================================

func (h *UpdateChannelHandler) Handle(c *gin.Context) {
	var req *UpdateChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	usernames := strings.Split(req.Username, ",")
	for _, username := range usernames {
		if err := h.chaturbate.UpdateChannel(strings.TrimSpace(username), req.Resolution, req.Framerate, req.IsFavorite); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	if err := h.chaturbate.SaveChannels(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &UpdateChannelResponse{})
}
