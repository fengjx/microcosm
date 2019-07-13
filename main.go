package main

import (
	"microcosm/internal/commands"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "microcosm"
	app.Usage = "a monitor server"
	app.Version = "v1.0"
	app.Copyright = "(c) 2018-2019 The microcosm <gzfengjianxin2012@gmail.com>"
	app.EnableBashCompletion = true
	app.Flags = commands.GlobalFlags

	app.Commands = []cli.Command{
		commands.StartCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}
