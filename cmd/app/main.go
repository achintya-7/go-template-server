package main

import (
	"fmt"

	"github.com/achintya-7/go-template-server/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Port)
}
