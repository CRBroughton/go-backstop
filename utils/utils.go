package utils

import (
	"fmt"
	"os/exec"
)

func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	err := cmd.Run()

	if IsError(err) {
		return fmt.Errorf("failed to run Docker command: %v", err)
	}
	return nil

}
