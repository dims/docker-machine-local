package main

import (
	"github.com/dims/docker-machine-local"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(local.NewDriver("default", "path"))
}
