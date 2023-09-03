package utils

import (
	"log"
	"os/exec"
	"strings"
)

func GenStringUUID() string {
	cmd := exec.Command("uuidgen")
	uuidBytes, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Convert the bytes to a string and remove leading/trailing whitespace
	uuid := strings.TrimSpace(string(uuidBytes))
	return uuid
}
