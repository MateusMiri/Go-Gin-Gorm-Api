package main

import (
	"api-avaliacao/src/config"
	"api-avaliacao/src/database"
	"api-avaliacao/src/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	database.ConnectDB()
	r := gin.Default()
	router.Rotas(r)
	r.Run(fmt.Sprintf(":%d", config.Porta))

}
