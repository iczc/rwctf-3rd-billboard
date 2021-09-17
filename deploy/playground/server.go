package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg      *Config
	engine   *gin.Engine
	verifier *Verifier
}

func NewServer(cfg *Config) *Server {
	server := &Server{
		cfg:      cfg,
		verifier: NewFlagVerifier(cfg.LCD, cfg.CheckMode),
	}
	gin.SetMode(gin.ReleaseMode)
	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/flag", s.handleGetFlag)

	s.engine = r
}

func (s *Server) Run() {
	log.Fatal(s.engine.Run(":" + s.cfg.Port))
}

type request struct {
	Token  string `form:"token"`
	TxHash string `form:"tx" binding:"required,len=64"`
}

type response struct {
	Err  string `json:"err"`
	Data string `json:"data"`
}

func (s *Server) handleGetFlag(context *gin.Context) {
	var req request
	if err := context.BindQuery(&req); err != nil {
		log.Println(err)
		return
	}

	if err := s.verifier.ValidateTx(req.Token, req.TxHash); err != nil {
		context.JSON(http.StatusOK, response{Err: err.Error()})
		return
	}

	log.Println(req.Token, req.TxHash)
	context.JSON(http.StatusOK, response{Data: s.cfg.Flag})
}
