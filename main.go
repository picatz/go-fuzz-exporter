package main

import (
	"fmt"
	"os"

	"github.com/picatz/go-fuzz-exporter/pkg/server"
)

func givenStdinPipe() (bool, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, fmt.Errorf("failed to stat process stdin: %w", err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true, nil
	}
	return false, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("unexpected number of args (%d), expected %d\n", len(os.Args), 2)
		os.Exit(1)
	}

	ok, err := givenStdinPipe()
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("no stdin pipe")
	}

	server.Start(os.Args[1], os.Stdin)
}
