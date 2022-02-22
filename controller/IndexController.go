package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexControllerS struct {
	Index func(context *gin.Context)
}

var IndexController IndexControllerS

func init() {

	IndexController.Index = func(context *gin.Context) {
		context.String(http.StatusOK, "index")
	}

}
