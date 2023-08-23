package tlschecker

import (
	"crypto/tls"
	"fmt"
)

func CheckTLS(domain string) (string, string) {
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		return "", fmt.Sprintf("Error establishing TLS connection: %v", err)
	}
	defer conn.Close()

	// Get the state of the connection to retrieve version and cipher suite.
	state := conn.ConnectionState()

	version := tlsVersionToString(state.Version)
	cipherSuite := cipherSuiteToString(state.Version, state.CipherSuite)

	return version, cipherSuite
}

func tlsVersionToString(version uint16) string {
	switch version {
	case tls.VersionSSL30:
		return "SSLv3 (Considered insecure!)"
	case tls.VersionTLS10:
		return "TLS 1.0 (Consider upgrading.)"
	case tls.VersionTLS11:
		return "TLS 1.1 (Consider upgrading.)"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown"
	}
}

func cipherSuiteToString(tlsVersion uint16, suite uint16) string {
	var secureCipherSuitesTLS12 = map[uint16]bool{
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256: true,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384: true,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256: true,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:   true,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:   true,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256:   true,
	}

	var secureCipherSuitesTLS13 = map[uint16]bool{
		tls.TLS_AES_256_GCM_SHA384:       true,
		tls.TLS_CHACHA20_POLY1305_SHA256: true,
		tls.TLS_AES_128_GCM_SHA256:       true,
	}

	// Check if the suite is considered secure for the TLS version
	if tlsVersion == tls.VersionTLS12 && secureCipherSuitesTLS12[suite] {
		return tls.CipherSuiteName(suite)
	} else if tlsVersion == tls.VersionTLS13 && secureCipherSuitesTLS13[suite] {
		return tls.CipherSuiteName(suite)
	} else {
		return "Insecure or uncommon cipher suite"
	}
}
