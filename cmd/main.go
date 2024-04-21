package main

import (
	"fmt"
	"shop/internal/config"
)

func main() {
	cfg := config.MustGetConfig("./config")
	fmt.Println(cfg.JWT)
}
