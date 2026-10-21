package client

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"net/http"
	"sync"
)

// distrustedSPKIHashes holds the hex-encoded SHA-256 hashes of the SubjectPublicKeyInfo of eight intermediate
// certificates that are no longer trustworthy: their private keys are weak, so anyone can use them to issue a
// certificate that a default trust store accepts. Go distrusts the same list as of Go 1.27.3 and 1.26.10
// (CVE-2026-111943, golang/go#82345); checking it here also protects applications built with older Go releases.
//
// Copied from Go commit 7f86f957633ee31a1c17a1eb8452ca7a3d9bc141 (src/crypto/x509/verify.go).
var distrustedSPKIHashes = map[string]struct{}{
	"ebed6475c2c2d80fcb199f3fe8dd6e6b3ce60b77eb15d4e42c3dec0c6f5b5105": {},
	"49dd353dc2b919a3770645117d0e60c16115e7d1bc259eefe734b00ea434fe5f": {},
	"e0254c7fcdc265e2f0f8a428dabb75ab4ae039b4872d62f5ee57f0877d5a8a41": {},
	"5c0660a2b164910d8ff589c18e14552e9e0ec237dcefe2cb736d41a4811f81a0": {},
	"43bffa19d6ef651451f6841e019931e96f6f2eaeff28ab4b6f1ce4576de2717b": {},
	"63ecefee431d49523bc1136c4e72d0206a477086caf606cd5cf025686c109cf8": {},
	"ec4ffa1ad0ac88081314b7aaffef6a91f97cd57291bb5ba686c5415f5eb4f6b6": {},
	"92c063d67e8f14db71d288e11266ffdaf69cbf115f7bab14f3bc2f68d25e0e30": {},
}

var errDistrustedIntermediate = errors.New("certificate chain contains a distrusted intermediate certificate")

func containsDistrustedSPKIHash(chain []*x509.Certificate) bool {
	for _, c := range chain {
		h := sha256.Sum256(c.RawSubjectPublicKeyInfo)
		if _, ok := distrustedSPKIHashes[hex.EncodeToString(h[:])]; ok {
			return true
		}
	}
	return false
}

// verifyConnection rejects a connection when every verified chain contains a distrusted certificate. Like Go, it
// allows the connection when at least one verified chain avoids them, and does nothing when verification was
// skipped (InsecureSkipVerify).
func verifyConnection(cs tls.ConnectionState) error {
	if len(cs.VerifiedChains) == 0 {
		return nil
	}
	for _, chain := range cs.VerifiedChains {
		if !containsDistrustedSPKIHash(chain) {
			return nil
		}
	}
	return errDistrustedIntermediate
}

// distrustTransport is the default HTTP client's transport. For each request it uses a copy of
// http.DefaultTransport with verifyConnection added. The copy is shared across requests so connections are still
// pooled, and it's rebuilt if http.DefaultTransport is replaced. When http.DefaultTransport isn't an
// *http.Transport (a test mock or an instrumentation wrapper, for example), it's used unchanged, as before.
type distrustTransport struct{}

var (
	distrustMu     sync.Mutex
	distrustSource *http.Transport
	distrustCopy   *http.Transport
)

func (distrustTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return currentDefaultTransport().RoundTrip(req)
}

func (distrustTransport) CloseIdleConnections() {
	if t, ok := currentDefaultTransport().(interface{ CloseIdleConnections() }); ok {
		t.CloseIdleConnections()
	}
}

func currentDefaultTransport() http.RoundTripper {
	source, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return http.DefaultTransport
	}
	distrustMu.Lock()
	defer distrustMu.Unlock()
	if source != distrustSource {
		if distrustCopy != nil {
			distrustCopy.CloseIdleConnections()
		}
		distrustSource = source
		distrustCopy = withDistrustCheck(source)
	}
	return distrustCopy
}

func withDistrustCheck(source *http.Transport) *http.Transport {
	// Clone sets up HTTP/2 the way the source would have, so the copy keeps using it.
	t := source.Clone()
	if t.TLSClientConfig == nil {
		t.TLSClientConfig = &tls.Config{}
	}
	next := t.TLSClientConfig.VerifyConnection
	t.TLSClientConfig.VerifyConnection = func(cs tls.ConnectionState) error {
		if err := verifyConnection(cs); err != nil {
			return err
		}
		if next != nil {
			return next(cs)
		}
		return nil
	}
	return t
}
