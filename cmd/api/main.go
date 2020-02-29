package main

import (
	"github.com/tanookiben/benkyo/api"
	"github.com/tanookiben/benkyo/internal/config"
)

func main() {
	c := config.Read()
	api.NewAPI(c).Serve()
}
