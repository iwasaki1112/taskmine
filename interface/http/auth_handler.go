package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	config oauth2.Config
}

func NewAuthHandler(config oauth2.Config) *AuthHandler {
	return &AuthHandler{
		config: config,
	}
}

func (h *AuthHandler) Auth(c *gin.Context) {
	url := h.config.AuthCodeURL("random")
	c.Redirect(http.StatusTemporaryRedirect, url)
}
