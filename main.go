package main

import (
	"fmt"
	"github.com/marcuschui2022/blog-aggregator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	cfg.SetUser("lane")
	fmt.Printf("%+v", cfg)
}
