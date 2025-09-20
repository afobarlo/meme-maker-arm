package main

import (
	"log"
	"meme-maker/cmd"
	
	imagick "gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	log.SetFlags(0)
	cmd.Execute()
}
