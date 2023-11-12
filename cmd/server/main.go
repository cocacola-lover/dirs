package main

import (
	"dirs/pkg/butler"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	butler.InitButler()
}
