package stringsx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitNonEmpty(t *testing.T) {
	// assert.Len(t, strings.Split("", " "), 1)
	assert.Len(t, Splitx("", " "), 0)
}
