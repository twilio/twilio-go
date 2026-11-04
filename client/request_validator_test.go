package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testURL   = "https://mycompany.com/myapp.php?foo=1&bar=2"
	signature = "vOEb5UThFn24KEfnOFLQY2AE5FY=" // of the testURL above with the params below
	bodyHash  = "0a1ff7634d9ab3b95db5c9a2dfe9416e41502b283a80c7cf19632632f96e6620"
)

var (
	validator = NewRequestValidator("12345")
	params    = map[string]string{
		"Digits":                "1234",
		"CallSid":               "CA1234567890ABCDE",
		"To":                    "+18005551212",
		"Caller":                "+14158675309",
		"From":                  "+14158675309",
		"ReasonConferenceEnded": "test",
		"Reason":                "Participant",
	}
	jsonBody = []byte(`{"property": "value", "boolean": true}`)
	formBody = []byte(`property=value&boolean=true`)
)

func TestRequestValidator_Validate(t *testing.T) {
	t.Parallel()

	t.Run("returns true when validation succeeds", func(t *testing.T) {
		assert.True(t, validator.Validate(testURL, params, signature))
	})

	t.Run("returns false when validation fails", func(t *testing.T) {
		assert.False(t, validator.Validate(testURL, params, "WRONG SIGNATURE"))
	})

	t.Run("returns true when https and port is specified but signature is generated without it", func(t *testing.T) {
		theURL := strings.Replace(testURL, ".com", ".com:1234", 1)
		assert.True(t, validator.Validate(theURL, params, signature))
	})

	t.Run("returns true when https and port is specified and signature is generated with it", func(t *testing.T) {
		expectedSignature := "vOEb5UThFn24KEfnOFLQY2AE5FY=" // hash of https uri without port
		assert.True(t, validator.Validate(testURL, params, expectedSignature))
	})

	t.Run("returns true when http and port port is specified but signature is generated without it", func(t *testing.T) {
		theURL := strings.Replace(testURL, ".com", ".com", 1)
		theURL = strings.Replace(theURL, "https", "http", 1)
		expectedSignature := "n2xBNyzSW7rfYStDtOFiFMv7qNo=" // hash of http uri without port
		assert.True(t, validator.Validate(theURL, params, expectedSignature))
	})

	t.Run("returns true when http and port is specified and signature is generated with it", func(t *testing.T) {
		theURL := strings.Replace(testURL, ".com", ".com:1234", 1)
		theURL = strings.Replace(theURL, "https", "http", 1)
		expectedSignature := "n2xBNyzSW7rfYStDtOFiFMv7qNo=" // hash of http uri with port 1234
		assert.True(t, validator.Validate(theURL, params, expectedSignature))
	})

	t.Run("return false when params are sorted incorrectly", func(t *testing.T) {
		incorrectSignature := "95+Bu0JVPi0r/SsESZCVf0dWAjw=" //Params ReasonConferenceEnded is sorted before Reason
		assert.False(t, validator.Validate(testURL, params, incorrectSignature))
	})
}

func TestRequestValidator_ValidateBody(t *testing.T) {
	t.Parallel()

	t.Run("returns true when validation succeeds with json body", func(t *testing.T) {
		theURL := testURL + "&bodySHA256=" + bodyHash
		signatureWithBodyHash := "a9nBmqA0ju/hNViExpshrM61xv4="
		assert.True(t, validator.ValidateBody(theURL, jsonBody, signatureWithBodyHash))
	})

	t.Run("returns true when validation succeeds with form body", func(t *testing.T) {
		expectedSignature := "NBdBDr/T/lgjI+tlgpXjKZQZs/k="
		assert.True(t, validator.ValidateBody(testURL, formBody, expectedSignature))
	})

	t.Run("returns false when validation fails with json body", func(t *testing.T) {
		assert.False(t, validator.ValidateBody(testURL, jsonBody, signature))
	})

	t.Run("returns true when there's no other parameters and the signature is right", func(t *testing.T) {
		theURL := "https://mycompany.com/myapp.php?bodySHA256=" + bodyHash
		signatureForURL := "y77kIzt2vzLz71DgmJGsen2scGs="
		assert.True(t, validator.ValidateBody(theURL, jsonBody, signatureForURL))
	})
}

func sha256Signature(data string) string {
	h := hmac.New(sha256.New, []byte("12345"))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func TestRequestValidator_ValidateWithAlgorithm(t *testing.T) {
	// Signed string: the URL, then every param name and value sorted by name.
	signed := testURL + "CallSidCA1234567890ABCDE" + "Caller+14158675309" + "Digits1234" + "From+14158675309" +
		"ReasonParticipant" + "ReasonConferenceEndedtest" + "To+18005551212"
	sha256Sig := sha256Signature(signed)

	t.Run("matches an independently computed HMAC-SHA256 signature", func(t *testing.T) {
		assert.True(t, validator.ValidateWithAlgorithm(testURL, params, sha256Sig, HmacSha256))
		assert.True(t, validator.ValidateWithAlgorithm(testURL, params, sha256Sig, "hmac_sha256"))
	})

	t.Run("SHA1 and SHA256 signatures are not interchangeable", func(t *testing.T) {
		assert.False(t, validator.Validate(testURL, params, sha256Sig))
		assert.False(t, validator.ValidateWithAlgorithm(testURL, params, signature, HmacSha256))
	})

	t.Run("defaults to HMAC-SHA1 when the algorithm is empty", func(t *testing.T) {
		assert.True(t, validator.ValidateWithAlgorithm(testURL, params, signature, ""))
		assert.True(t, validator.ValidateWithAlgorithm(testURL, params, signature, HmacSha1))
	})

	t.Run("returns false and logs for an unsupported algorithm", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer log.SetOutput(os.Stderr)

		assert.False(t, validator.ValidateWithAlgorithm(testURL, params, signature, "HMAC_SHA512"))
		assert.Contains(t, buf.String(), "HMAC_SHA512")
	})

	t.Run("validates a json body signed with HMAC-SHA256", func(t *testing.T) {
		theURL := testURL + "&bodySHA256=" + bodyHash
		sig := sha256Signature(theURL)
		assert.True(t, validator.ValidateBodyWithAlgorithm(theURL, jsonBody, sig, HmacSha256))
		assert.False(t, validator.ValidateBodyWithAlgorithm(theURL, []byte("{}"), sig, HmacSha256))
		assert.False(t, validator.ValidateBody(theURL, jsonBody, sig))
	})

	t.Run("validates a form body signed with HMAC-SHA256", func(t *testing.T) {
		sig := sha256Signature(testURL + "booleantrue" + "propertyvalue")
		assert.True(t, validator.ValidateBodyWithAlgorithm(testURL, formBody, sig, HmacSha256))
	})
}

func TestGetSignatureAlgorithm(t *testing.T) {
	t.Run("reads the algorithm from the settings header", func(t *testing.T) {
		assert.Equal(t, HmacSha256, GetSignatureAlgorithm("key=webhooks_sharedkey_1,algorithm=HMAC_SHA256"))
		assert.Equal(t, HmacSha1, GetSignatureAlgorithm("algorithm=HMAC_SHA1,key=k"))
		assert.Equal(t, HmacSha256, GetSignatureAlgorithm(" key=k , algorithm = HMAC_SHA256 "))
	})

	t.Run("defaults to HMAC-SHA1 when missing, empty or without algorithm", func(t *testing.T) {
		for _, header := range []string{"", "key=k", "key=k,algorithm="} {
			assert.Equal(t, HmacSha1, GetSignatureAlgorithm(header))
		}
	})

	t.Run("exposes the header names", func(t *testing.T) {
		assert.Equal(t, "X-Twilio-Signature", SignatureHeader)
		assert.Equal(t, "X-Twilio-Webhooks-Signature-Settings", SignatureSettingsHeader)
	})
}
