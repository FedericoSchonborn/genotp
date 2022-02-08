package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "gotp: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <SECRET>\n", filepath.Base(args[0]))
		os.Exit(2)
	}

	code, err := totp.GenerateCode(args[1], time.Now())
	if err != nil {
		return err
	}

	fmt.Println(code)
	return nil
}
