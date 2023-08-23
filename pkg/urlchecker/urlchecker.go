package urlchecker

import (
	"net/url"
)

func IsValidURL(inputURL string) bool {
	u, err := url.Parse(inputURL)
	if err != nil {
		return false
	}

	// Only allowing HTTP or HTTPS protocols
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	// Optional: Check if domain is reachable
	// Please note: This check might not always be desired because a valid and well-formed URL might be temporarily unreachable.
	/*
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(u.Host, "80"), 5*time.Second)
		if err != nil {
			if strings.HasSuffix(u.Host, ":443") {
				conn, err = net.DialTimeout("tcp", net.JoinHostPort(u.Host, "443"), 5*time.Second)
				if err != nil {
					return false
				}
			} else {
				return false
			}
		}
		defer conn.Close()
	*/

	return true
}
