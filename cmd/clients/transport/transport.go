package transport

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"time"
)

func ConfigureDefaultHttpTransport() (*http.Transport, *http.Client) {

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	return t, httpClient
}

func ConfigureInSecureHttpTransport() (t *http.Transport, c *http.Client) {

	t, c = ConfigureDefaultHttpTransport()
	log.Println(c)
	t.TLSClientConfig = &tls.Config{
		Rand: nil,
		Time: func() time.Time {
			return time.Now()
		},
		Certificates:      []tls.Certificate{},
		NameToCertificate: map[string]*tls.Certificate{},
		GetCertificate: func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
			return nil, nil
		},
		GetClientCertificate: func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return nil, nil

		},
		GetConfigForClient: func(*tls.ClientHelloInfo) (*tls.Config, error) {
			return nil, nil

		},
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			return nil

		},
		VerifyConnection: func(tls.ConnectionState) error {
			return nil
		},
		RootCAs:                     &x509.CertPool{},
		NextProtos:                  []string{},
		ServerName:                  "",
		ClientAuth:                  0,
		ClientCAs:                   &x509.CertPool{},
		InsecureSkipVerify:          true,
		CipherSuites:                []uint16{},
		PreferServerCipherSuites:    false,
		SessionTicketsDisabled:      false,
		SessionTicketKey:            [32]byte{},
		ClientSessionCache:          nil,
		MinVersion:                  0,
		MaxVersion:                  0,
		CurvePreferences:            []tls.CurveID{},
		DynamicRecordSizingDisabled: false,
		Renegotiation:               0,
		KeyLogWriter:                nil,
	}
	c = &http.Client{
		Transport: t,
	}

	return t, c

}

func ConfigureSecureHttpTransport() (*http.Transport, *http.Client) {

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	return t, httpClient
}
