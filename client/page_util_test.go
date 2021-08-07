package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLimits(t *testing.T) {
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

func TestGetNextPageUri(t *testing.T) {
	payload := map[string]interface{}{
		"next_page_uri": "/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1",
		"page_size":     50,
	}
	baseUrl := "https://api.twilio.com/"
	nextPageUrl := getNextPageUrl(baseUrl, payload)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)

	payload["next_page_uri"] = "2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1"
	baseUrl = "https://api.twilio.com"
	nextPageUrl = getNextPageUrl(baseUrl, payload)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)

	payload = map[string]interface{}{}
	nextPageUrl = getNextPageUrl(baseUrl, payload)
	assert.Equal(t, "", nextPageUrl)
}

func TestGetNextPageUrl(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"next_page_url": "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1",
			"page_size":     50,
		},
	}

	nextPageUrl := getNextPageUrl("https://apitest.twilio.com", payload)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)
}
