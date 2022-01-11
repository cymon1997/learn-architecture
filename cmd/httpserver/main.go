package main

import (
	"github.com/cymon1997/learn-architecture/internal/core/services/gamesrv"
	"github.com/cymon1997/learn-architecture/internal/handlers/gamehdl"
	"github.com/cymon1997/learn-architecture/internal/repositories/gamesrepo"
	"github.com/gin-gonic/gin"
)

func main() {
	gamesRepository := gamesrepo.NewMemKVS()
	gamesService := gamesrv.New(gamesRepository)
	gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)
	router.PUT("/games/:id", gamesHandler.RevealCell)

	router.Run(":8080")
}
