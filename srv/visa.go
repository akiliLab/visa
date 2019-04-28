package visa

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"net/http"
)

const (
	productionBaseURL = "api.visa.com"
	stagingBaseURL    = "stage.api.visa.com"
	devBaseURL        = "dev.api.visa.com"
	sandboxBaseURL    = "sandbox.api.visa.com"
	userAgent         = "go-visa"
)

// A Client manages communication with the Visa API.
type Client struct { // HTTP client used to communicate with the API.
	HTTPClient *http.Client // Oauth 1 client to set authentication up

	// Base URL for API requests. Defaults to the public Visa API.
	BaseURL string

	// User agent used when communicating with the Visa API.
	UserAgent string

	Authorization string
}

// EnvType equals the type of environment to work in
type EnvType int

// Environment types
const (
	PRODUCTION = 0
	STAGING    = 1
	DEV        = 2
	SANDBOX    = 3
)

// BaseURL Returns a string that corresponds to the environment type
func (e EnvType) BaseURL() string {
	switch e {
	case PRODUCTION:
		return productionBaseURL
	case STAGING:
		return stagingBaseURL
	case DEV:
		return devBaseURL
	case SANDBOX:
		return sandboxBaseURL
	}
	return sandboxBaseURL
}

// NewClient returns an instance of a Mastercard api client that allows
// accessing endpoints with more ease. Oauth1 authentification is managed
// internally, but you need to pass your mastercard consumerKey, the path
// to the .p12 file and the keystore password for the client to retrieve
// your RSA private key and sign requests correctly
func NewClient(userID string, password string, cert tls.Certificate, ca []byte, env EnvType) (*Client, error) {

	client := &Client{
		BaseURL:       env.BaseURL(),
		UserAgent:     userAgent,
		Authorization: "Basic " + basicAuth(userID, password),
	}

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(ca)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      clientCertPool,
	}

	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{TLSClientConfig: tlsConfig}

	client.HTTPClient = &http.Client{Transport: transport}

	return client, nil
}

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
