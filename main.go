package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/halkyon/go-rest/pkg/server"
)

const (
	exitFail = 1
)

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(stdout io.Writer) error {
	var serverConfig server.Config

	flag.StringVar(&serverConfig.Listen, "listen", "0.0.0.0", "HTTP server address")
	flag.StringVar(&serverConfig.Port, "port", "8000", "HTTP server port")
	flag.DurationVar(&serverConfig.IdleTimeout, "idle-timeout", 10*time.Second, "HTTP server idle timeout")
	flag.DurationVar(&serverConfig.ReadTimeout, "read-timeout", 10*time.Second, "HTTP server read timeout")
	flag.DurationVar(&serverConfig.WriteTimeout, "write-timeout", 10*time.Second, "HTTP server write timeout")
	flag.Int64Var(&serverConfig.MaxBodySize, "max-body-size", 51<<20, "Max body size in bytes")
	flag.BoolVar(&serverConfig.DebugPerf, "debug-perf", false, "Enable performance debugging")
	flag.Parse()

	stdoutLog := log.New(stdout, "", log.Ldate|log.Ltime)

	server, err := server.New(serverConfig, stdoutLog)
	if err != nil {
		return err
	}

	return server.Start()
}
