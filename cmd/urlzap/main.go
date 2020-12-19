package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/brunoluiz/urlzap"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"exec", "e"},
				Usage:   "Generate static content for redirecting to URLs",
				Action:  run(generate),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "config, c",
						Value: "./config.yml",
						Usage: "Site configs",
					},
				},
			},
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "Serve config and return HTTP 301 for configured keys",
				Action:  run(serve),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "config, c",
						Value: "./config.yml",
						Usage: "Site configs",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(cb func(c *cli.Context, config urlzap.Config) error) cli.ActionFunc {
	return func(c *cli.Context) error {
		freader, err := os.Open(c.String("config"))
		if err != nil {
			return err
		}

		conf, err := urlzap.FromYAML(freader)
		if err != nil {
			return err
		}

		return cb(c, conf)
	}
}

func generate(c *cli.Context, conf urlzap.Config) error {
	if conf.Path == "" {
		conf.Path = "./"
	}

	return urlzap.Generate(c.Context, conf)
}

func serve(c *cli.Context, conf urlzap.Config) error {
	if conf.HTTP.Address == "" || conf.HTTP.BasePath == "" {
		return errors.New("Missing http.address or http.path configs")
	}

	return http.ListenAndServe(conf.HTTP.Address, urlzap.NewServer(c.Context, conf))
}
