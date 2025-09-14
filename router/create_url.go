package router

import (
	"net/http"

	"github.com/MalikSaddique/url_shortner/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) CreateURL(c *gin.Context) {
	var req models.CreateURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	url := models.URL{
		LongURL:  req.LongURL,
		ExpireAt: req.ExpireAt,
	}
	response, err := r.URLShortnerAPI.CreateURL(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.CreateURLResponse{
		ShortURL: response.ShortURL,
		LongURL:  response.LongURL,
		ExpireAt: response.ExpireAt,
	})

}
