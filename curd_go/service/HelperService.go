package service

import (
	"os/exec"
	"strings"
	"time"
)

func NewUuid() (string, error) {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(newUUID), "\n"), nil
}

func calculateAge(dob time.Time) int {
	today := time.Now()
	age := today.Year() - dob.Year()

	// Check if the birthday has occurred this year or not
	if today.YearDay() < dob.YearDay() {
		age--
	}

	return age
}
