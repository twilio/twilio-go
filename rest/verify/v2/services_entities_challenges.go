/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateChallenge'
type CreateChallengeParams struct {
	// Optional payload used to verify the Challenge upon creation. Only used with a Factor of type `totp` to carry the TOTP code that needs to be verified. For `TOTP` this value must be between 3 and 8 characters long.
	AuthPayload *string `json:"AuthPayload,omitempty"`
	// A list of objects that describe the Fields included in the Challenge. Each object contains the label and value of the field, the label can be up to 36 characters in length and the value can be up to 128 characters in length. Used when `factor_type` is `push`. There can be up to 20 details fields.
	DetailsFields *[]map[string]interface{} `json:"Details.Fields,omitempty"`
	// Shown to the user when the push notification arrives. Required when `factor_type` is `push`. Can be up to 256 characters in length
	DetailsMessage *string `json:"Details.Message,omitempty"`
	// The date-time when this Challenge expires, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. The default value is five (5) minutes after Challenge creation. The max value is sixty (60) minutes after creation.
	ExpirationDate *time.Time `json:"ExpirationDate,omitempty"`
	// The unique SID identifier of the Factor.
	FactorSid *string `json:"FactorSid,omitempty"`
	// Details provided to give context about the Challenge. Not shown to the end user. It must be a stringified JSON with only strings values eg. `{\\\"ip\\\": \\\"172.168.1.234\\\"}`. Can be up to 1024 characters in length
	HiddenDetails *map[string]interface{} `json:"HiddenDetails,omitempty"`
}

func (params *CreateChallengeParams) SetAuthPayload(AuthPayload string) *CreateChallengeParams {
	params.AuthPayload = &AuthPayload
	return params
}
func (params *CreateChallengeParams) SetDetailsFields(DetailsFields []map[string]interface{}) *CreateChallengeParams {
	params.DetailsFields = &DetailsFields
	return params
}
func (params *CreateChallengeParams) SetDetailsMessage(DetailsMessage string) *CreateChallengeParams {
	params.DetailsMessage = &DetailsMessage
	return params
}
func (params *CreateChallengeParams) SetExpirationDate(ExpirationDate time.Time) *CreateChallengeParams {
	params.ExpirationDate = &ExpirationDate
	return params
}
func (params *CreateChallengeParams) SetFactorSid(FactorSid string) *CreateChallengeParams {
	params.FactorSid = &FactorSid
	return params
}
func (params *CreateChallengeParams) SetHiddenDetails(HiddenDetails map[string]interface{}) *CreateChallengeParams {
	params.HiddenDetails = &HiddenDetails
	return params
}

// Create a new Challenge for the Factor
func (c *ApiService) CreateChallenge(ServiceSid string, Identity string, params *CreateChallengeParams) (*VerifyV2Challenge, error) {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}/Challenges"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	if params != nil && params.AuthPayload != nil {
		data.Set("AuthPayload", *params.AuthPayload)
	}
	if params != nil && params.DetailsFields != nil {
		for _, item := range *params.DetailsFields {
			v, err := json.Marshal(item)

			if err != nil {
				return nil, err
			}

			data.Add("Details.Fields", string(v))
		}
	}
	if params != nil && params.DetailsMessage != nil {
		data.Set("Details.Message", *params.DetailsMessage)
	}
	if params != nil && params.ExpirationDate != nil {
		data.Set("ExpirationDate", fmt.Sprint((*params.ExpirationDate).Format(time.RFC3339)))
	}
	if params != nil && params.FactorSid != nil {
		data.Set("FactorSid", *params.FactorSid)
	}
	if params != nil && params.HiddenDetails != nil {
		v, err := json.Marshal(params.HiddenDetails)

		if err != nil {
			return nil, err
		}

		data.Set("HiddenDetails", string(v))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Challenge{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch a specific Challenge.
func (c *ApiService) FetchChallenge(ServiceSid string, Identity string, Sid string) (*VerifyV2Challenge, error) {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}/Challenges/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Challenge{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListChallenge'
type ListChallengeParams struct {
	// The unique SID identifier of the Factor.
	FactorSid *string `json:"FactorSid,omitempty"`
	// The Status of the Challenges to fetch. One of `pending`, `expired`, `approved` or `denied`.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListChallengeParams) SetFactorSid(FactorSid string) *ListChallengeParams {
	params.FactorSid = &FactorSid
	return params
}
func (params *ListChallengeParams) SetStatus(Status string) *ListChallengeParams {
	params.Status = &Status
	return params
}
func (params *ListChallengeParams) SetPageSize(PageSize int) *ListChallengeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListChallengeParams) SetLimit(Limit int) *ListChallengeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Challenge records from the API. Request is executed immediately.
func (c *ApiService) PageChallenge(ServiceSid string, Identity string, params *ListChallengeParams, pageToken string, pageNumber string) (*ListChallengeResponse, error) {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}/Challenges"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	if params != nil && params.FactorSid != nil {
		data.Set("FactorSid", *params.FactorSid)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChallengeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Challenge records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListChallenge(ServiceSid string, Identity string, params *ListChallengeParams) ([]VerifyV2Challenge, error) {
	if params == nil {
		params = &ListChallengeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageChallenge(ServiceSid, Identity, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VerifyV2Challenge

	for response != nil {
		records = append(records, response.Challenges...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListChallengeResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListChallengeResponse)
	}

	return records, err
}

// Streams Challenge records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamChallenge(ServiceSid string, Identity string, params *ListChallengeParams) (chan VerifyV2Challenge, error) {
	if params == nil {
		params = &ListChallengeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageChallenge(ServiceSid, Identity, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VerifyV2Challenge, 1)

	go func() {
		for response != nil {
			for item := range response.Challenges {
				channel <- response.Challenges[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListChallengeResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListChallengeResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListChallengeResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChallengeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateChallenge'
type UpdateChallengeParams struct {
	// The optional payload needed to verify the Challenge. E.g., a TOTP would use the numeric code. For `TOTP` this value must be between 3 and 8 characters long. For `Push` this value can be up to 5456 characters in length
	AuthPayload *string `json:"AuthPayload,omitempty"`
}

func (params *UpdateChallengeParams) SetAuthPayload(AuthPayload string) *UpdateChallengeParams {
	params.AuthPayload = &AuthPayload
	return params
}

// Verify a specific Challenge.
func (c *ApiService) UpdateChallenge(ServiceSid string, Identity string, Sid string, params *UpdateChallengeParams) (*VerifyV2Challenge, error) {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}/Challenges/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	if params != nil && params.AuthPayload != nil {
		data.Set("AuthPayload", *params.AuthPayload)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Challenge{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
