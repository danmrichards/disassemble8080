package main

import (
	"flag"
	"log"

	"github.com/danmrichards/disassemble8080/internal/rom"
)

var romPath string

func main() {
	flag.StringVar(&romPath, "rom", "", "Path to 8080 ROM file")
	flag.Parse()

	rb, err := rom.Load(romPath)
	if err != nil {
		log.Fatal(err)
	}

	rom.Disassemble(rb)
}
