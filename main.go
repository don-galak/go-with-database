package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/don-galak/go-with-database/apiserver"
	"github.com/don-galak/go-with-database/storage"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	apiServerAddrFlagName string = "addr"
	apiServerPostgresUrl  string = "postgres-url"
	apiServerNeo4jUrl     string = "neo4j-url"
)

func main() {
	if err := app().Run(os.Args); err != nil {
		logrus.WithError(err).Fatal("could not run application")
	}
}

func app() *cli.App {
	return &cli.App{
		Name:  "api-server",
		Usage: "The API",
		Commands: []*cli.Command{
			apiServerCmd(),
		},
	}
}

func apiServerCmd() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "starts the API server",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: apiServerAddrFlagName, EnvVars: []string{"API_SERVER_ADDR"}},
			&cli.StringFlag{Name: apiServerPostgresUrl, EnvVars: []string{"POSTGRES_DB_URL"}},
		},
		Action: func(c *cli.Context) error {

			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			stopper := make(chan struct{})
			go func() {
				<-done
				close(stopper)
			}()

			databaseURL := c.String(apiServerPostgresUrl)
			s, err := storage.NewStorage(databaseURL)
			if err != nil {
				return fmt.Errorf("could not initialize storage: %w", err)
			}

			addr := c.String(apiServerAddrFlagName)
			server, err := apiserver.NewAPIServer(addr, s)
			if err != nil {
				return err
			}

			return server.Start(stopper)
		},
	}
}
