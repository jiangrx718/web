package middleware

import (
	capture "web/gopkg/gins/capture"
	"web/gopkg/utils"

	"github.com/gin-gonic/gin"
)

func RequestCapture() gin.HandlerFunc {
	return capture.RequestCapture(capture.Options{
		FilterPaths: []string{
			"/api/file/upload-file-proxy",
		},
	}, func(ctx *gin.Context, request *capture.Request) {
		ctx.Set(utils.ClientIPKey, request.ClientIP)
	}, func(ctx *gin.Context, capture *capture.Capture) {
		//log.Sugar().Info(ctx, "request", zap.Any("capture", capture))
	})
}
