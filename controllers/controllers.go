package controllers

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Jéssica",
	})
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API DIZ: ": "E aí " + nome + ", tudo beleza?!",
	})
}
