package urlx

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	a, _ := url.Parse("http://localhost/1234")
	b := Copy(a)

	b.Path = "foo"
	assert.NotEqual(t, a.Path, b.Path)
}
