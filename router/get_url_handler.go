package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *Router) RedirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	url, err := r.URLShortnerAPI.GetByShortURL(shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if url == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	if time.Now().After(url.ExpireAt) {
		c.JSON(http.StatusGone, gin.H{"error": "URL has expired"})
		return
	}

	// redirect
	c.Redirect(http.StatusFound, url.LongURL)
}
