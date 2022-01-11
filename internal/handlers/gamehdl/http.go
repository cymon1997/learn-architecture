package gamehdl

import (
	"strconv"

	"github.com/cymon1997/go-architecture/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	gamesService ports.GamesService
}

func NewHTTPHandler(gamesService ports.GamesService) *HTTPHandler {
	return &HTTPHandler{
		gamesService: gamesService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	game, err := hdl.gamesService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, game)
}

func (hdl *HTTPHandler) Create(c *gin.Context) {
	name := c.Param("name")
	size, _ := strconv.ParseInt(c.Param("size"), 10, 64)
	bombs, _ := strconv.ParseInt(c.Param("bombs"), 10, 64)
	game, err := hdl.gamesService.Create(name, uint(size), uint(bombs))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, game)
}
