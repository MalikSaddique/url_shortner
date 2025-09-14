package router

import (
	"net/http"

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

	// redirect
	c.Redirect(http.StatusFound, url.LongURL)
}
