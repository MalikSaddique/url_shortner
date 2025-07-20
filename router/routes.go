package router

import "github.com/gin-gonic/gin"

func (r *Router) Routes(router *gin.Engine) {
	r.setupCORS(router)
}
