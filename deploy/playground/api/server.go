package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/iczc/billboard/playground/config"
	"github.com/iczc/billboard/playground/internal"
)

type Server struct {
	cfg      *config.Config
	engine   *gin.Engine
	verifier *internal.Verifier
}

func NewServer(cfg *config.Config) *Server {
	server := &Server{
		cfg:      cfg,
		verifier: internal.NewFlagVerifier(cfg.LCD, cfg.CheckMode),
	}
	gin.SetMode(gin.ReleaseMode)
	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/flag", s.getFlagByTxHash)

	s.engine = r
}

func (s *Server) Run() {
	log.Fatal(s.engine.Run(":" + s.cfg.Port))
}
