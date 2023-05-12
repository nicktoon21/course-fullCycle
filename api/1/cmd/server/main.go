package main

import "github.com/nicktoon21/goexpert/apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBName)
}
