/*
 * Twilio - Events
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

// Fetch a specific schema and version.
func (c *ApiService) FetchSchemaVersion(Id string, SchemaVersion int) (*EventsV1SchemaSchemaVersion, error) {
	path := "/v1/Schemas/{Id}/Versions/{SchemaVersion}"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)
	path = strings.Replace(path, "{"+"SchemaVersion"+"}", fmt.Sprint(SchemaVersion), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SchemaSchemaVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSchemaVersion'
type ListSchemaVersionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSchemaVersionParams) SetPageSize(PageSize int) *ListSchemaVersionParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a paginated list of versions of the schema.
func (c *ApiService) ListSchemaVersion(Id string, params *ListSchemaVersionParams) (*ListSchemaVersionResponse, error) {
	path := "/v1/Schemas/{Id}/Versions"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)

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

	ps := &ListSchemaVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
