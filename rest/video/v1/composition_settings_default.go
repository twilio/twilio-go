/*
 * Twilio - Video
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
)

// Optional parameters for the method 'CreateCompositionSettings'
type CreateCompositionSettingsParams struct {
	// The SID of the stored Credential resource.
	AwsCredentialsSid *string `json:"AwsCredentialsSid,omitempty"`
	// The URL of the AWS S3 bucket where the compositions should be stored. We only support DNS-compliant URLs like `https://documentation-example-twilio-bucket/compositions`, where `compositions` is the path in which you want the compositions to be stored. This URL accepts only URI-valid characters, as described in the <a href='https://tools.ietf.org/html/rfc3986#section-2'>RFC 3986</a>.
	AwsS3Url *string `json:"AwsS3Url,omitempty"`
	// Whether all compositions should be written to the `aws_s3_url`. When `false`, all compositions are stored in our cloud.
	AwsStorageEnabled *bool `json:"AwsStorageEnabled,omitempty"`
	// Whether all compositions should be stored in an encrypted form. The default is `false`.
	EncryptionEnabled *bool `json:"EncryptionEnabled,omitempty"`
	// The SID of the Public Key resource to use for encryption.
	EncryptionKeySid *string `json:"EncryptionKeySid,omitempty"`
	// A descriptive string that you create to describe the resource and show to the user in the console
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateCompositionSettingsParams) SetAwsCredentialsSid(AwsCredentialsSid string) *CreateCompositionSettingsParams {
	params.AwsCredentialsSid = &AwsCredentialsSid
	return params
}
func (params *CreateCompositionSettingsParams) SetAwsS3Url(AwsS3Url string) *CreateCompositionSettingsParams {
	params.AwsS3Url = &AwsS3Url
	return params
}
func (params *CreateCompositionSettingsParams) SetAwsStorageEnabled(AwsStorageEnabled bool) *CreateCompositionSettingsParams {
	params.AwsStorageEnabled = &AwsStorageEnabled
	return params
}
func (params *CreateCompositionSettingsParams) SetEncryptionEnabled(EncryptionEnabled bool) *CreateCompositionSettingsParams {
	params.EncryptionEnabled = &EncryptionEnabled
	return params
}
func (params *CreateCompositionSettingsParams) SetEncryptionKeySid(EncryptionKeySid string) *CreateCompositionSettingsParams {
	params.EncryptionKeySid = &EncryptionKeySid
	return params
}
func (params *CreateCompositionSettingsParams) SetFriendlyName(FriendlyName string) *CreateCompositionSettingsParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateCompositionSettings(params *CreateCompositionSettingsParams) (*VideoV1CompositionSettings, error) {
	path := "/v1/CompositionSettings/Default"

	data := url.Values{}
	if params != nil && params.AwsCredentialsSid != nil {
		data.Set("AwsCredentialsSid", *params.AwsCredentialsSid)
	}
	if params != nil && params.AwsS3Url != nil {
		data.Set("AwsS3Url", *params.AwsS3Url)
	}
	if params != nil && params.AwsStorageEnabled != nil {
		data.Set("AwsStorageEnabled", fmt.Sprint(*params.AwsStorageEnabled))
	}
	if params != nil && params.EncryptionEnabled != nil {
		data.Set("EncryptionEnabled", fmt.Sprint(*params.EncryptionEnabled))
	}
	if params != nil && params.EncryptionKeySid != nil {
		data.Set("EncryptionKeySid", *params.EncryptionKeySid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1CompositionSettings{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchCompositionSettings() (*VideoV1CompositionSettings, error) {
	path := "/v1/CompositionSettings/Default"

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1CompositionSettings{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
