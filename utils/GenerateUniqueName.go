package utils

import (
	"path/filepath"

	"github.com/google/uuid"
)

func GenerateUniqueName(name string) string {
	id := uuid.New()

	extension := filepath.Ext(name)

	return id.String() + extension
}
