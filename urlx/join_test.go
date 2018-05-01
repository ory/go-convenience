package urlx

import (
	"testing"
	"net/url"
	"github.com/stretchr/testify/assert"
)

func TestAppendPaths(t *testing.T) {
	u, _ := url.Parse("https://localhost/1234/")
	assert.Equal(t, "https://localhost/1234/foo", AppendPaths(u, "/foo").String())
}
