package urlx

import (
	"net/url"
	"path"
)

// AppendPaths appends the provided paths to the url.
func AppendPaths(u *url.URL, paths ...string) (ep *url.URL) {
	ep = Copy(u)
	ep.Path = path.Join(append([]string{ep.Path}, paths...)...)
	return ep
}
