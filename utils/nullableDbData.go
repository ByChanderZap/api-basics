package utils

import (
	"database/sql"
	"time"
)

func NullableString(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

func NullableTime(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}
