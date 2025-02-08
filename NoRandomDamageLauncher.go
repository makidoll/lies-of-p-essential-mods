// go build -o NoRandomDamageLauncher NoRandomDamageLauncher.go

package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getProtonPath() (string, error) {
	for _, arg := range os.Args[1:] {
		if strings.HasSuffix(arg, "proton") {
			return arg, nil
		}
	}
	return "", errors.New("Not found")
}

func main() {
	// rootDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	gameArgs := os.Args[1:]

	protonPath, err := getProtonPath()
	if err != nil {
		os.Exit(1)
	}

	// run game

	gameCmd := exec.Command(gameArgs[0])
	gameCmd.Args = gameArgs
	// gameCmd.Stdin = os.Stdin
	gameCmd.Stdout = os.Stdout
	gameCmd.Stderr = os.Stderr
	gameCmd.Start()
	gameCmd.Process.Release()

	// run mod

	time.Sleep(5 * time.Second)

	modArgs := []string{protonPath, "run", "NoRandomDamage.exe"}
	// _ = os.WriteFile("lop.txt", []byte(strings.Join(modArgs, " ")), 0644)

	// will inject whenever game is found though
	modCmd := exec.Command(modArgs[0])
	modCmd.Args = modArgs
	modCmd.Stdin = os.Stdin
	modCmd.Stdout = os.Stdout
	modCmd.Stderr = os.Stderr
	modCmd.Start()
	modCmd.Process.Release()

	// both processes are released so this program will close
	// steam's stop button however breaks unfortunately
	// tried to pipe signals and shit but it was too complex
}
