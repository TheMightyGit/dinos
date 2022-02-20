package main

import (
	"log"

	"github.com/TheMightyGit/dinos/cartridge"
	"github.com/TheMightyGit/marv/marvlib"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	marvlib.API.ConsoleBoot(
		"dinos",
		cartridge.Resources,
		cartridge.Start,
		cartridge.Update,
	)
}
