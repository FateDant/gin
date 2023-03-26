package study

import (
	"gin/logs_source"
	"github.com/gin-gonic/gin"
)

func LogTest(ctx *gin.Context) {
	logs_source.Log.Info("zheshi info")
}
