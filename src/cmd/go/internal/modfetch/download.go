//go:build ignore

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cmd/go/internal/cfg"
	"cmd/go/internal/modfetch"

	"golang.org/x/mod/module"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: download path[/subdirs] version\n")
	os.Exit(2)
}

func main() {
	cfg.BuildX = true
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 2 {
		usage()
	}
	ctx := context.Background()

	dir, err := modfetch.Download(ctx, module.Version{
		Path:    flag.Arg(0),
		Version: flag.Arg(1),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "?%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Downloaded to %s\n", dir)

}
