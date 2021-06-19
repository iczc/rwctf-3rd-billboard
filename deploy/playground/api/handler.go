package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Err  string `json:"err"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func resp(context *gin.Context, err, msg, data string) {
	resp := Resp{
		Err:  err,
		Msg:  msg,
		Data: data,
	}

	context.JSON(http.StatusOK, resp)
}

func (s *Server) getFlagByTxHash(context *gin.Context) {
	token := context.Query("token")
	txHash := context.Query("tx")
	if len(txHash) == 0 || len(txHash) != 64 {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := s.verifier.ValidateTx(txHash, token); err != nil {
		resp(context, err.Error(), "", "")
		return
	}

	resp(context, "", "", s.cfg.Flag)
}
