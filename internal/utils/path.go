package utils

import (
	"errors"
	"fmt"
	"os"
)

func GetStoragePath() (string, error) {

	fmt.Println("Datos ", Config)

	switch Config.Env {
	case "DOCKER":
		return Config.Storage, nil
	case "LOCAL":
		dir, _ := os.Getwd()
		return fmt.Sprintf("%s/../%s", dir, Config.Storage), nil
	default:
		return "", errors.New("Cannot get storage path. Reason: 'Invalid enviroment'")
	}
}

func DeleteFileFromStorage(path string) {
	os.Remove(path)
}
