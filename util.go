package xrpldata

import "time"

// FormatTime formats a time for passing to the XRPL Data API endpoints.
// https://xrpl.org/data-api.html#timestamps
func FormatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
