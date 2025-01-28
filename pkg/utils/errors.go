package utils

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
	"regexp"
)

func HandleUniqueConstraintError(err error) (int, string, bool) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		// Extract column name and value for the detail message
		re := regexp.MustCompile(`Key \(([^)]+)\)=\([^)]*\) already exists.`)
		matches := re.FindStringSubmatch(pgErr.Detail)
		if len(matches) == 2 {
			columnName := matches[1]
			return http.StatusConflict, fmt.Sprintf("%s already exists", columnName), true
		}
		return http.StatusConflict, fmt.Sprintf("%s already exists", pgErr.ConstraintName), true
	}
	return 0, "", false
}