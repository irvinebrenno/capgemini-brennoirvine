package main

import (
	"capgemini-brennoirvine/webservice/sequence"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// definimos os grupos de rotas
	sequenceGroup := r.Group("")

	sequence.Router(sequenceGroup)

	r.Run() // rodando em localhost:8080
}
