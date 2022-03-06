package main

import (
	"github.com/TheMightyGit/dinos/cartridge"
	"github.com/TheMightyGit/marv/marvlib"
)

func main() {
	marvlib.API.ConsoleBoot(
		"dinos",
		cartridge.Resources,
		cartridge.Start,
		cartridge.Update,
	)
}
