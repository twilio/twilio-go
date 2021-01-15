/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"fmt"
    twilio "github.com/twilio/twilio-go/client"
    "net/url"
    "strings"
    ""
)

type SummaryTagApiService struct {
    baseURL string
    client  *twilio.Client
}

func NewSummaryTagApiService(client *twilio.Client) *SummaryTagApiService {
    return &SummaryTagApiService{
        client: client,
        baseURL: fmt.Sprintf("https://studio.%s", client.BaseURL),
    }
}
// FetchSummaryParams Optional parameters for the method 'FetchSummary'
type FetchSummaryParams struct {
    ProcessingState *string `json:"ProcessingState,omitempty"`
}

/*
FetchSummary Method for FetchSummary
 * @param callSid
 * @param optional nil or *FetchSummaryOpts - Optional Parameters:
 * @param "ProcessingState" (string) - 
@return InsightsV1CallSummary
*/
func (c *SummaryTagApiService) FetchSummary(callSid string, params *FetchSummaryParams) (*InsightsV1CallSummary, error) {
    path := "/v1/Voice/{CallSid}/Summary"
    path = strings.Replace(path, "{"+"CallSid"+"}", callSid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.ProcessingState != nil {
        data.Set("ProcessingState", *params.ProcessingState)
    }


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &InsightsV1CallSummary{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
