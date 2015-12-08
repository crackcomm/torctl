package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/golang/glog"
	"github.com/oschwald/geoip2-golang"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/crackcomm/onion/proxyutil"
	"github.com/crackcomm/torctl"
	"github.com/tower-services/proxies"
)

// Command - Command: torctl
var Command = cli.Command{
	Name:  "torctl",
	Usage: "tor service control",
	Subcommands: []cli.Command{
		LaunchCommand,
		ServiceCommand,
	},
}

// LaunchCommand - Command: torctl launch
var LaunchCommand = cli.Command{
	Name:  "launch",
	Usage: "launches a tor proxy",
	Flags: []cli.Flag{
		torBinFlag,
		cli.StringFlag{
			Name:   "data-dir",
			Usage:  "Tor data directory (default: temporary directory is created if empty)",
			EnvVar: "TOR_DATA_DIR",
		},
		cli.StringFlag{
			Name:   "torrc",
			Usage:  "Tor config (by default we generate one)",
			EnvVar: "TOR_TORRC",
		},
		cli.StringFlag{
			Name:   "proxy-address",
			Value:  "127.0.0.1",
			Usage:  "Tor proxy listening address",
			EnvVar: "TOR_PROXY_ADDR",
		},
		cli.IntFlag{
			Name:   "proxy-port",
			Usage:  "Tor proxy listening port",
			EnvVar: "TOR_PROXY_PORT",
		},
		cli.StringFlag{
			Name:   "control-address",
			Value:  "127.0.0.1",
			Usage:  "Tor control listening address",
			EnvVar: "TOR_CTL_ADDR",
		},
		cli.IntFlag{
			Name:   "control-port",
			Usage:  "Tor control listening port",
			EnvVar: "TOR_CTL_PORT",
		},
		torControlPasswordHash,
		verboseFlag,
	},
	Action: func(c *cli.Context) {
		opts := &torctl.LaunchOptions{
			Path:            c.String("tor-bin"),
			ProxyAddress:    c.String("proxy-address"),
			ProxyPort:       c.Int("proxy-port"),
			ControlAddress:  c.String("control-address"),
			ControlPort:     c.Int("control-port"),
			ControlPassword: c.String("control-password"),
			DataDir:         c.String("data-dir"),
			TempDir:         c.String("temp-dir"),
			Torrc:           c.String("torrc"),
			Quiet:           !c.Bool("verbose"),
		}
		if opts.ProxyPort == 0 {
			opts.ProxyPort = proxyutil.FreePort()
		}
		if opts.ControlPort == 0 {
			opts.ControlPort = proxyutil.FreePort()
		}
		service, err := torctl.Launch(context.Background(), opts)
		if err != nil {
			glog.Fatalf("Error running tor: %v", err)
		}

		endpoint := service.Endpoint()
		glog.Infof("Proxy address: %s", endpoint.Address)
		glog.Infof("Public address: %s", endpoint.PublicIp)

		<-make(chan bool)
	},
}

// ServiceCommand - Command: torctl service
var ServiceCommand = cli.Command{
	Name:  "service",
	Usage: "starts a tor proxy service",
	Flags: []cli.Flag{
		verboseFlag,
		maxMindDB,
		torBinFlag,
		tempDirFlag,
		torControlPasswordHash,
		cli.StringFlag{
			Name:   "listen",
			Value:  "0.0.0.0",
			Usage:  "tor service server listening address",
			EnvVar: "LISTEN",
		},
		cli.IntFlag{
			Name:   "port",
			Value:  9055,
			Usage:  "tor service server listening port",
			EnvVar: "PORT",
		},
		cli.IntFlag{
			Name:   "limit",
			Value:  50,
			Usage:  "Tor instances limit",
			EnvVar: "INSTANCES_LIMIT",
		},
		cli.DurationFlag{
			Name:   "new-nym-interval",
			Value:  100 * time.Millisecond,
			Usage:  "tor new nym signal interval (new IP)",
			EnvVar: "NEW_NYM_INTERVAL",
		},
	},
	Action: func(c *cli.Context) {
		torctl.NewNymInterval = c.Duration("new-nym-interval")
		addr := fmt.Sprintf("%s:%s", c.String("listen"), c.String("port"))

		// Maxmind DB
		var db *geoip2.Reader
		var err error
		if path := c.String("maxmind-db"); path != "" {
			db, err = geoip2.Open(path)
			if err != nil {
				glog.Fatal(err)
			}
			defer db.Close()
		}

		// Construct the TOR launcher
		launcher := torctl.NewLauncher(&torctl.LauncherOptions{
			Limit: c.Int("limit"),
			Defaults: &torctl.LaunchOptions{
				Path:                c.String("tor-bin"),
				TempDir:             c.String("temp-dir"),
				Quiet:               !c.Bool("verbose"),
				ControlPassword:     c.String("control-password-hash"),
				ControlPasswordHash: c.String("control-password-hash"),
				MaxMindDB:           db,
			},
		})

		// Start the tor service server listener
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			glog.Fatal(err)
		}

		// Construct the TOR gRPC service server
		server := grpc.NewServer()
		proxies.RegisterProxiesServer(server, launcher)

		glog.Infof("Serving on %s", listener.Addr())

		// Start serving
		err = server.Serve(listener)
		if err != nil {
			glog.Fatal(err)
		}
	},
}

var torBinFlag = cli.StringFlag{
	Name:   "tor-bin",
	Value:  "bin/linux/amd64/tor",
	Usage:  "Tor executable",
	EnvVar: "TOR_BIN",
}

var torControlPasswordHash = cli.StringFlag{
	Name:   "control-password-hash",
	Usage:  "Hashed tor control password",
	EnvVar: "TOR_CTL_PASS_HASH",
}

var maxMindDB = cli.StringFlag{
	Name:   "maxmind-db",
	Usage:  "MaxMind Countries database",
	EnvVar: "MAXMIND_DB",
}

var tempDirFlag = cli.StringFlag{
	Name:   "temp-dir",
	Usage:  "Temporary directory to create tor data-dir",
	EnvVar: "TOR_TEMP_DIR",
}

var verboseFlag = cli.BoolFlag{
	Name:   "verbose",
	Usage:  "runs in verbose mode",
	EnvVar: "TOR_VERBOSE",
}

func main() {
	defer glog.Flush()
	flag.CommandLine.Parse([]string{"-logtostderr"})

	app := cli.NewApp()
	app.Name = "torctl"
	app.HelpName = "torctl"
	app.Version = "0.0.0"
	app.Usage = "TOR control tool"
	app.Commands = []cli.Command{
		LaunchCommand,
		ServiceCommand,
	}

	app.Run(os.Args)
}
