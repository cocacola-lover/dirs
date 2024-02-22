package main

import (
	"dirs/pkg/logger"
	"dirs/pkg/thinker"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logger := logger.NewLogger()

	thinker.InitThinker(logger)
}
