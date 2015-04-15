package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "googl"
	app.Usage = "short and expand url with Google URL Shortener"
	app.Version = VERSION
	app.Author = "Diego Moreira"
	app.Email = "diegoalvesmoreira@gmail.com"
	app.Commands = Commands
	app.Run(os.Args)
}
