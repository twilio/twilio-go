package client

import (
	"net/http"
	"net/url"
	"time"
)

// MockBaseClient is a hand-written test double for BaseClient.
//
// Each method delegates to its corresponding Fn field when set, and otherwise
// returns the zero value, so a test only has to populate the behavior it cares
// about:
//
//	c := NewMockBaseClient()
//	c.SendRequestFn = func(method, rawURL string, data url.Values,
//		headers map[string]interface{}, body ...byte) (*http.Response, error) {
//		return &http.Response{Body: io.NopCloser(strings.NewReader(`{}`))}, nil
//	}
//
// Calls are counted so a test can assert on how the client was exercised.
type MockBaseClient struct {
	AccountSidFn  func() string
	SetTimeoutFn  func(timeout time.Duration)
	SendRequestFn func(method string, rawURL string, data url.Values,
		headers map[string]interface{}, body ...byte) (*http.Response, error)
	SetOauthFn func(auth OAuth)
	OAuthFn    func() OAuth

	AccountSidCalls  int
	SetTimeoutCalls  int
	SendRequestCalls int
	SetOauthCalls    int
	OAuthCalls       int
}

// NewMockBaseClient returns a MockBaseClient with no behavior configured.
func NewMockBaseClient() *MockBaseClient {
	return &MockBaseClient{}
}

func (m *MockBaseClient) AccountSid() string {
	m.AccountSidCalls++
	if m.AccountSidFn == nil {
		return ""
	}
	return m.AccountSidFn()
}

func (m *MockBaseClient) SetTimeout(timeout time.Duration) {
	m.SetTimeoutCalls++
	if m.SetTimeoutFn != nil {
		m.SetTimeoutFn(timeout)
	}
}

func (m *MockBaseClient) SendRequest(method string, rawURL string, data url.Values,
	headers map[string]interface{}, body ...byte) (*http.Response, error) {
	m.SendRequestCalls++
	if m.SendRequestFn == nil {
		return nil, nil
	}
	return m.SendRequestFn(method, rawURL, data, headers, body...)
}

func (m *MockBaseClient) SetOauth(auth OAuth) {
	m.SetOauthCalls++
	if m.SetOauthFn != nil {
		m.SetOauthFn(auth)
	}
}

func (m *MockBaseClient) OAuth() OAuth {
	m.OAuthCalls++
	if m.OAuthFn == nil {
		return nil
	}
	return m.OAuthFn()
}

// Compile-time assertion that the double satisfies the interface it stands in for.
var _ BaseClient = (*MockBaseClient)(nil)
