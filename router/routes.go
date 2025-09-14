package router

func (r *Router) Routes() {
	r.setupCORS(r.Engine)

	r.Engine.POST("/createurl", r.CreateURL)
	r.Engine.GET("/:shortURL", r.RedirectURL)
}
