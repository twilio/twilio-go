package client_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"

	apiV2010 "github.com/twilio/twilio-go/rest/api/v2010"
)

func NewRequestHandler(accountSid string, authToken string) *client.RequestHandler {
	c := NewClient(accountSid, authToken)
	return client.NewRequestHandler(c)
}

func TestRequestHandler_BuildUrlSetRegion(t *testing.T) {
	// Region set via url
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://api.region.twilio.com", requestHandler.BuildUrl("https://api.region.twilio.com"))

	// Region set via requestHandler
	requestHandler.Region = "region"
	assert.Equal(t, "https://api.region.twilio.com", requestHandler.BuildUrl("https://api.twilio.com"))
	assert.Equal(t, "https://api.region.twilio.com", requestHandler.BuildUrl("https://api.urlRegion.twilio.com"))
}

func TestRequestHandler_BuildUrlSetEdgeDefaultRegion(t *testing.T) {
	// Edge set via client
	requestHandler := NewRequestHandler("user", "pass")
	requestHandler.Edge = "edge"
	assert.Equal(t, "https://api.edge.us1.twilio.com", requestHandler.BuildUrl("https://api.twilio.com"))
}

func TestRequestHandler_BuildUrlSetEdgeRegion(t *testing.T) {
	//Edge and Region set via url
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://api.edge.region.twilio.com", requestHandler.BuildUrl("https://api.edge.region.twilio.com"))

	// Edge and Region set via client
	requestHandler.Edge = "edge"
	assert.Equal(t, "https://api.edge.region.twilio.com", requestHandler.BuildUrl("https://api.region.twilio.com"))
	requestHandler.Region = "region"
	assert.Equal(t, "https://api.edge.region.twilio.com", requestHandler.BuildUrl("https://api.twilio.com"))
	assert.Equal(t, "https://api.edge.region.twilio.com", requestHandler.BuildUrl("https://api.urlEdge.urlRegion.twilio.com"))
}

func TestRequestHandler_BuildHostRawHostWithoutPeriods(t *testing.T) {
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://prism_twilio:4010", requestHandler.BuildUrl("https://prism_twilio:4010"))
}

func TestRequestHandler_ReadLimits(t *testing.T) {
	requestHandler := NewRequestHandler("user", "pass")
	metaData := client.PaginationData{
		PageSize: 0,
		Limit:    5,
	}
	assert.Equal(t, 5, requestHandler.ReadLimits(metaData).PageSize)
	assert.Equal(t, 5, requestHandler.ReadLimits(metaData).Limit)

	metaData.PageSize = 10
	assert.Equal(t, 10, requestHandler.ReadLimits(metaData).PageSize)
	assert.Equal(t, 5, requestHandler.ReadLimits(metaData).Limit)

	metaData = client.PaginationData{
		Limit: 5000,
	}
	assert.Equal(t, 1000, requestHandler.ReadLimits(metaData).PageSize)
}

func TestRequestHandler_List(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	testClient := client.NewMockBaseClient(mockCtrl)

	testClient.EXPECT().AccountSid().DoAndReturn(func() string {
		return "AC222222222222222222222222222222"
	}).AnyTimes()

	testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}) (*http.Response, error) {
			response := map[string]interface{}{
				"end":            4,
				"first_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0",
				"messages": []map[string]interface{}{
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
				},
				"uri":           "“/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0&PageToken=",
				"page_size":     5,
				"start":         0,
				"next_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=1&PageToken=PASMc49f620580b24424bcfa885b1f741130",
				"page":          0,
			}

			resp, _ := json.Marshal(response)

			return &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader(resp)),
			}, nil
		},
		).AnyTimes()

	twilio := apiV2010.NewApiServiceWithClient(testClient)

	params := apiV2010.ListMessageParams{}
	params.SetFrom("from")
	params.SetTo("to")

	data := client.PaginationData{
		PageSize: 5,
		Limit:    5,
	}

	resp := twilio.MessageList(&params, data)
	assert.Len(t, resp, 5)

	data.PageSize = 2
	resp = twilio.MessageList(&params, data)
	assert.Len(t, resp, 5)

	data.Limit = 2
	resp = twilio.MessageList(&params, data)
	assert.Len(t, resp, 2)

	data.PageSize = 5
	resp = twilio.MessageList(&params, data)
	assert.Len(t, resp, 2)

	data.PageLimit = 1
	data.PageSize = 5
	data.Limit = 10
	resp = twilio.MessageList(&params, data)
	assert.Len(t, resp, 5)
}

func TestRequestHandler_List_Paging(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	testClient := client.NewMockBaseClient(mockCtrl)
	testClient.EXPECT().AccountSid().DoAndReturn(func() string {
		return "AC222222222222222222222222222222"
	}).AnyTimes()

	page0 := testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}) (*http.Response, error) {
			response := map[string]interface{}{
				"end":            4,
				"first_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0",
				"messages": []map[string]interface{}{
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
				},
				"uri":           "“/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0&PageToken=",
				"page_size":     5,
				"start":         0,
				"next_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=1&PageToken=PASMc49f620580b24424bcfa885b1f741130",
				"page":          0,
			}

			resp, _ := json.Marshal(response)

			return &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader(resp)),
			}, nil
		},
		)

	page1 := testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}) (*http.Response, error) {
			response := map[string]interface{}{
				"end":            3,
				"first_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0",
				"messages": []map[string]interface{}{
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
				},
				"uri":           "“/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=1&PageToken=",
				"page_size":     2,
				"start":         0,
				"next_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=2&PageToken=PASMc49f620580b24424bcfa885b1f741130",
				"page":          1,
			}

			resp, _ := json.Marshal(response)

			return &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader(resp)),
			}, nil
		},
		)

	page2 := testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}) (*http.Response, error) {
			return &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader(nil)),
			}, nil
		},
		)

	gomock.InOrder(
		page0,
		page1,
		page2,
	)

	twilio := apiV2010.NewApiServiceWithClient(testClient)

	params := apiV2010.ListMessageParams{}
	params.SetFrom("from")
	params.SetTo("to")

	data := client.PaginationData{
		PageSize: 5,
		Limit:    10,
	}

	resp := twilio.MessageList(&params, data)
	assert.Len(t, resp, 9)
}

func TestRequestHandler_Stream(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	testClient := client.NewMockBaseClient(mockCtrl)
	testClient.EXPECT().AccountSid().DoAndReturn(func() string {
		return "AC222222222222222222222222222222"
	}).AnyTimes()

	testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}) (*http.Response, error) {
			response := map[string]interface{}{
				"end":            4,
				"first_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0",
				"messages": []map[string]interface{}{
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
				},
				"uri":           "“/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0&PageToken=",
				"page_size":     5,
				"start":         0,
				"next_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=1&PageToken=PASMc49f620580b24424bcfa885b1f741130",
				"page":          0,
			}

			resp, _ := json.Marshal(response)

			return &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader(resp)),
			}, nil
		},
		).AnyTimes()

	twilio := apiV2010.NewApiServiceWithClient(testClient)

	params := apiV2010.ListMessageParams{}
	params.SetFrom("from")
	params.SetTo("to")

	data := client.PaginationData{
		PageSize: 2,
		Limit:    5,
	}

	count := 0
	for range twilio.MessageStream(&params, data) {
		count += 1
	}
	assert.Equal(t, 5, count)

	data.PageLimit = 1
	data.Limit = 10
	count = 0
	for range twilio.MessageStream(&params, data) {
		count += 1
	}
	assert.Equal(t, 5, count)
}

func TestRequestHandler_StreamPaging(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	testClient := client.NewMockBaseClient(mockCtrl)
	testClient.EXPECT().AccountSid().DoAndReturn(func() string {
		return "AC222222222222222222222222222222"
	}).AnyTimes()

	page0 := testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}) (*http.Response, error) {
			response := map[string]interface{}{
				"end":            4,
				"first_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0",
				"messages": []map[string]interface{}{
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
				},
				"uri":           "“/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0&PageToken=",
				"page_size":     5,
				"start":         0,
				"next_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=1&PageToken=PASMc49f620580b24424bcfa885b1f741130",
				"page":          0,
			}

			resp, _ := json.Marshal(response)

			return &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader(resp)),
			}, nil
		},
		)

	page1 := testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}) (*http.Response, error) {
			response := map[string]interface{}{
				"end":            2,
				"first_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0",
				"messages": []map[string]interface{}{
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "delivered",
					},
				},
				"uri":           "“/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=1&PageToken=",
				"page_size":     2,
				"start":         0,
				"next_page_uri": "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=2&PageToken=PASMc49f620580b24424bcfa885b1f741130",
				"page":          1,
			}

			resp, _ := json.Marshal(response)

			return &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader(resp)),
			}, nil
		},
		)

	page2 := testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}) (*http.Response, error) {
			return &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader(nil)),
			}, nil
		},
		)

	gomock.InOrder(
		page0,
		page1,
		page2,
	)

	twilio := apiV2010.NewApiServiceWithClient(testClient)

	params := apiV2010.ListMessageParams{}
	params.SetFrom("4444444444")
	params.SetTo("9999999999")

	data := client.PaginationData{
		PageSize: 5,
		Limit:    10,
	}

	count := 0
	for range twilio.MessageStream(&params, data) {
		count += 1
	}
	assert.Equal(t, 8, count)
}
