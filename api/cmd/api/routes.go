package main

import "github.com/gin-gonic/gin"

func (a *Application) routes() *gin.Engine {

	router := gin.Default()

	router.Handle("GET", "/equestrian-events", listEvents)

	return router
}
