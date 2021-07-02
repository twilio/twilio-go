package client_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/pkg/errors"

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
	assert.Equal(t, 5, requestHandler.ReadLimits(nil, 5))
	assert.Equal(t, 10, requestHandler.ReadLimits(setPageSize(10), 5))
	assert.Equal(t, 1000, requestHandler.ReadLimits(nil, 5000))
	assert.Equal(t, 10, requestHandler.ReadLimits(setPageSize(10), 0))
	assert.Equal(t, 50, requestHandler.ReadLimits(nil, 0))

}

func setPageSize(pageSize int) *int {
	PageSize := &pageSize
	return PageSize
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
	params.SetFrom("4444444444")
	params.SetTo("9999999999")
	params.SetPageSize(5)

	resp, err := twilio.MessageList(&params, 5)
	assert.Equal(t, "outbound-api", resp[0].(map[string]interface{})["direction"])
	assert.Equal(t, "4444444444", resp[0].(map[string]interface{})["from"])
	assert.Len(t, resp, 5)
	assert.Nil(t, err)

	params.SetPageSize(2)
	resp, _ = twilio.MessageList(&params, 5)
	assert.Len(t, resp, 5)

	resp, _ = twilio.MessageList(&params, 2)
	assert.Len(t, resp, 2)

	params.SetPageSize(5)
	resp, err = twilio.MessageList(&params, 2)
	assert.Len(t, resp, 2)
	assert.Nil(t, err)
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
	params.SetPageSize(5)

	resp, err := twilio.MessageList(&params, 10)
	assert.Len(t, resp, 9)
	assert.Nil(t, err)
}

func TestRequestHandler_ListError(t *testing.T) {
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
			}, errors.New("Listing error")
		},
		).AnyTimes()

	twilio := apiV2010.NewApiServiceWithClient(testClient)

	params := apiV2010.ListMessageParams{}
	params.SetFrom("from")
	params.SetTo("to")
	params.SetPageSize(5)

	resp, err := twilio.MessageList(&params, 5)
	assert.Len(t, resp, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "Listing error", err.Error())
}

func TestRequestHandler_StreamError(t *testing.T) {
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
			}, errors.New("Streaming error")
		},
		).AnyTimes()

	twilio := apiV2010.NewApiServiceWithClient(testClient)

	params := apiV2010.ListMessageParams{}
	params.SetFrom("from")
	params.SetTo("to")
	params.SetPageSize(5)

	resp, err := twilio.MessageStream(&params, 5)
	assert.Len(t, resp, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "Streaming error", err.Error())
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
	params.SetPageSize(2)

	count := getStreamCount(twilio, params, 0)
	assert.Equal(t, 10, count)

	count = getStreamCount(twilio, params, 0)

	assert.Equal(t, 10, count)
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
	params.SetPageSize(5)

	count := getStreamCount(twilio, params, 0)
	assert.Equal(t, 8, count)
}

func getStreamCount(twilio *apiV2010.ApiService, params apiV2010.ListMessageParams, count int) int {
	for {
		channel, _ := twilio.MessageStream(&params, 10)
		for range channel {
			count += 1
		}

		if len(channel) == 0 {
			break
		}
	}
	return count
}
