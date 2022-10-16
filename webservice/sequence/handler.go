package sequence

import (
	"capgemini-brennoirvine/application/sequence"
	"net/http"

	"github.com/gin-gonic/gin"
)

// finderSequence define um handler para encontrar uma sequência em uma matriz
func finderSequence(c *gin.Context) {
	reqSequence := sequence.Sequence{}

	if err := c.BindJSON(&reqSequence); err != nil {
		c.JSON(400, gin.H{
			"error": "Dados inválidos",
		})
		return
	}

	isValid, err := sequence.FinderSequence(reqSequence)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_valid": isValid,
	})

}

// getStats define um handler para buscar stats das matrizes testadas
func getStats(c *gin.Context) {
	res, err := sequence.GetStats()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, res)
}
