package main

import (
	"fmt"
	"fullstack/backend/auth-ms/internal/config"
)

func main() {
	cfg := config.GetConfig()
	fmt.Print(cfg)
}
