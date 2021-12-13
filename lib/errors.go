package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var messages map[string]string = map[string]string{
	"90000001": "Internal Server Error.",
	"10000001": "Not found.",
	"10000002": "Already exist.",
}

var codes map[string]int = map[string]int{
	"90000001": http.StatusInternalServerError,
	"10000001": http.StatusNotFound,
	"10000002": http.StatusConflict,
}

type ErrorMessage struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

var SuccessfulMessage gin.H = gin.H{
	"code": "0",
}

func ErrorMaker(err error) (int, ErrorMessage) {
	code := codes[err.Error()]
	message := ErrorMessage{messages[err.Error()], err.Error()}
	if code == 0 {
		code = http.StatusBadRequest
	}
	if len(message.Error) == 0 {
		message = ErrorMessage{err.Error(), "80000001"}
	}
	return code, message
}
