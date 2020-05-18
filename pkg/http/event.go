package http

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/the-maldridge/nemo/pkg/models"
)

func (s *Server) eventCreate(c echo.Context) error {
	event := &models.Event{}
	if err := c.Bind(event); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, event)
}

func (s *Server) eventFetch(c echo.Context) error {
	return nil
}

func (s *Server) eventList(c echo.Context) error {
	return nil
}

func (s *Server) eventArchive(c echo.Context) error {
	return nil
}
