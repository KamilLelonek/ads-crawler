package publisher

import (
	"fmt"
	"strings"
)

func contains(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}

func relationship(r string) Relationship {
	switch {
	case contains(r, "DIRECT"):
		return Direct
	case contains(r, "RESELLER"):
		return Reseller
	default:
		return Other
	}
}

func normalize(row string) []string {
	fields := strings.Split(row, ",")

	for i := range fields {
		fields[i] = strings.TrimSpace(fields[i])
	}

	return fields
}

func removeComment(element string) string {
	i := strings.Index(element, "#")

	switch i {
	case -1:
		return element
	default:
		return element[:i-1]
	}
}

func buildRow(fields []string) Row {
	var r Row

	r.Domain = fields[0]
	r.AccountId = fields[1]
	r.Relationship = relationship(fields[2])

	if len(fields) > 3 {
		r.Authority = removeComment(fields[3])
	}

	return r
}

func NewRow(row string) (Row, error) {
	fields := normalize(row)

	if len := len(fields); len != 3 && len != 4 {
		return Row{}, fmt.Errorf("The given row is invalid: %s", row)
	}

	r := buildRow(fields)

	return r, nil
}
