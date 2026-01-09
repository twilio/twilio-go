package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestPageUtil_ReadLimits(t *testing.T) {
	assert.Equal(t, 5, ReadLimits(nil, setLimit(5)))
	assert.Equal(t, 5, ReadLimits(setPageSize(10), setLimit(5)))
	assert.Equal(t, 1000, ReadLimits(nil, setLimit(5000)))
	assert.Equal(t, 10, ReadLimits(setPageSize(10), nil))
	assert.Equal(t, 50, ReadLimits(nil, nil))
}

func setLimit(limit int) *int {
	return &limit
}

func setPageSize(pageSize int) *int {
	return &pageSize
}

func TestPageUtil_GetNextPageUri(t *testing.T) {
	payload := map[string]interface{}{
		"next_page_uri": "/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1",
		"page_size":     50,
	}
	baseUrl := "https://api.twilio.com/"
	nextPageUrl, err := getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)

	payload["next_page_uri"] = "2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1"
	baseUrl = "https://api.twilio.com"
	nextPageUrl, err = getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)

	payload = map[string]interface{}{}
	nextPageUrl, err = getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "", nextPageUrl)
}

func TestPageUtil_GetNextPageUrl(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"next_page_url": "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1",
			"page_size":     50,
		},
	}

	nextPageUrl, err := getNextPageUrl("https://apitest.twilio.com", payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)
}

func getTestClient(t *testing.T) *MockBaseClient {
	mockCtrl := gomock.NewController(t)
	testClient := NewMockBaseClient(mockCtrl)
	testClient.EXPECT().AccountSid().DoAndReturn(func() string {
		return "AC222222222222222222222222222222"
	}).AnyTimes()

	testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}, body ...interface{}) (*http.Response, error) {
			response := map[string]interface{}{
				"end":            4,
				"first_page_uri": "/2010-04-01/Accounts/ACXX/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=2&Page=0",
				"messages": []map[string]interface{}{
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 0",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 1",
						"status":    "delivered",
					},
				},
				"uri":           "/2010-04-01/Accounts/ACXX/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=2&Page=0&PageToken=dummy",
				"page_size":     5,
				"start":         0,
				"next_page_uri": "/2010-04-01/Accounts/ACXX/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=2&Page=1&PageToken=PASMXX",
				"page":          0,
			}

			resp, _ := json.Marshal(response)

			return &http.Response{
				Body: io.NopCloser(bytes.NewReader(resp)),
			}, nil
		},
		)

	return testClient
}

type testResponse struct {
	End             int           `json:"end,omitempty"`
	FirstPageUri    string        `json:"first_page_uri,omitempty"`
	Messages        []testMessage `json:"messages,omitempty"`
	NextPageUri     string        `json:"next_page_uri,omitempty"`
	Page            int           `json:"page,omitempty"`
	PageSize        int           `json:"page_size,omitempty"`
	PreviousPageUri string        `json:"previous_page_uri,omitempty"`
	Start           int           `json:"start,omitempty"`
	Uri             string        `json:"uri,omitempty"`
}

type testMessage struct {
	// The message text
	Body *string `json:"body,omitempty"`
	// The direction of the message
	Direction *string `json:"direction,omitempty"`
	// The phone number that initiated the message
	From *string `json:"from,omitempty"`
	// The status of the message
	Status *string `json:"status,omitempty"`
	// The phone number that received the message
	To *string `json:"to,omitempty"`
}

func getSomething(nextPageUrl string) (interface{}, error) {
	return nextPageUrl, nil
}

func TestPageUtil_GetNext(t *testing.T) {
	testClient := getTestClient(t)
	baseUrl := "https://api.twilio.com"
	response, _ := testClient.SendRequest("get", "", nil, nil) //nolint:bodyclose
	ps := &testResponse{}
	_ = json.NewDecoder(response.Body).Decode(ps)

	nextPageUrl, err := GetNext(baseUrl, ps, getSomething)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=2&Page=1&PageToken=PASMXX", nextPageUrl)
	assert.Nil(t, err)

	nextPageUrl, err = GetNext(baseUrl, nil, getSomething)
	assert.Empty(t, nextPageUrl)
	assert.Nil(t, err)
}

func TestPageUtil_ToMap(t *testing.T) {
	testMap, err := toMap("invalid")
	assert.NotNil(t, err)
	assert.Nil(t, testMap)

	valid := testResponse{
		End:             0,
		FirstPageUri:    "first",
		Messages:        nil,
		NextPageUri:     "next",
		Page:            0,
		PageSize:        0,
		PreviousPageUri: "previous",
		Start:           0,
		Uri:             "uri",
	}
	testMap, err = toMap(valid)
	assert.Nil(t, err)
	assert.NotNil(t, testMap)
}

// Tests for Safe Extraction Helpers

func TestGetStringFromMap_ValidPath(t *testing.T) {
	data := map[string]interface{}{
		"meta": map[string]interface{}{
			"next_token": "TOKEN123",
		},
	}

	value, err := getStringFromMap(data, "meta", "next_token")
	assert.Nil(t, err)
	assert.Equal(t, "TOKEN123", value)
}

func TestGetStringFromMap_MissingKey(t *testing.T) {
	data := map[string]interface{}{
		"meta": map[string]interface{}{},
	}

	value, err := getStringFromMap(data, "meta", "next_token")
	assert.Nil(t, err)
	assert.Equal(t, "", value)
}

func TestGetStringFromMap_WrongType(t *testing.T) {
	data := map[string]interface{}{
		"meta": map[string]interface{}{
			"next_token": 123,
		},
	}

	value, err := getStringFromMap(data, "meta", "next_token")
	assert.NotNil(t, err)
	assert.Equal(t, "", value)
}

func TestGetStringFromMap_EmptyPath(t *testing.T) {
	data := map[string]interface{}{}

	value, err := getStringFromMap(data)
	assert.NotNil(t, err)
	assert.Equal(t, "", value)
}

func TestGetIntFromMap_ValidInt(t *testing.T) {
	data := map[string]interface{}{
		"meta": map[string]interface{}{
			"page_size": 50,
		},
	}

	value, err := getIntFromMap(data, "meta", "page_size")
	assert.Nil(t, err)
	assert.Equal(t, 50, value)
}

func TestGetIntFromMap_ValidFloat64(t *testing.T) {
	data := map[string]interface{}{
		"meta": map[string]interface{}{
			"page_size": 50.0,
		},
	}

	value, err := getIntFromMap(data, "meta", "page_size")
	assert.Nil(t, err)
	assert.Equal(t, 50, value)
}

func TestGetIntFromMap_MissingKey(t *testing.T) {
	data := map[string]interface{}{
		"meta": map[string]interface{}{},
	}

	value, err := getIntFromMap(data, "meta", "page_size")
	assert.Nil(t, err)
	assert.Equal(t, 0, value)
}

func TestGetIntFromMap_WrongType(t *testing.T) {
	data := map[string]interface{}{
		"meta": map[string]interface{}{
			"page_size": "fifty",
		},
	}

	value, err := getIntFromMap(data, "meta", "page_size")
	assert.NotNil(t, err)
	assert.Equal(t, 0, value)
}

// Tests for Token URL Building

func TestBuildTokenUrl_WithTokenAndPageSize(t *testing.T) {
	url := buildTokenUrl("https://api.twilio.com/v1/Services", "TOKEN123", 50)
	assert.Equal(t, "https://api.twilio.com/v1/Services?PageSize=50&PageToken=TOKEN123", url)
}

func TestBuildTokenUrl_WithTokenOnly(t *testing.T) {
	url := buildTokenUrl("https://api.twilio.com/v1/Services", "TOKEN123", 0)
	assert.Equal(t, "https://api.twilio.com/v1/Services?PageToken=TOKEN123", url)
}

func TestBuildTokenUrl_EmptyToken(t *testing.T) {
	url := buildTokenUrl("https://api.twilio.com/v1/Services", "", 50)
	assert.Equal(t, "", url)
}

func TestBuildTokenUrl_WithExistingQueryParams(t *testing.T) {
	url := buildTokenUrl("https://api.twilio.com/v1/Services?Status=active&Limit=100", "TOKEN123", 50)
	assert.Contains(t, url, "Status=active")
	assert.Contains(t, url, "Limit=100")
	assert.Contains(t, url, "PageSize=50")
	assert.Contains(t, url, "PageToken=TOKEN123")
	// Ensure proper delimiter is used (should have & not double ?)
	assert.NotContains(t, url, "??")
	assert.Contains(t, url, "?Status=active&Limit=100&PageSize=50&PageToken=TOKEN123")
}

func TestBuildTokenUrl_WithExistingQueryParamsTokenOnly(t *testing.T) {
	url := buildTokenUrl("https://api.twilio.com/v1/Services?Filter=test", "TOKEN123", 0)
	assert.Contains(t, url, "Filter=test")
	assert.Contains(t, url, "PageToken=TOKEN123")
	assert.NotContains(t, url, "??")
	assert.Contains(t, url, "?Filter=test&PageToken=TOKEN123")
}

// Tests for Token Metadata Extraction

func TestExtractTokenMeta_SnakeCase(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"next_token":     "TOKEN123",
			"previous_token": "TOKEN122",
			"page_size":      50,
			"key":            "items",
		},
	}

	meta, err := extractTokenMeta(payload)
	assert.Nil(t, err)
	assert.NotNil(t, meta)
	assert.Equal(t, "TOKEN123", meta.NextToken)
	assert.Equal(t, "TOKEN122", meta.PreviousToken)
	assert.Equal(t, 50, meta.PageSize)
	assert.Equal(t, "items", meta.Key)
}

func TestExtractTokenMeta_CamelCase(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"nextToken":     "TOKEN123",
			"previousToken": "TOKEN122",
			"pageSize":      25,
			"key":           "records",
		},
	}

	meta, err := extractTokenMeta(payload)
	assert.Nil(t, err)
	assert.NotNil(t, meta)
	assert.Equal(t, "TOKEN123", meta.NextToken)
	assert.Equal(t, "TOKEN122", meta.PreviousToken)
	assert.Equal(t, 25, meta.PageSize)
	assert.Equal(t, "records", meta.Key)
}

func TestExtractTokenMeta_NoMeta(t *testing.T) {
	payload := map[string]interface{}{
		"items": []interface{}{},
	}

	meta, err := extractTokenMeta(payload)
	assert.Nil(t, err)
	assert.Nil(t, meta)
}

func TestExtractTokenMeta_NoTokens(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"page_size": 50,
		},
	}

	meta, err := extractTokenMeta(payload)
	assert.Nil(t, err)
	assert.Nil(t, meta)
}

func TestExtractTokenMeta_NextTokenOnly(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"next_token": "TOKEN123",
			"page_size":  50,
		},
	}

	meta, err := extractTokenMeta(payload)
	assert.Nil(t, err)
	assert.NotNil(t, meta)
	assert.Equal(t, "TOKEN123", meta.NextToken)
	assert.Equal(t, "", meta.PreviousToken)
	assert.Equal(t, 50, meta.PageSize)
}

func TestExtractTokenMeta_PreviousTokenOnly(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"previous_token": "TOKEN122",
		},
	}

	meta, err := extractTokenMeta(payload)
	assert.Nil(t, err)
	assert.NotNil(t, meta)
	assert.Equal(t, "", meta.NextToken)
	assert.Equal(t, "TOKEN122", meta.PreviousToken)
}

// Tests for Token-Based Pagination Integration

func TestGetNextPageUrl_TokenBased(t *testing.T) {
	payload := map[string]interface{}{
		"items": []interface{}{},
		"meta": map[string]interface{}{
			"next_token": "TOKEN123",
			"page_size":  25,
		},
	}

	nextUrl, err := getNextPageUrl("https://api.twilio.com/v1/Services", payload)
	assert.Nil(t, err)
	assert.Contains(t, nextUrl, "PageToken=TOKEN123")
	assert.Contains(t, nextUrl, "PageSize=25")
}

func TestGetNextPageUrl_TokenBasedCamelCase(t *testing.T) {
	payload := map[string]interface{}{
		"records": []interface{}{},
		"meta": map[string]interface{}{
			"nextToken": "TOKEN456",
			"pageSize":  50,
		},
	}

	nextUrl, err := getNextPageUrl("https://api.twilio.com/v1/Services", payload)
	assert.Nil(t, err)
	assert.Contains(t, nextUrl, "PageToken=TOKEN456")
	assert.Contains(t, nextUrl, "PageSize=50")
}

func TestGetNextPageUrl_UrlBasedTakesPrecedence(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"next_page_url": "https://api.twilio.com/v1/Services?Page=2",
			"next_token":    "TOKEN123",
		},
	}

	nextUrl, err := getNextPageUrl("https://api.twilio.com/v1/Services", payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.twilio.com/v1/Services?Page=2", nextUrl)
	assert.NotContains(t, nextUrl, "TOKEN123")
}

func TestGetNextPageUrl_TokenBasedWithExistingQueryParams(t *testing.T) {
	payload := map[string]interface{}{
		"items": []interface{}{},
		"meta": map[string]interface{}{
			"next_token": "TOKEN123",
			"page_size":  25,
		},
	}

	// Base URL already has query parameters
	baseUrl := "https://api.twilio.com/v1/Services?Status=active&Type=sms"
	nextUrl, err := getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Contains(t, nextUrl, "Status=active")
	assert.Contains(t, nextUrl, "Type=sms")
	assert.Contains(t, nextUrl, "PageToken=TOKEN123")
	assert.Contains(t, nextUrl, "PageSize=25")
	// Verify proper concatenation with &
	assert.Contains(t, nextUrl, "?Status=active&Type=sms&PageSize=25&PageToken=TOKEN123")
}

func TestGetPreviousPageUrl_TokenBased(t *testing.T) {
	payload := map[string]interface{}{
		"items": []interface{}{},
		"meta": map[string]interface{}{
			"previous_token": "TOKEN122",
			"page_size":      25,
		},
	}

	prevUrl, err := getPreviousPageUrl("https://api.twilio.com/v1/Services", payload)
	assert.Nil(t, err)
	assert.Contains(t, prevUrl, "PageToken=TOKEN122")
	assert.Contains(t, prevUrl, "PageSize=25")
}

func TestGetPreviousPageUrl_NoPreviousToken(t *testing.T) {
	payload := map[string]interface{}{
		"items": []interface{}{},
		"meta": map[string]interface{}{
			"next_token": "TOKEN123",
			"page_size":  25,
		},
	}

	prevUrl, err := getPreviousPageUrl("https://api.twilio.com/v1/Services", payload)
	assert.Nil(t, err)
	assert.Equal(t, "", prevUrl)
}

// Tests for Backward Compatibility

func TestGetNextPageUrl_BackwardCompatibility_MetaNextPageUrl(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"next_page_url": "https://api.twilio.com/2010-04-01/Accounts/ACXX/Messages.json?PageSize=50&Page=1",
			"page_size":     50,
		},
	}

	nextPageUrl, err := getNextPageUrl("https://apitest.twilio.com", payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/Messages.json?PageSize=50&Page=1", nextPageUrl)
}

func TestGetNextPageUrl_BackwardCompatibility_NextPageUri(t *testing.T) {
	payload := map[string]interface{}{
		"next_page_uri": "/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1",
		"page_size":     50,
	}
	baseUrl := "https://api.twilio.com/"
	nextPageUrl, err := getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)
}

// Tests for KeyError

func TestKeyError_ErrorMessage(t *testing.T) {
	err := NewKeyError("next_token", "required for pagination")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "next_token")
	assert.Contains(t, err.Error(), "required for pagination")
	assert.Contains(t, err.Error(), "KeyError")
}

func TestKeyError_Fields(t *testing.T) {
	err := NewKeyError("test_key", "test message")
	assert.Equal(t, "test_key", err.Key)
	assert.Equal(t, "test message", err.Message)
}
