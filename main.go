package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rbrick/elections/internal/data/sources/us/texas"
	"github.com/rbrick/elections/internal/env"
)

func init() {
	godotenv.Load()
}

func main() {
	app := gin.Default()

	// Gets available elections
	app.GET("/:country/:region/elections", func(ctx *gin.Context) {
		country := ctx.Param("country")
		region := ctx.Param("region")

		if country == "us" && region == "texas" {
			texasTurnoutSource := texas.NewTexasTurnoutSource()
			elections, err := texasTurnoutSource.GetElections()

			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(200, elections)
			return
		}

		ctx.JSON(404, gin.H{"error": "source not found"})
	})

	app.Run(env.GetOrDefault("HOST", ":3000"))
}
