package client

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testCertificate(t *testing.T, cn string) *x509.Certificate {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: cn},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	assert.NoError(t, err)
	cert, err := x509.ParseCertificate(der)
	assert.NoError(t, err)
	return cert
}

// distrust adds cert's SPKI hash to the list for the rest of the test.
func distrust(t *testing.T, cert *x509.Certificate) {
	h := sha256.Sum256(cert.RawSubjectPublicKeyInfo)
	hash := hex.EncodeToString(h[:])
	distrustedSPKIHashes[hash] = struct{}{}
	t.Cleanup(func() { delete(distrustedSPKIHashes, hash) })
}

// replaceDefaultTransport swaps http.DefaultTransport for the rest of the test.
func replaceDefaultTransport(t *testing.T, transport http.RoundTripper) {
	original := http.DefaultTransport
	http.DefaultTransport = transport
	t.Cleanup(func() { http.DefaultTransport = original })
}

// countingTLSServer returns a TLS test server and a transport that trusts its certificate.
func countingTLSServer(t *testing.T) (*httptest.Server, *int32, *http.Transport) {
	var requests int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requests, 1)
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)
	roots := x509.NewCertPool()
	roots.AddCert(server.Certificate())
	return server, &requests, &http.Transport{TLSClientConfig: &tls.Config{RootCAs: roots}}
}

func get(t *testing.T, c *Client, url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	assert.NoError(t, err)
	res, err := c.doWithErr(req)
	if res != nil {
		res.Body.Close()
	}
	return res, err
}

func TestVerifyConnection(t *testing.T) {
	leaf := testCertificate(t, "leaf")
	bad := testCertificate(t, "distrusted intermediate")
	good := testCertificate(t, "intermediate")
	distrust(t, bad)

	tests := []struct {
		name    string
		chains  [][]*x509.Certificate
		wantErr bool
	}{
		{"verification skipped", nil, false},
		{"clean chain", [][]*x509.Certificate{{leaf, good}}, false},
		{"distrusted intermediate", [][]*x509.Certificate{{leaf, bad}}, true},
		{"distrusted leaf", [][]*x509.Certificate{{bad, good}}, true},
		{"one clean alternative", [][]*x509.Certificate{{leaf, bad}, {leaf, good}}, false},
		{"every chain distrusted", [][]*x509.Certificate{{leaf, bad}, {leaf, good, bad}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := verifyConnection(tls.ConnectionState{VerifiedChains: tt.chains})
			if tt.wantErr {
				assert.ErrorIs(t, err, errDistrustedIntermediate)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDefaultClientRefusesDistrustedChain(t *testing.T) {
	server, requests, transport := countingTLSServer(t)
	replaceDefaultTransport(t, transport)
	distrust(t, server.Certificate())

	_, err := get(t, &Client{Credentials: NewCredentials("user", "pass")}, server.URL)

	assert.ErrorIs(t, err, errDistrustedIntermediate)
	assert.Equal(t, int32(0), atomic.LoadInt32(requests), "the request must not reach the server")
}

func TestClientWithTimeoutRefusesDistrustedChain(t *testing.T) {
	server, requests, transport := countingTLSServer(t)
	replaceDefaultTransport(t, transport)
	distrust(t, server.Certificate())
	c := &Client{Credentials: NewCredentials("user", "pass")}
	c.SetTimeout(5 * time.Second)

	_, err := get(t, c, server.URL)

	assert.ErrorIs(t, err, errDistrustedIntermediate)
	assert.Equal(t, int32(0), atomic.LoadInt32(requests))
}

func TestDefaultClientAllowsCleanChain(t *testing.T) {
	server, requests, transport := countingTLSServer(t)
	replaceDefaultTransport(t, transport)

	res, err := get(t, &Client{Credentials: NewCredentials("user", "pass")}, server.URL)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, int32(1), atomic.LoadInt32(requests))
}

// The default client's transport is a copy of http.DefaultTransport with an extra hook, so check it kept hostname
// verification: a trusted, non-distrusted certificate issued for another host must still be refused.
func TestDefaultClientRejectsHostnameMismatch(t *testing.T) {
	rootKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	rootTemplate := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "Test Root"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		IsCA:                  true,
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageCertSign,
	}
	rootDER, err := x509.CreateCertificate(rand.Reader, rootTemplate, rootTemplate, &rootKey.PublicKey, rootKey)
	assert.NoError(t, err)
	root, err := x509.ParseCertificate(rootDER)
	assert.NoError(t, err)

	leafKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	leafDER, err := x509.CreateCertificate(rand.Reader, &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: "example.com"},
		DNSNames:     []string{"example.com"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}, root, &leafKey.PublicKey, rootKey)
	assert.NoError(t, err)

	var requests int32
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requests, 1)
	}))
	server.TLS = &tls.Config{Certificates: []tls.Certificate{{Certificate: [][]byte{leafDER}, PrivateKey: leafKey}}}
	server.StartTLS()
	t.Cleanup(server.Close)
	roots := x509.NewCertPool()
	roots.AddCert(root)
	replaceDefaultTransport(t, &http.Transport{TLSClientConfig: &tls.Config{RootCAs: roots}})

	_, err = get(t, &Client{Credentials: NewCredentials("user", "pass")}, server.URL)

	var hostnameErr x509.HostnameError
	assert.True(t, errors.As(err, &hostnameErr), "expected a hostname error, got %v", err)
	assert.NotErrorIs(t, err, errDistrustedIntermediate)
	assert.Equal(t, int32(0), atomic.LoadInt32(&requests), "the request must not reach the server")
}

func TestCustomHTTPClientIsUnchanged(t *testing.T) {
	server, requests, _ := countingTLSServer(t)
	distrust(t, server.Certificate())
	c := &Client{Credentials: NewCredentials("user", "pass"), HTTPClient: server.Client()}

	_, err := get(t, c, server.URL)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), atomic.LoadInt32(requests))
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestDefaultClientUsesReplacedDefaultTransport(t *testing.T) {
	// Test mocks and instrumentation often replace http.DefaultTransport; the default client must still use it.
	called := false
	replaceDefaultTransport(t, roundTripFunc(func(req *http.Request) (*http.Response, error) {
		called = true
		return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody, Request: req}, nil
	}))

	res, err := get(t, &Client{Credentials: NewCredentials("user", "pass")}, "https://api.twilio.com/")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.True(t, called)
}

func TestCurrentDefaultTransport(t *testing.T) {
	var nextCalled bool
	source := &http.Transport{TLSClientConfig: &tls.Config{
		VerifyConnection: func(tls.ConnectionState) error {
			nextCalled = true
			return nil
		},
	}}
	replaceDefaultTransport(t, source)

	first := currentDefaultTransport()
	assert.Same(t, first, currentDefaultTransport(), "the copy is reused, so connections are pooled")
	assert.NotSame(t, source, first, "http.DefaultTransport itself is not modified")

	copied := first.(*http.Transport)
	assert.NoError(t, copied.TLSClientConfig.VerifyConnection(tls.ConnectionState{}))
	assert.True(t, nextCalled, "an existing VerifyConnection still runs")

	replaceDefaultTransport(t, &http.Transport{})
	assert.NotSame(t, first, currentDefaultTransport(), "a replaced http.DefaultTransport is picked up")
}

func TestDefaultClientKeepsHTTP2(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	roots := x509.NewCertPool()
	roots.AddCert(server.Certificate())
	// Like http.DefaultTransport, which sets ForceAttemptHTTP2.
	replaceDefaultTransport(t, &http.Transport{TLSClientConfig: &tls.Config{RootCAs: roots}, ForceAttemptHTTP2: true})

	res, err := get(t, &Client{Credentials: NewCredentials("user", "pass")}, server.URL)

	assert.NoError(t, err)
	assert.Equal(t, 2, res.ProtoMajor)
}
