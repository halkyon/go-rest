package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/halkyon/go-rest-server/pkg/server"
)

const (
	exitFail = 1
)

func main() {
	if err := run(os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(stdout, stderr io.Writer) error {
	var serverConfig server.Config

	flag.StringVar(&serverConfig.Listen, "listen", "0.0.0.0", "Address")
	flag.StringVar(&serverConfig.Port, "port", "8000", "Port")
	flag.DurationVar(&serverConfig.IdleTimeout, "idle-timeout", 10*time.Second, "Idle timeout")
	flag.DurationVar(&serverConfig.ReadTimeout, "read-timeout", 10*time.Second, "Read timeout")
	flag.DurationVar(&serverConfig.WriteTimeout, "write-timeout", 10*time.Second, "Write timeout")
	flag.Int64Var(&serverConfig.MaxBodySize, "max-body-size", 2<<20, "Max body size in bytes")
	flag.BoolVar(&serverConfig.DebugPerf, "debug-perf", false, "Enable performance debugging")
	flag.Parse()

	server, err := server.New(serverConfig, stdout, stderr)
	if err != nil {
		return err
	}

	return server.Start()
}
