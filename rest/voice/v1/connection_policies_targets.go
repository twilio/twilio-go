/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
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

// Optional parameters for the method 'CreateConnectionPolicyTarget'
type CreateConnectionPolicyTargetParams struct {
	// Whether the Target is enabled. The default is `true`.
	Enabled *bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target.
	Priority *int `json:"Priority,omitempty"`
	// The SIP address you want Twilio to route your calls to. This must be a `sip:` schema. `sips` is NOT supported.
	Target *string `json:"Target,omitempty"`
	// The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority.
	Weight *int `json:"Weight,omitempty"`
}

func (params *CreateConnectionPolicyTargetParams) SetEnabled(Enabled bool) *CreateConnectionPolicyTargetParams {
	params.Enabled = &Enabled
	return params
}
func (params *CreateConnectionPolicyTargetParams) SetFriendlyName(FriendlyName string) *CreateConnectionPolicyTargetParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateConnectionPolicyTargetParams) SetPriority(Priority int) *CreateConnectionPolicyTargetParams {
	params.Priority = &Priority
	return params
}
func (params *CreateConnectionPolicyTargetParams) SetTarget(Target string) *CreateConnectionPolicyTargetParams {
	params.Target = &Target
	return params
}
func (params *CreateConnectionPolicyTargetParams) SetWeight(Weight int) *CreateConnectionPolicyTargetParams {
	params.Weight = &Weight
	return params
}

func (c *ApiService) CreateConnectionPolicyTarget(ConnectionPolicySid string, params *CreateConnectionPolicyTargetParams) (*VoiceV1ConnectionPolicyConnectionPolicyTarget, error) {
	path := "/v1/ConnectionPolicies/{ConnectionPolicySid}/Targets"
	path = strings.Replace(path, "{"+"ConnectionPolicySid"+"}", ConnectionPolicySid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Priority != nil {
		data.Set("Priority", fmt.Sprint(*params.Priority))
	}
	if params != nil && params.Target != nil {
		data.Set("Target", *params.Target)
	}
	if params != nil && params.Weight != nil {
		data.Set("Weight", fmt.Sprint(*params.Weight))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicyConnectionPolicyTarget{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteConnectionPolicyTarget(ConnectionPolicySid string, Sid string) error {
	path := "/v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid}"
	path = strings.Replace(path, "{"+"ConnectionPolicySid"+"}", ConnectionPolicySid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchConnectionPolicyTarget(ConnectionPolicySid string, Sid string) (*VoiceV1ConnectionPolicyConnectionPolicyTarget, error) {
	path := "/v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid}"
	path = strings.Replace(path, "{"+"ConnectionPolicySid"+"}", ConnectionPolicySid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicyConnectionPolicyTarget{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConnectionPolicyTarget'
type ListConnectionPolicyTargetParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListConnectionPolicyTargetParams) SetPageSize(PageSize int) *ListConnectionPolicyTargetParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListConnectionPolicyTarget(ConnectionPolicySid string, params *ListConnectionPolicyTargetParams) (*ListConnectionPolicyTargetResponse, error) {
	path := "/v1/ConnectionPolicies/{ConnectionPolicySid}/Targets"
	path = strings.Replace(path, "{"+"ConnectionPolicySid"+"}", ConnectionPolicySid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConnectionPolicyTargetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateConnectionPolicyTarget'
type UpdateConnectionPolicyTargetParams struct {
	// Whether the Target is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The relative importance of the target. Can be an integer from 0 to 65535, inclusive. The lowest number represents the most important target.
	Priority *int `json:"Priority,omitempty"`
	// The SIP address you want Twilio to route your calls to. This must be a `sip:` schema. `sips` is NOT supported.
	Target *string `json:"Target,omitempty"`
	// The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive. Targets with higher values receive more load than those with lower ones with the same priority.
	Weight *int `json:"Weight,omitempty"`
}

func (params *UpdateConnectionPolicyTargetParams) SetEnabled(Enabled bool) *UpdateConnectionPolicyTargetParams {
	params.Enabled = &Enabled
	return params
}
func (params *UpdateConnectionPolicyTargetParams) SetFriendlyName(FriendlyName string) *UpdateConnectionPolicyTargetParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateConnectionPolicyTargetParams) SetPriority(Priority int) *UpdateConnectionPolicyTargetParams {
	params.Priority = &Priority
	return params
}
func (params *UpdateConnectionPolicyTargetParams) SetTarget(Target string) *UpdateConnectionPolicyTargetParams {
	params.Target = &Target
	return params
}
func (params *UpdateConnectionPolicyTargetParams) SetWeight(Weight int) *UpdateConnectionPolicyTargetParams {
	params.Weight = &Weight
	return params
}

func (c *ApiService) UpdateConnectionPolicyTarget(ConnectionPolicySid string, Sid string, params *UpdateConnectionPolicyTargetParams) (*VoiceV1ConnectionPolicyConnectionPolicyTarget, error) {
	path := "/v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid}"
	path = strings.Replace(path, "{"+"ConnectionPolicySid"+"}", ConnectionPolicySid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Priority != nil {
		data.Set("Priority", fmt.Sprint(*params.Priority))
	}
	if params != nil && params.Target != nil {
		data.Set("Target", *params.Target)
	}
	if params != nil && params.Weight != nil {
		data.Set("Weight", fmt.Sprint(*params.Weight))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicyConnectionPolicyTarget{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
