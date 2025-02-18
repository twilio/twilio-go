/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateIpCommand'
type CreateIpCommandParams struct {
	// The `sid` or `unique_name` of the [Super SIM](https://www.twilio.com/docs/iot/supersim/api/sim-resource) to send the IP Command to.
	Sim *string `json:"Sim,omitempty"`
	// The data that will be sent to the device. The payload cannot exceed 1300 bytes. If the PayloadType is set to text, the payload is encoded in UTF-8. If PayloadType is set to binary, the payload is encoded in Base64.
	Payload *string `json:"Payload,omitempty"`
	// The device port to which the IP Command will be sent.
	DevicePort *int `json:"DevicePort,omitempty"`
	//
	PayloadType *string `json:"PayloadType,omitempty"`
	// The URL we should call using the `callback_method` after we have sent the IP Command.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be `GET` or `POST`, and the default is `POST`.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
}

func (params *CreateIpCommandParams) SetSim(Sim string) *CreateIpCommandParams {
	params.Sim = &Sim
	return params
}
func (params *CreateIpCommandParams) SetPayload(Payload string) *CreateIpCommandParams {
	params.Payload = &Payload
	return params
}
func (params *CreateIpCommandParams) SetDevicePort(DevicePort int) *CreateIpCommandParams {
	params.DevicePort = &DevicePort
	return params
}
func (params *CreateIpCommandParams) SetPayloadType(PayloadType string) *CreateIpCommandParams {
	params.PayloadType = &PayloadType
	return params
}
func (params *CreateIpCommandParams) SetCallbackUrl(CallbackUrl string) *CreateIpCommandParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateIpCommandParams) SetCallbackMethod(CallbackMethod string) *CreateIpCommandParams {
	params.CallbackMethod = &CallbackMethod
	return params
}

// Send an IP Command to a Super SIM.
func (c *ApiService) CreateIpCommand(params *CreateIpCommandParams) (*SupersimV1IpCommand, error) {
	path := "/v1/IpCommands"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Payload != nil {
		data.Set("Payload", *params.Payload)
	}
	if params != nil && params.DevicePort != nil {
		data.Set("DevicePort", fmt.Sprint(*params.DevicePort))
	}
	if params != nil && params.PayloadType != nil {
		data.Set("PayloadType", *params.PayloadType)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1IpCommand{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch IP Command instance from your account.
func (c *ApiService) FetchIpCommand(Sid string) (*SupersimV1IpCommand, error) {
	path := "/v1/IpCommands/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1IpCommand{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIpCommand'
type ListIpCommandParams struct {
	// The SID or unique name of the Sim resource that IP Command was sent to or from.
	Sim *string `json:"Sim,omitempty"`
	// The ICCID of the Sim resource that IP Command was sent to or from.
	SimIccid *string `json:"SimIccid,omitempty"`
	// The status of the IP Command. Can be: `queued`, `sent`, `received` or `failed`. See the [IP Command Status Values](https://www.twilio.com/docs/iot/supersim/api/ipcommand-resource#status-values) for a description of each.
	Status *string `json:"Status,omitempty"`
	// The direction of the IP Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`.
	Direction *string `json:"Direction,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int64 `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int64 `json:"limit,omitempty"`
}

func (params *ListIpCommandParams) SetSim(Sim string) *ListIpCommandParams {
	params.Sim = &Sim
	return params
}
func (params *ListIpCommandParams) SetSimIccid(SimIccid string) *ListIpCommandParams {
	params.SimIccid = &SimIccid
	return params
}
func (params *ListIpCommandParams) SetStatus(Status string) *ListIpCommandParams {
	params.Status = &Status
	return params
}
func (params *ListIpCommandParams) SetDirection(Direction string) *ListIpCommandParams {
	params.Direction = &Direction
	return params
}
func (params *ListIpCommandParams) SetPageSize(PageSize int64) *ListIpCommandParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListIpCommandParams) SetLimit(Limit int64) *ListIpCommandParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of IpCommand records from the API. Request is executed immediately.
func (c *ApiService) PageIpCommand(params *ListIpCommandParams, pageToken, pageNumber string) (*ListIpCommandResponse, error) {
	path := "/v1/IpCommands"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.SimIccid != nil {
		data.Set("SimIccid", *params.SimIccid)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIpCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists IpCommand records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIpCommand(params *ListIpCommandParams) ([]SupersimV1IpCommand, error) {
	response, errors := c.StreamIpCommand(params)

	records := make([]SupersimV1IpCommand, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams IpCommand records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIpCommand(params *ListIpCommandParams) (chan SupersimV1IpCommand, chan error) {
	if params == nil {
		params = &ListIpCommandParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1IpCommand, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageIpCommand(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamIpCommand(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamIpCommand(response *ListIpCommandResponse, params *ListIpCommandParams, recordChannel chan SupersimV1IpCommand, errorChannel chan error) {
	var curRecord int64 = 1

	for response != nil {
		responseRecords := response.IpCommands
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListIpCommandResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListIpCommandResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListIpCommandResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIpCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
