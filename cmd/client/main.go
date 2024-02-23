package main

import (
	envp "dirs/pkg/environment"
	"dirs/pkg/thinker"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	thinker.InitThinker(envp.NewEnvironment())
}
