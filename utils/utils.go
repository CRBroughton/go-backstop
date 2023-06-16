package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func RunCommand(command string, loading *bool) bool {
	*loading = true
	result, err := exec.Command("bash", "-c", command).Output()

	if IsError(err) {
		log.Fatal("There was a fatal error running your command: ", err, result)
	}
	*loading = false
	return *loading
}
