package urlx

import (
	"net/url"
	"path"
)

// AppendPaths appends the provided paths to the url.
func AppendPaths(u *url.URL, paths ...string) (ep *url.URL) {
	ep = Copy(u)
	if len(paths) == 0 {
		return ep
	}

	ep.Path = path.Join(append([]string{ep.Path}, paths...)...)

	last := paths[len(paths)-1]
	if last[len(last)-1] == '/' {
		ep.Path = ep.Path + "/"
	}

	return ep
}
