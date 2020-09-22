package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/kinbiko/semver"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("expected exactly one argument of either 'major', 'minor', or 'patch'")
	}

	current, err := getLatestTag()
	if err != nil {
		log.Fatal(err)
	}

	newVersion, err := semver.Upversion(args[0], current)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newVersion)
}

func getLatestTag() (string, error) {
	out, err := exec.Command("git", "describe", "--abbrev=0", "--tags").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
