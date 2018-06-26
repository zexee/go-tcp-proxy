package main

import (
	"flag"
	"fmt"
	"net"
	"os"
  "encoding/json"
  "../../"
)

var (
	matchid = uint64(0)
	connid  = uint64(0)
	logger  proxy.ColorLogger

	localAddr   = flag.String("l", ":9999", "local address")
	remoteAddr  = flag.String("r", "localhost:80", "remote address")
	configFile  = flag.String("c", "/etc/tcp-proxy.json", "Config file")
	quiet       = flag.Bool("q", false, "display nothing")
	verbose     = flag.Bool("v", false, "display server actions")
	veryverbose = flag.Bool("vv", false, "display server actions and all tcp data")
	nagles      = flag.Bool("n", false, "disable nagles algorithm")
	hex         = flag.Bool("h", false, "output hex")
	colors      = flag.Bool("color", false, "output ansi colors")
	unwrapTLS   = flag.Bool("unwrap-tls", false, "remote connection with TLS exposed unencrypted locally")
)

type Config struct {
  Local string `json:"local"`
  Remote string `json:"remote"`
}

func main() {
	flag.Parse()

	logger := proxy.ColorLogger{
    Quiet:   false,
		Verbose: *verbose,
		Color:   *colors,
	}

  if len(*configFile) > 0 {
    logger.Info("Config file %v", *configFile)
    var config Config;
    f, err := os.Open(*configFile)
    defer f.Close()
    if err != nil {
      fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(f)
    jsonParser.Decode(&config)
    logger.Info(config.Local)
    logger.Info(config.Remote)
    *localAddr = config.Local
    *remoteAddr = config.Remote
  }

	logger.Info("Proxying from %v to %v", *localAddr, *remoteAddr)

	laddr, err := net.ResolveTCPAddr("tcp", *localAddr)
	if err != nil {
		logger.Warn("Failed to resolve local address: %s", err)
		os.Exit(1)
	}
	raddr, err := net.ResolveTCPAddr("tcp", *remoteAddr)
	if err != nil {
		logger.Warn("Failed to resolve remote address: %s", err)
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		logger.Warn("Failed to open local port to listen: %s", err)
		os.Exit(1)
	}

	if *veryverbose {
		*verbose = true
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logger.Warn("Failed to accept connection '%s'", err)
			continue
		}
		connid++

		var p *proxy.Proxy
		if *unwrapTLS {
			logger.Info("Unwrapping TLS")
			p = proxy.NewTLSUnwrapped(conn, laddr, raddr, *remoteAddr)
		} else {
			p = proxy.New(conn, laddr, raddr)
		}

		p.Nagles = *nagles
		p.OutputHex = *hex
		p.Log = proxy.ColorLogger{
			Quiet:       *quiet,
			Verbose:     *verbose,
			VeryVerbose: *veryverbose,
			Prefix:      fmt.Sprintf("Connection #%03d ", connid),
			Color:       *colors,
		}

		go p.Start()
	}
}

