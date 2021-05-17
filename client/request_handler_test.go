package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"
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
	assert.Equal(t, "prism_twilio:4010", requestHandler.BuildUrl("prism_twilio:4010"))
}
