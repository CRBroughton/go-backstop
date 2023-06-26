package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func RunCommand(command string, args ...string) (bool, error) {
	cmd := exec.Command(command, args...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()

	output := stdout.String()

	if strings.Contains(output, "Mismatch errors found") {
		// here return something useful to show
		// that the tests have failed
		log.Fatal("Mismatch errors found, check your HTML report")
		return true, nil
	}

	if IsError(err) {
		return true, fmt.Errorf("failed to run Docker command: %v", err)
	}

	return false, nil
}
