package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/the-maldridge/nemo/pkg/models"
)

func (s *Server) seriesCreate(c echo.Context) error {
	series := &models.Series{}
	if err := c.Bind(series); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := s.db.Save(series); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusCreated)
}

func (s *Server) listSeries(c echo.Context) error {
	qp := c.QueryParam("archived")
	archived, err := strconv.ParseBool(qp)
	if err != nil {
		archived = false
	}

	series := []models.Series{}
	if err := s.db.Where(&models.Series{Archived: archived}).Find(&series); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, series)
}
