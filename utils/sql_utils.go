package utils

import (
	"fmt"
	"strings"
)

func GenerateSQLPlaceholders(count int) (string, []interface{}) {
	placeholders := make([]string, count)
	args := make([]interface{}, count)

	for i := 0; i < count; i++ {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = nil
	}

	return strings.Join(placeholders, ","), args
}
