package main

import "fmt"

type Config struct {
	Port string
}

func (c Config) Load() {
	c.Port = ":8090"
	fmt.Println("Config load")
}
