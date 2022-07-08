package main

import "fmt"

type Config struct {
	Port string
}

func (c Config) Load() {
	c.Port = ":8080"
	fmt.Println("Config load")
}
