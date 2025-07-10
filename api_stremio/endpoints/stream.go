package endpoints

import (
	"FSA/api_stremio/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterStreamEndpoint(e *echo.Echo) {
	e.GET("/stream/:type/:id.json", func(c echo.Context) error {
		cType := c.Param("type")
		ID := c.Param("id")

		println(cType, ID)

		stream := entities.Stream{
			Title: "The Matrix",
			URL:   "https://www.w3schools.com/html/mov_bbb.mp4",
		}

		streamResponse := entities.StreamResponse{
			Streams: []entities.Stream{stream},
		}

		return c.JSONPretty(http.StatusOK, streamResponse, " ")
	})
}
