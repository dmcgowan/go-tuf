package main

import (
	"github.com/flynn/go-tuf"
	"github.com/flynn/go-tuf/Godeps/_workspace/src/github.com/flynn/go-docopt"
)

func init() {
	register("clean", cmdClean, `
usage: tuf clean

Remove all staged manifests.
  `)
}

func cmdClean(args *docopt.Args, repo *tuf.Repo) error {
	return repo.Clean()
}
