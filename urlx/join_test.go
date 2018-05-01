package urlx

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestAppendPaths(t *testing.T) {
	u, _ := url.Parse("https://localhost/1234/")
	assert.Equal(t, "https://localhost/1234/foo", AppendPaths(u, "/foo").String())
}
