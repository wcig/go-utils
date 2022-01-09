package xuuid

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// new uuid
func NewUuid() string {
	return uuid.NewString()
}

// new uuid with no dashes
func NewShortUuid() string {
	return strings.ReplaceAll(NewUuid(), "-", "")
}

// valid uuid
func IsValidUuid(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
