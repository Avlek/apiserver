package main

import (
	"flag"
	"github.com/Avlek/apiserver/impl"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config.toml", "path to config")
}

func main() {
	flag.Parse()
	config := impl.NewConfig()
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}
	server := impl.NewAPIServer(config)
	if err = server.Start(); err != nil {
		log.Fatal(err)
	}
}
