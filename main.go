package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dru-go/noah-toolbox/constant"
	"github.com/urfave/cli/v2"
)

var load bool
var dump bool
var file string

func main() {
	app := &cli.App{
		Name:      "toolbox",
		Version:   "0.1",
		UsageText: constant.Toolbox_usage,
		Action: func(ctx *cli.Context) error {
			fmt.Print(ctx.App.UsageText)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "material",
				Aliases: []string{"m"},
				Usage:   "manage materials in the data store",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:    "add",
						Aliases: []string{"new", "create"},
						UsageText: `

						`,
						Action: func(ctx *cli.Context) error {
							fmt.Println("added material: ", ctx.App.UsageText)
							return nil
						},
					},
					{
						Name:  "import",
						Usage: "import transaction from a source",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("added material: ", cCtx.Args().First())
							return nil
						},
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "load",
								Usage:       "load source file to the terminal",
								Destination: &load,
							},
							&cli.BoolFlag{
								Name:        "dump",
								Usage:       "dump loaded materials to the datastore",
								Destination: &dump,
							},
							&cli.StringFlag{
								Name:        "file",
								Usage:       "path of the file to be imported",
								Destination: &file,
								Required:    true,
							},
						},
					},
				},
			},
			{
				Name:    "transaction",
				Aliases: []string{"t"},
				Usage:   "add a task to the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:  "import",
						Usage: "import transaction from a source",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("added material: ", cCtx.Args().First())
							return nil
						},
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "load",
								Usage:       "load source file to the terminal",
								Destination: &load,
							},
							&cli.BoolFlag{
								Name:        "dump",
								Usage:       "dump loaded materials to the datastore",
								Destination: &dump,
							},
							&cli.StringFlag{
								Name:        "file",
								Usage:       "path of the file to be imported",
								Destination: &file,
							},
						},
					},
					{
						Name:      "compute",
						UsageText: ``,
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "load",
								Aliases:     []string{"l"},
								Usage:       "load source file to the terminal",
								Destination: &load,
							},
							&cli.BoolFlag{
								Name:        "dump",
								Aliases:     []string{"d"},
								Usage:       "dump loaded materials to the datastore",
								Destination: &dump,
							},
							&cli.StringFlag{
								Name:        "file",
								Aliases:     []string{"f"},
								Usage:       "path of the file to be imported",
								Destination: &file,
								Required:    true,
							},
						},
						Action: func(ctx *cli.Context) error {
							fmt.Println("added material: ", ctx.App.UsageText)
							return nil
						},
					},
				},
			},
			{
				Name:    "report",
				Aliases: []string{"t"},
				Usage:   "options to generate report",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
