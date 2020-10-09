package config

import (
	"flag"
	"os"
)

type Config struct {
	Port   string
	Listen string
}

func Get() *Config {
	config := &Config{}


	// Read in the environment variables
	flag.StringVar(&config.Port, "port", os.Getenv("SERVER_PORT"), "Port server will run on")
	flag.StringVar(&config.Listen, "listen", os.Getenv("LISTEN_ADDRESS"), "Where to listen for requests")

	flag.Parse()

	return config
}
