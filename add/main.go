package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	dateFormat = "20060102-030405"
	usage      = `
# == Adds a database upgrade script to this repository.
#
# == Usage
#  sem-add <path>
#
# == Example
#  sem-add ./new-script.sql
#
`
	targetDir = "scripts"
)

func must1[T any](v T, err error) T {
	must(err)
	return v
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mustRelDir(dir string) {
	parent := must1(os.Getwd())
	scriptsDir := filepath.Join(parent, dir)
	must(os.MkdirAll(scriptsDir, os.ModePerm))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "**** ERROR: Need file path")
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}
	file := os.Args[1]
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "File[%s] could not be found\n", file)
		os.Exit(1)
	}
	if filepath.Ext(file) != ".sql" {
		fmt.Fprintf(os.Stderr, "File[%s] must end with .sql\n", file)
		os.Exit(1)
	}

	mustRelDir(targetDir)

	now := time.Now().UTC().Format(dateFormat)
	targetFile := now + ".sql"
	target := filepath.Join(targetDir, targetFile)

	for _, err := os.Stat(target); !errors.Is(err, os.ErrNotExist); time.Sleep(1 * time.Millisecond) {
		now = time.Now().UTC().Format(dateFormat)
		target = filepath.Join(targetDir, targetFile)
	}

	fmt.Printf("Adding %s\n", target)
	must(os.Rename(file, target))

	must(exec.Command("git", "add", target).Run())
	fmt.Println("File staged in git. You need to commit and push")
}
