package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "gotp: %v", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	code, err := totp.GenerateCode(args[1], time.Now())
	if err != nil {
		return err
	}

	fmt.Println(code)
	return nil
}
