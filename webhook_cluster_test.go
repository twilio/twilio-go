//go:build webhook_cluster
// +build webhook_cluster

package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/localtunnel/go-localtunnel"
	"github.com/stretchr/testify/assert"
	twilio "github.com/twilio/twilio-go/client"
	StudioV2 "github.com/twilio/twilio-go/rest/studio/v2"
)

var testClient *RestClient
var authToken string

func TestMain(m *testing.M) {
	var apiKey = os.Getenv("TWILIO_API_KEY")
	var secret = os.Getenv("TWILIO_API_SECRET")
	var accountSid = os.Getenv("TWILIO_ACCOUNT_SID")
	authToken = os.Getenv("TWILIO_AUTH_TOKEN")
	testClient = NewRestClientWithParams(ClientParams{apiKey, secret, accountSid, nil})
	ret := m.Run()
	os.Exit(ret)
}

func createValidationServer(channel chan bool) *http.Server {
	server := &http.Server{}
	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.Header["X-Forwarded-Proto"][0] + "://" + r.Header["X-Forwarded-Host"][0] + r.URL.RequestURI()
		signatureHeader := r.Header["X-Twilio-Signature"]
		r.ParseForm()
		params := make(map[string]string)
		for k, v := range r.PostForm {
			params[k] = v[0]
		}
		requestValidator := twilio.NewRequestValidator(authToken)
		if len(signatureHeader) != 0 {
			channel <- requestValidator.Validate(url, params, r.Header["X-Twilio-Signature"][0])
		} else {
			channel <- false
		}
	})
	return server
}

func createStudioFlowParams(url string, method string) *StudioV2.CreateFlowParams {
	jsonStr := fmt.Sprintf(`{
		"description": "Studio Flow",
		"states": [
		  {
			"name": "Trigger",
			"type": "trigger",
			"transitions": [
			  {
				"next": "httpRequest",
				"event": "incomingRequest"
			  }
			],
			"properties": {
			}
		  },
		  {
			"name": "httpRequest",
			"type": "make-http-request",
			"transitions": [],
			"properties": {
			  "method": "%s",
			  "content_type": "application/x-www-form-urlencoded;charset=utf-8",
			  "url": "%s"
			}
		  }
		],
		"initial_state": "Trigger",
		"flags": {
		  "allow_concurrent_calls": true
		}
	  }`, method, url)

	var definition interface{}
	_ = json.Unmarshal([]byte(jsonStr), &definition)

	params := &StudioV2.CreateFlowParams{
		Definition: &definition,
	}
	params.SetFriendlyName("Go Cluster Test Flow")
	params.SetStatus("published")
	return params
}

func createStudioExecutionParams() *StudioV2.CreateExecutionParams {
	executionParams := &StudioV2.CreateExecutionParams{}
	executionParams.SetTo("To")
	executionParams.SetFrom("From")
	return executionParams
}

func executeFlow(t *testing.T, flowSid string) {
	_, exeErr := testClient.StudioV2.CreateExecution(flowSid, createStudioExecutionParams())
	if exeErr != nil {
		t.Fatal("Error with Studio Execution Creation: ", exeErr)
	}
}

func requestValidation(t *testing.T, method string) {
	//Invoke Localtunnel
	listener, ltErr := localtunnel.Listen(localtunnel.Options{})
	if ltErr != nil {
		t.Fatal("Error with Localtunnel: ", ltErr)
	}
	//Create Validation Server & Listen
	channel := make(chan bool)
	server := createValidationServer(channel)
	go server.Serve(listener)

	//Extra time for server to set up
	time.Sleep(1 * time.Second)

	//Create Studio Flow
	params := createStudioFlowParams(listener.URL(), method)
	resp, flowErr := testClient.StudioV2.CreateFlow(params)
	if flowErr != nil {
		t.Fatal("Error with Studio Flow Creation: ", flowErr)
	}
	flowSid := *resp.Sid
	executeFlow(t, flowSid)

	//Await for Request Validation
	afterCh := time.After(5 * time.Second)
	select {
	case validate := <-channel:
		assert.True(t, validate)
	case <-afterCh:
		t.Fatal("No request was sent to validation server")
	}
	defer testClient.StudioV2.DeleteFlow(flowSid)
	defer server.Close()
}

func TestRequestValidation_GETMethod(t *testing.T) {
	requestValidation(t, "GET")
}

func TestRequestValidation_POSTMethod(t *testing.T) {
	requestValidation(t, "POST")
}
