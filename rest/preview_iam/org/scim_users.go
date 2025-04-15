/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Organization Public API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateOrganizationUser'
type CreateOrganizationUserParams struct {
	//
	ScimUser *ScimUser `json:"ScimUser,omitempty"`
}

func (params *CreateOrganizationUserParams) SetScimUser(ScimUser ScimUser) *CreateOrganizationUserParams {
	params.ScimUser = &ScimUser
	return params
}

func (c *ApiService) CreateOrganizationUser(OrganizationSid string, params *CreateOrganizationUserParams) (*ScimUser, error) {
	path := "/Organizations/{OrganizationSid}/scim/Users"
	path = strings.Replace(path, "{"+"OrganizationSid"+"}", OrganizationSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.ScimUser != nil {
		b, err := json.Marshal(*params.ScimUser)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ScimUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteOrganizationUser(OrganizationSid string, UserSid string) error {
	path := "/Organizations/{OrganizationSid}/scim/Users/{UserSid}"
	path = strings.Replace(path, "{"+"OrganizationSid"+"}", OrganizationSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchOrganizationUser(OrganizationSid string, UserSid string) (*ScimUser, error) {
	path := "/Organizations/{OrganizationSid}/scim/Users/{UserSid}"
	path = strings.Replace(path, "{"+"OrganizationSid"+"}", OrganizationSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ScimUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListOrganizationUsers'
type ListOrganizationUsersParams struct {
	//
	Filter *string `json:"filter,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListOrganizationUsersParams) SetFilter(Filter string) *ListOrganizationUsersParams {
	params.Filter = &Filter
	return params
}
func (params *ListOrganizationUsersParams) SetLimit(Limit int) *ListOrganizationUsersParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of OrganizationUsers records from the API. Request is executed immediately.
func (c *ApiService) PageOrganizationUsers(OrganizationSid string, params *ListOrganizationUsersParams, pageToken, pageNumber string) (*ScimUserPage, error) {
	path := "/Organizations/{OrganizationSid}/scim/Users"

	path = strings.Replace(path, "{"+"OrganizationSid"+"}", OrganizationSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Filter != nil {
		data.Set("filter", *params.Filter)
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

	ps := &ScimUserPage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists OrganizationUsers records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListOrganizationUsers(OrganizationSid string, params *ListOrganizationUsersParams) ([]ScimUser, error) {
	response, errors := c.StreamOrganizationUsers(OrganizationSid, params)

	records := make([]ScimUser, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams OrganizationUsers records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamOrganizationUsers(OrganizationSid string, params *ListOrganizationUsersParams) (chan ScimUser, chan error) {
	if params == nil {
		params = &ListOrganizationUsersParams{}
	}

	recordChannel := make(chan ScimUser, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageOrganizationUsers(OrganizationSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamOrganizationUsers(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamOrganizationUsers(response *ScimUserPage, params *ListOrganizationUsersParams, recordChannel chan ScimUser, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Resources
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextScimUserPage)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ScimUserPage)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextScimUserPage(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ScimUserPage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateOrganizationUser'
type UpdateOrganizationUserParams struct {
	//
	IfMatch *string `json:"If-Match,omitempty"`
	//
	ScimUser *ScimUser `json:"ScimUser,omitempty"`
}

func (params *UpdateOrganizationUserParams) SetIfMatch(IfMatch string) *UpdateOrganizationUserParams {
	params.IfMatch = &IfMatch
	return params
}
func (params *UpdateOrganizationUserParams) SetScimUser(ScimUser ScimUser) *UpdateOrganizationUserParams {
	params.ScimUser = &ScimUser
	return params
}

func (c *ApiService) UpdateOrganizationUser(OrganizationSid string, UserSid string, params *UpdateOrganizationUserParams) (*ScimUser, error) {
	path := "/Organizations/{OrganizationSid}/scim/Users/{UserSid}"
	path = strings.Replace(path, "{"+"OrganizationSid"+"}", OrganizationSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.ScimUser != nil {
		b, err := json.Marshal(*params.ScimUser)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}
	resp, err := c.requestHandler.Put(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ScimUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
