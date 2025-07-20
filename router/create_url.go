package router

import (
	"net/http"

	"github.com/MalikSaddique/url_shortner/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) CreateURL(c *gin.Context) {
	var req *models.URL
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	url := models.URL{
		Long_URL: req.Long_URL,
	}
	response, err := r.URLShortnerAPI.CreateURL(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}
