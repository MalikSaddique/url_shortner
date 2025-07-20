package router

func (r *Router) Routes() {
	r.setupCORS(r.Engine)

	r.Engine.POST("/createurl", r.CreateURL)
}
