package main

import (
	"strconv"

	"github.com/amalshaji/localport/internal/client/config"
	"github.com/amalshaji/localport/internal/constants"
	"github.com/urfave/cli/v2"
)

func tcpCmd() *cli.Command {
	return &cli.Command{
		Name:  "tcp",
		Usage: "Expose tcp port",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port to expose",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			port, err := strconv.Atoi(c.String("port"))
			if err != nil {
				return err
			}
			return startTunnels(c, &config.Tunnel{
				Port:      port,
				Subdomain: "",
				Type:      constants.Tcp,
			})
		},
	}
}
