package main

import (
	"os"
	"github.com/codegangsta/cli"
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
