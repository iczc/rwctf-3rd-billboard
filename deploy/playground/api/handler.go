package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Token  string `form:"token"`
	TxHash string `form:"tx" binding:"required,len=64"`
}

type response struct {
	Err  string `json:"err"`
	Data string `json:"data"`
}

func (s *Server) getFlagByTxHash(context *gin.Context) {
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
