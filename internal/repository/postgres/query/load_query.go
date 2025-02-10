package query

import (
	"embed"
	"fmt"
)

//go:embed sql-queries/*/*.sql
var queryFS embed.FS

func loadQuery(target, method string) (string, error) {
	data, err := queryFS.ReadFile(fmt.Sprintf("sql-queries/%s/%s.sql", target, method))
	if err != nil {
		return "", fmt.Errorf("failed to load query: %w", err)
	}
	return string(data), nil
}
