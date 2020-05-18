package http

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Server binds the types necessary to serve the Nemo application.
type Server struct {
	*echo.Echo

	db *gorm.DB

	bindStr string
}

// New returns an initialized server, but the db must still be
// injected.
func New(d *gorm.DB) (*Server, error) {
	s := &Server{
		Echo: echo.New(),
		db:   d,
	}

	s.bindStr = os.Getenv("NEMO_BIND")
	if s.bindStr == "" {
		s.bindStr = ":1234"
	}

	return s, nil
}

// Serve initializes the server and never returns.
func (s *Server) Serve() error {
	s.POST("/series", s.seriesCreate)
	s.GET("/series", s.listSeries)

	s.POST("/series/:sid/events", s.eventCreate)
	s.GET("/series/:sid/events", s.eventList)
	s.GET("/series/:sid/events/:eid", s.eventFetch)
	s.DELETE("/series/:sid/events/:eid", s.eventArchive)

	return s.Start(s.bindStr)
}
