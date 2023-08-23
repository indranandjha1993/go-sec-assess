package headerchecker

import (
	"net/http"
	"time"
)

var headersToCheck = []string{
	"X-Content-Type-Options",
	"X-Frame-Options",
	"Content-Security-Policy",
	"X-XSS-Protection",
	"Strict-Transport-Security",
}

type HeaderResult struct {
	Name  string
	Value string
}

func CheckHeaders(url string) ([]HeaderResult, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var results []HeaderResult

	for _, header := range headersToCheck {
		value := resp.Header.Get(header)
		results = append(results, HeaderResult{
			Name:  header,
			Value: value,
		})
	}

	return results, nil
}
