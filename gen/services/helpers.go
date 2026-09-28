package services

import (
	"fmt"
	"strings"
)

// requireValues reports the first name whose paired value is empty.
func requireValues(pairs ...string) error {
	for i := 0; i+1 < len(pairs); i += 2 {
		if pairs[i+1] == "" {
			return fmt.Errorf("blogger: %s is required", pairs[i])
		}
	}
	return nil
}

// joinList renders a list of strings as a comma-separated wire form.
func joinList(items []string) string {
	b := &strings.Builder{}
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		if b.Len() > 0 {
			b.WriteByte(',')
		}
		b.WriteString(item)
	}
	return b.String()
}

// itoa converts an int64 to its decimal string representation.
func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	var b [20]byte
	neg := n < 0
	if neg {
		n = -n
	}
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}
