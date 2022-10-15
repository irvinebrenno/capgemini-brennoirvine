package sequence

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("sequence", finderSequence)
	r.GET("/stats", getStats)
}
