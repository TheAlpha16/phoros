package utils

import (
	"os"
	"log"
	"errors"
	"regexp"
	"strings"
)

func ValidateFileName(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("filename is empty")
	}

	filename = strings.ToLower(filename)
	
	re := regexp.MustCompile(`[^a-z0-9._-]`)
	filename = re.ReplaceAllString(filename, "_")

	if len(filename) > 255 {
		filename = filename[:255]
	}

	return filename, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}