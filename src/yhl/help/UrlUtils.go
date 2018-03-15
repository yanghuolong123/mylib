package help

import (
	"net/url"
)

func UrlEncode(ustr string) string {
	u, _ := url.Parse(ustr)
	u.RawQuery = u.Query().Encode()

	return u.String()
}
