package doc

import (
	"net/url"
	"strings"
)

func String(v string) string {

	s := strings.Split(v, "\n")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}

	return strings.TrimSpace(strings.Join(s, " "))
}

func MustURLJoinPath(base string, elem ...string) string {
	u, err := url.JoinPath(base, elem...)
	if err != nil {
		panic(err)
	}
	return u
}
