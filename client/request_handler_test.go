package client_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	assert.Equal(t, 5, requestHandler.ReadLimits(setPageSize(10), 5))
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
						"direction": "inbound",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "accepted",
					},
					{
						"direction": "outbound-reply",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "read",
					},
					{
						"direction": "outbound-call",
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

	resp, err := twilio.AccountsMessagesList(&params, 5)
	assert.Equal(t, "outbound-api", *resp[0].Messages[0].Direction)
	assert.Equal(t, "4444444444", *resp[0].Messages[0].From)
	assert.Equal(t, "inbound", *resp[0].Messages[1].Direction)
	assert.Equal(t, "read", *resp[0].Messages[2].Status)
	assert.Len(t, resp, 1)
	assert.Equal(t, 5, getRecordCount(resp))
	assert.Nil(t, err)

	resp, _ = twilio.AccountsMessagesList(&params, 5)
	assert.Len(t, resp, 1)
	assert.Equal(t, 5, getRecordCount(resp))

	resp, _ = twilio.AccountsMessagesList(&params, 10)
	assert.Len(t, resp, 2)
	assert.Equal(t, 10, getRecordCount(resp))
}

func getRecordCount(resp []apiV2010.ListMessageResponse) int {
	messagesCount := 0
	for page := range resp {
		messagesCount += len(resp[page].Messages)
	}
	return messagesCount
}

func TestRequestHandler_ListPaging(t *testing.T) {
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
						"direction": "outbound-call",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "accepted",
					},
					{
						"direction": "outbound-reply",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "canceled",
					},
					{
						"direction": "inbound",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "received",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "queued",
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
						"direction": "outbound-call",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi!",
						"status":    "queued",
					},
					{
						"direction": "outbound-reply",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "sent",
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

	resp, err := twilio.AccountsMessagesList(&params, 10)
	assert.Equal(t, "delivered", *resp[0].Messages[0].Status)
	assert.Equal(t, "Hi!", *resp[1].Messages[1].Body)
	assert.Len(t, resp, 3)
	assert.Equal(t, 9, getRecordCount(resp))
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

	resp, err := twilio.AccountsMessagesList(&params, 5)
	assert.Len(t, resp, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "Listing error", err.Error())
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
						"direction": "outbound-call",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Hi",
						"status":    "queued",
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
						"body":      "Hello",
						"status":    "sent",
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
	params.SetPageSize(5)

	pageCount := 0

	pageCount, messageCount := getStreamCount(twilio, params, 10)
	assert.Equal(t, 2, pageCount)
	assert.Equal(t, 10, messageCount)

	pageCount, messageCount = getStreamCount(twilio, params, 15)
	assert.Equal(t, 3, pageCount)
	assert.Equal(t, 15, messageCount)

	pageCount, messageCount = getStreamCount(twilio, params, 40)
	assert.Equal(t, 8, pageCount)
	assert.Equal(t, 40, messageCount)
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
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 2",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 3",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 4",
						"status":    "delivered",
					},
				},
				"uri":           "/2010-04-01/Accounts/AC12345678123456781234567812345678/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=5&Page=0&PageToken=dummy",
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
						"body":      "Message 5",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 6",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 7",
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

	gomock.InOrder(
		page0,
		page1,
	)

	twilio := apiV2010.NewApiServiceWithClient(testClient)

	params := apiV2010.ListMessageParams{}
	params.SetFrom("4444444444")
	params.SetTo("9999999999")
	params.SetPageSize(5)

	pageCount := 0
	messageCount := 0

	for {
		channel, _ := twilio.AccountsMessagesStream(&params, 8)
		for record := range channel {
			pageCount += 1
			messages := record.(*apiV2010.ListMessageResponse).Messages
			for message := range messages {
				text := fmt.Sprintf("Message %d", messageCount)
				assert.Equal(t, text, *messages[message].Body)
				messageCount += 1
			}
		}

		if len(channel) == 0 {
			break
		}
	}

	assert.Equal(t, 2, pageCount)
	assert.Equal(t, 8, messageCount)
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
	params.SetPageSize(5)

	resp, err := twilio.AccountsMessagesStream(&params, 5)
	assert.Len(t, resp, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "Streaming error", err.Error())
}

func getStreamCount(twilio *apiV2010.ApiService, params apiV2010.ListMessageParams, limit int) (int, int) {
	count := 0
	messageCount := 0

	for {
		channel, _ := twilio.AccountsMessagesStream(&params, limit)
		for record := range channel {
			count += 1
			messages := record.(*apiV2010.ListMessageResponse).Messages
			for range messages {
				messageCount += 1
			}
		}

		if len(channel) == 0 {
			break
		}
	}
	return count, messageCount
}
