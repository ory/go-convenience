package urlx

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"github.com/stretchr/testify/require"
	"fmt"
)

func TestAppendPaths(t *testing.T) {
	u, err := url.Parse("http://localhost/home/")
	require.NoError(t, err)
	assert.Equal(t, "http://localhost/home/", AppendPaths(u).String())

	for k, tc := range []struct {
		give   []string
		expect string
	}{
		{
			give:   []string{"http://localhost/", "/home"},
			expect: "http://localhost/home",
		},
		{
			give:   []string{"http://localhost", "/home"},
			expect: "http://localhost/home",
		},
		{
			give:   []string{"https://localhost/", "/home"},
			expect: "https://localhost/home",
		},
		{
			give:   []string{"http://localhost/", "/home", "home/", "/home/"},
			expect: "http://localhost/home/home/home/",
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			u, err := url.Parse(tc.give[0])
			require.NoError(t, err)
			assert.Equal(t, tc.expect, AppendPaths(u, tc.give[1:]...).String())
		})
	}
}
