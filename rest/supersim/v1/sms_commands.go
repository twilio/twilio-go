/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'CreateSmsCommand'
type CreateSmsCommandParams struct {
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_method` after we have sent the command.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The message body of the SMS Command.
	Payload *string `json:"Payload,omitempty"`
	// The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the SMS Command to.
	Sim *string `json:"Sim,omitempty"`
}

func (params *CreateSmsCommandParams) SetCallbackMethod(CallbackMethod string) *CreateSmsCommandParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *CreateSmsCommandParams) SetCallbackUrl(CallbackUrl string) *CreateSmsCommandParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateSmsCommandParams) SetPayload(Payload string) *CreateSmsCommandParams {
	params.Payload = &Payload
	return params
}
func (params *CreateSmsCommandParams) SetSim(Sim string) *CreateSmsCommandParams {
	params.Sim = &Sim
	return params
}

// Send SMS Command to a Sim.
func (c *ApiService) CreateSmsCommand(params *CreateSmsCommandParams) (*SupersimV1SmsCommand, error) {
	path := "/v1/SmsCommands"

	data := url.Values{}
	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Payload != nil {
		data.Set("Payload", *params.Payload)
	}
	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1SmsCommand{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch SMS Command instance from your account.
func (c *ApiService) FetchSmsCommand(Sid string) (*SupersimV1SmsCommand, error) {
	path := "/v1/SmsCommands/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1SmsCommand{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSmsCommand'
type ListSmsCommandParams struct {
	// The SID or unique name of the Sim resource that SMS Command was sent to or from.
	Sim *string `json:"Sim,omitempty"`
	// The status of the SMS Command. Can be: `queued`, `sent`, `delivered`, `received` or `failed`. See the [SMS Command Status Values](https://www.twilio.com/docs/wireless/api/smscommand-resource#status-values) for a description of each.
	Status *string `json:"Status,omitempty"`
	// The direction of the SMS Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`.
	Direction *string `json:"Direction,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSmsCommandParams) SetSim(Sim string) *ListSmsCommandParams {
	params.Sim = &Sim
	return params
}
func (params *ListSmsCommandParams) SetStatus(Status string) *ListSmsCommandParams {
	params.Status = &Status
	return params
}
func (params *ListSmsCommandParams) SetDirection(Direction string) *ListSmsCommandParams {
	params.Direction = &Direction
	return params
}
func (params *ListSmsCommandParams) SetPageSize(PageSize int) *ListSmsCommandParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of SMS Commands from your account.
func (c *ApiService) ListSmsCommand(params *ListSmsCommandParams) (*ListSmsCommandResponse, error) {
	path := "/v1/SmsCommands"

	data := url.Values{}
	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
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
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSmsCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
