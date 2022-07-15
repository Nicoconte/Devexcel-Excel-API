package utils

import (
	"fmt"
	"os"
)

func GetStoragePath() string {
	dir, _ := os.Getwd()
	return fmt.Sprintf("%s/../%s", dir, Config.Storage)
}
