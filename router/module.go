package router

import (
	"time"

	"database/sql"

	"github.com/MalikSaddique/url_shortner/api"
	"github.com/MalikSaddique/url_shortner/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	URLShortnerAPI api.URLShortnerAPI
	DBLayer        db.URLDB
	Engine         *gin.Engine
}

func NewRouter(database *sql.DB) *Router {
	dbLayer := db.NewURLDB(database)

	apiInput := api.NewURLShortnerAPIInput{
		URLdb: dbLayer,
	}

	apiLayer := api.NewURLShortnerAPI(apiInput)
	engine := gin.Default()

	routes := &Router{
		URLShortnerAPI: apiLayer,
		DBLayer:        dbLayer,
		Engine:         engine,
	}
	routes.Routes()
	return routes
}

func (r *Router) setupCORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
