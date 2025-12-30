package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rbrick/elections/internal/env"
)

func init() {
	godotenv.Load()
}

func main() {
	app := gin.Default()

	app.GET("/:country/:region/", func(ctx *gin.Context) {})

	app.Run(env.GetOrDefault("HOST", ":3000"))
}
