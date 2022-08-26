package main

import (
	"errors"
	"os"

	"github.com/github/git-bundle-server/internal/core"
)

type Delete struct{}

func (Delete) subcommand() string {
	return "delete"
}

func (Delete) run(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: git-bundle-server delete <route>")
	}

	route := args[0]

	repo, err := core.CreateRepository(route)
	if err != nil {
		return err
	}

	err = core.RemoveRoute(route)
	if err != nil {
		return err
	}

	err = os.RemoveAll(repo.WebDir)
	if err != nil {
		return err
	}

	err = os.RemoveAll(repo.RepoDir)
	if err != nil {
		return err
	}

	return nil
}
