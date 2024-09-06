package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/thiagokokada/sem-go/internal/utils"
)

const (
	dateFormat = "20060102-030405"
	usage      = `
# == Adds a database upgrade script to this repository.
#
# == Usage
#  sem-add <path>...
#
# == Example
#  sem-add ./new-script.sql
#
`
	targetDir = "scripts"
)

func getTargetFile() string {
	now := time.Now().UTC().Format(dateFormat)
	return now + ".sql"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "**** ERROR: Need file path")
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}

	for _, file := range os.Args[1:] {
		if !utils.FileExist(file) {
			fmt.Fprintf(os.Stderr, "File[%s] could not be found\n", file)
			os.Exit(1)
		}
		if filepath.Ext(file) != ".sql" {
			fmt.Fprintf(os.Stderr, "File[%s] must end with .sql\n", file)
			os.Exit(1)
		}

		utils.Must(utils.MkRelDir(targetDir))

		target := filepath.Join(targetDir, getTargetFile())
		for utils.FileExist(target) {
			time.Sleep(1 * time.Millisecond)
			target = filepath.Join(targetDir, getTargetFile())
		}

		fmt.Printf("Adding %s\n", target)
		utils.Must(os.Rename(file, target))

		utils.Must(exec.Command("git", "add", target).Run())
		fmt.Println("File staged in git. You need to commit and push")
	}
}
