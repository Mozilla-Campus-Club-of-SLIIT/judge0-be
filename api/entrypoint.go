package api

import (
	"net/http"

	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/routes"
	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	api := app.Group("/api")
	routes.RegisterAllRoutes(api)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
