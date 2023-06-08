package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cmd/go/internal/cfg"
	"cmd/go/internal/mod/fetch"

	"golang.org/x/mod/v2/module"
)

func usage() {
	fmt.Fprintf(os.Stderr, `usage: dl path[/subdirs] version

Path is to the root of a git repository.
A tag of the form "subdir/version" must exist.
Version is a semver of the form "v1.2.3".

This is using the Golang module system and infrastructure (see https://go.dev/ref/mod)
By default the GOPROXY 'https://proxy.golang.org,direct' is used.
For private repo
 1. set GOPRIVATE
    e.g., export GOPRIVATE=github.com/org1/repo,github.com/org2/repo
 2. add "url with insteadOf" entry to ~/.gitconfig.
    e.g.,
[url "git@github.com:org1/repo.git"]
  insteadOf = https://github.com/org1repo.git

The "url with insteadOf" entry is needed as all downloads are tried using https which requires a username and password.
The instead of allows the use of ssh private key.
Note if you have multiple ssh keys a url like "git@github.com-org1:org1/repo.git" and associated entry in ~/.ssh/config is useful.


Note GOPRIVATE (or GONOPROXY) can be used for repos which can be downloaded but not through the gomod proxy (e.g., don't contain go code).

Flags:
`)
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	Debug = flag.Bool("x", false, `prints commands as they are executed.
This is useful for debugging version control commands when a module is downloaded directly
from a repository`)
	ModFilename = flag.String("modfilename", "adl.bundle.json",
		`The filename which contains the module (package, crate, bundle) information, specifically the global identifier.
For Golang this should be "go.mod" with a 'module path' entry.
Golang mod format and json are supported.
For json files of the form <name.json> an element name "module" is required.
Json files with the form <a.b.XXX.json> an element name XXX is required.`)
)

func main() {
	ctx := context.Background()

	flag.Usage = usage
	flag.Parse()

	if *Debug {
		cfg.BuildX = true
	}

	if flag.NArg() != 2 {
		usage()
	}

	dir, err := modfetch.Download(ctx, module.Version{
		Path:        flag.Arg(0),
		ModFilename: *ModFilename,
		Version:     flag.Arg(1),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "?%s\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Downloaded to %s\n", dir)
	fmt.Printf("%s\n", dir)

}
