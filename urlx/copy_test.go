package urlx

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestCopy(t *testing.T) {
	a, _ := url.Parse("http://localhost/1234")
	b := Copy(a)

	b.Path = "foo"
	assert.NotEqual(t, a.Path, b.Path)
}
