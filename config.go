package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
}

func (c *Config) Load() {
	c.Port = os.Getenv("PORT")
	fmt.Println("Config load")
}
