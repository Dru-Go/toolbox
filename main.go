package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dru-go/noah-toolbox/adapter/repository"
	"github.com/dru-go/noah-toolbox/constant"
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/dru-go/noah-toolbox/usecase"
	"github.com/urfave/cli/v2"
)

var load bool
var dump bool
var file string

func main() {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	mu := usecase.MaterialUsecase{
		Repo: repository.NewRepository(db),
		Ctx:  context.Background(),
	}
	tu := usecase.TransactionUsecase{
		Repo: repository.NewRepository(db),
		Ctx:  context.Background(),
	}
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
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "name of the material",
							},
							&cli.StringFlag{
								Name:  "category",
								Usage: "category of the material",
							},
							&cli.StringFlag{
								Name:  "measurement",
								Usage: "measurement of the material",
							},
						},
						Action: func(ctx *cli.Context) error {
							fmt.Println("added material: ", ctx.App.UsageText)
							new_material, err := mu.Create(ctx.String("name"), ctx.String("category"), ctx.String("measurement"))
							fmt.Println(new_material)
							return err
						},
					},
					{
						Name:  "import",
						Usage: "import transaction from a source",
						Action: func(ctx *cli.Context) error {
							filePath := ctx.String("file")
							category := ctx.String("category")
							if filePath == "" {
								return fmt.Errorf("file path was not provided")
							}
							if category == "" {
								return fmt.Errorf("category was not provided")
							}
							return mu.BulkImport(ctx.String("file"), ctx.String("category"))
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "category",
								Aliases:  []string{"c"},
								Usage:    "category of the materials being imported",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "file",
								Aliases:  []string{"f"},
								Usage:    "path of the file to be imported",
								Required: true,
							},
						},
					},
					{
						Name:    "load",
						Aliases: []string{"l"},
						Usage:   "load materials from a csv file",
						Action: func(ctx *cli.Context) error {
							return mu.LoadCSV(ctx.String("file"))
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "file",
								Usage:    "path of the file to be imported",
								Required: true,
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
						Action: func(ctx *cli.Context) error {
							var transactions []domain.Transaction
							var err error
							var filePath = ctx.String("file")
							var materialId = ctx.String("material")
							var transactionId = ctx.String("transaction")
							if filePath != "" {
								transactions, err = tu.LoadCSV(filePath)
								if err != nil {
									return err
								}
							} else if materialId != "" {
								transactions, err = tu.Repo.Fetch(domain.ComputeFilter{MaterialId: materialId})
								if err != nil {
									return err
								}
							} else if transactionId != "" {
								transactions, err = tu.Repo.FetchSubsequentTransactions(transactionId)
								if err != nil {
									return err
								}
							} else {
								return fmt.Errorf("you need to provide a file path for this command to work, --f")
							}
							tu.BulkCompute(transactions)
							return err
						},
						Flags: []cli.Flag{

							&cli.StringFlag{
								Name:    "file",
								Aliases: []string{"f"},
								Usage:   "path of the file to be imported",
							},
						},
					},
					{
						Name:    "load",
						Aliases: []string{"l"},
						Usage:   "load transaction from a csv file",
						Action: func(ctx *cli.Context) error {
							tu.LoadCSV(ctx.String("file"))
							return nil
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "file",
								Usage:    "source file to the csv",
								Required: true,
							},
						},
					},
					{
						Name:      "compute",
						UsageText: ``,
						Action: func(ctx *cli.Context) error {
							var transactions []domain.Transaction
							var err error
							var materialId = ctx.String("material")
							var transactionId = ctx.String("transaction")
							var company = ctx.String("company")
							var project = ctx.String("project")
							if materialId != "" {
								if company == "" || project == "" {
									return fmt.Errorf("you have added the material, so you need to provide the company id and the project id")
								}
								transactions, err = tu.Repo.Fetch(domain.ComputeFilter{MaterialId: materialId, Company: company, Project: project})
								if err != nil {
									return err
								}
							} else if transactionId != "" {
								transactions, err = tu.Repo.FetchSubsequentTransactions(transactionId)
								if err != nil {
									return err
								}
							} else {
								return fmt.Errorf("you need to provide a file path for this command to work, --f")
							}
							if len(transactions) == 0 {
								return fmt.Errorf("unable to find transactions for the provided filters")
							}
							computedTransactions := tu.BulkCompute(transactions)
							if ctx.Bool("update") {
								tu.BulkUpdate(computedTransactions)
							}
							return nil
						},
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "update",
								Aliases: []string{"u"},
								Usage:   "Update the transactions in the datastore",
							},
							&cli.StringFlag{
								Name:     "material",
								Aliases:  []string{"m"},
								Usage:    "materialId for the transactions",
								Required: false,
							},
							&cli.StringFlag{
								Name:     "project",
								Aliases:  []string{"p"},
								Usage:    "project identifier for the transactions",
								Required: false,
							},
							&cli.StringFlag{
								Name:     "company",
								Aliases:  []string{"c"},
								Usage:    "company identifier for the transactions",
								Required: false,
							},
							&cli.StringFlag{
								Name:     "transaction",
								Aliases:  []string{"t"},
								Usage:    "materialId for the transactions",
								Required: false,
							},
						},
					},
				},
			},
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "serve the server locally port :3400",
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
