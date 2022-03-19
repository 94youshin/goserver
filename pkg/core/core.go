package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmhk/goserver/pkg/errno"
)

// swagger: model
type ErrorResonpse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//
// WriteResponse write an error or the response data into http response body.
//
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		code, message := errno.DecodeErr(err)
		c.JSON(http.StatusOK, ErrorResonpse{Code: code, Message: message})
		return
	}
	c.JSON(http.StatusOK, data)
}
