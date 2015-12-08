# torctl

[![GoDoc](https://godoc.org/github.com/crackcomm/torctl?status.svg)](https://godoc.org/github.com/crackcomm/torctl)

This is a TOR proxies service that implements [tower-services/proxies](https://github.com/tower-services/proxies) interface.

## Install

To install service and control command line tool:

```
go install github.com/crackcomm/torctl/cmd/torctl
```

## Usage

```
NAME:
   torctl - tor service control

USAGE:
   torctl [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   launch	launches a tor proxy
   service	starts a tor proxy service
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

### torctl launch

```
NAME:
   torctl launch - launches a tor proxy

USAGE:
   torctl launch [command options] [arguments...]

OPTIONS:
   --tor-bin "bin/linux/amd64/tor"	Tor executable [$TOR_BIN]
   --data-dir 				Tor data directory (default: temporary directory is created if empty) [$TOR_DATA_DIR]
   --torrc 				Tor config (by default we generate one) [$TOR_TORRC]
   --proxy-address "127.0.0.1"		Tor proxy listening address [$TOR_PROXY_ADDR]
   --proxy-port "0"			Tor proxy listening port [$TOR_PROXY_PORT]
   --control-address "127.0.0.1"	Tor control listening address [$TOR_CTL_ADDR]
   --control-port "0"			Tor control listening port [$TOR_CTL_PORT]
   --control-password-hash 		Hashed tor control password [$TOR_CTL_PASS_HASH]
   --verbose				runs in verbose mode [$TOR_VERBOSE]
```

### torctl service

```
NAME:
   torctl service - starts a tor proxy service

USAGE:
   torctl service [command options] [arguments...]

OPTIONS:
   --verbose				runs in verbose mode [$TOR_VERBOSE]
   --maxmind-db 			MaxMind Countries database [$MAXMIND_DB]
   --tor-bin "bin/linux/amd64/tor"	Tor executable [$TOR_BIN]
   --temp-dir 				Temporary directory to create tor data-dir [$TOR_TEMP_DIR]
   --control-password-hash 		Hashed tor control password [$TOR_CTL_PASS_HASH]
   --listen "0.0.0.0"			tor service server listening address [$LISTEN]
   --port "9055"			tor service server listening port [$PORT]
   --limit "50"				Tor instances limit [$INSTANCES_LIMIT]
   --new-nym-interval "100ms"		tor new nym signal interval (new IP) [$NEW_NYM_INTERVAL]
```
