package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) Routes() {
	r.setupCORS(r.Engine)

	r.Engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.Engine.POST("/createurl", r.CreateURL)
	r.Engine.GET("/:shortURL", r.RedirectURL)
}
