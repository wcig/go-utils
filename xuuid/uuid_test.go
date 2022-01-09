package xuuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUuid(t *testing.T) {
	uuid := NewUuid()
	assert.Equal(t, 36, len(uuid))
}

func TestNewShortUuid(t *testing.T) {
	uuid := NewShortUuid()
	assert.Equal(t, 32, len(uuid))
}

func TestIsValidUuid(t *testing.T) {
	uuid := NewUuid()
	assert.True(t, IsValidUuid(uuid))
	assert.False(t, IsValidUuid(uuid+"-"))
}
