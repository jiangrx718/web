package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"web/commands"
	"web/commands/server"
	"web/gopkg/log"
)

func main() {
	app := cli.NewApp()
	app.Action = server.Run
	app.Before = server.InitConfig
	app.After = server.Flush
	app.Commands = commands.All()
	app.Flags = server.Flags()
	if err := app.Run(os.Args); err != nil {
		log.Sugar().Fatal(err)
	}
}
