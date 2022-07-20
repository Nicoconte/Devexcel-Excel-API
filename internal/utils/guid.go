package utils

import (
	"strings"

	"github.com/pborman/uuid"
)

func NewGuid() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}
