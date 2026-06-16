// urlzap is a CLI tool for generating redirect pages and serving HTTP redirects.
package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/brunoluiz/urlzap"
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
		slog.Error(err.Error()) //nolint:gosec // CLI args, not untrusted input
		os.Exit(1)
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

	return (&http.Server{
		Addr:         conf.HTTP.Address,
		Handler:      urlzap.NewServer(c.Context, conf),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}).ListenAndServe()
}
