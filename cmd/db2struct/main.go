package main

import (
	"os"

	"github.com/urfave/cli"
	db2struct "github.com/seiji/go-db2struct"
)

var (
	version string = "HEAD"
)

func main() {
	newApp().Run(os.Args)
}

func newApp() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "db2struct"
	app.Usage = "Generate go codes from db tables"
	app.Version = version
	app.Action = db2struct.Generate
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "omitempty, o",
			Usage: "Add field tag of omitempty",
		},
		cli.StringFlag{
			Name:  "package, p",
			Value: "db",
			Usage: "Name for this package",
		},
	}
	return
}
