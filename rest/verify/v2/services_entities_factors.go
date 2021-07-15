/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
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

// Optional parameters for the method 'CreateNewFactor'
type CreateNewFactorParams struct {
	// The algorithm used when `factor_type` is `push`. Algorithm supported: `ES256`
	BindingAlg *string `json:"Binding.Alg,omitempty"`
	// The Ecdsa public key in PKIX, ASN.1 DER format encoded in Base64.  Required when `factor_type` is `push`
	BindingPublicKey *string `json:"Binding.PublicKey,omitempty"`
	// The shared secret for TOTP factors encoded in Base32. This can be provided when creating the Factor, otherwise it will be generated.  Used when `factor_type` is `totp`
	BindingSecret *string `json:"Binding.Secret,omitempty"`
	// The algorithm used to derive the TOTP codes. Can be `sha1`, `sha256` or `sha512`. Defaults to `sha1`.  Used when `factor_type` is `totp`
	ConfigAlg *string `json:"Config.Alg,omitempty"`
	// The ID that uniquely identifies your app in the Google or Apple store, such as `com.example.myapp`. It can be up to 100 characters long.  Required when `factor_type` is `push`.
	ConfigAppId *string `json:"Config.AppId,omitempty"`
	// Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. The default value is defined at the service level in the property `totp.code_length`. If not configured defaults to 6.  Used when `factor_type` is `totp`
	ConfigCodeLength *int `json:"Config.CodeLength,omitempty"`
	// The transport technology used to generate the Notification Token. Can be `apn` or `fcm`.  Required when `factor_type` is `push`.
	ConfigNotificationPlatform *string `json:"Config.NotificationPlatform,omitempty"`
	// For APN, the device token. For FCM the registration token. It used to send the push notifications. Must be between 32 and 255 characters long.  Required when `factor_type` is `push`.
	ConfigNotificationToken *string `json:"Config.NotificationToken,omitempty"`
	// The Verify Push SDK version used to configure the factor  Required when `factor_type` is `push`
	ConfigSdkVersion *string `json:"Config.SdkVersion,omitempty"`
	// The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. The default value is defined at the service level in the property `totp.skew`. If not configured defaults to 1.  Used when `factor_type` is `totp`
	ConfigSkew *int `json:"Config.Skew,omitempty"`
	// Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. The default value is defined at the service level in the property `totp.time_step`. Defaults to 30 seconds if not configured.  Used when `factor_type` is `totp`
	ConfigTimeStep *int `json:"Config.TimeStep,omitempty"`
	// The Type of this Factor. Currently `push` and `totp` are supported. For `totp` to work, you need to contact [Twilio sales](https://www.twilio.com/help/sales) first to have the Verify TOTP feature enabled for your Twilio account.
	FactorType *string `json:"FactorType,omitempty"`
	// The friendly name of this Factor. This can be any string up to 64 characters, meant for humans to distinguish between Factors. For `factor_type` `push`, this could be a device name. For `factor_type` `totp`, this value is used as the “account name” in constructing the `binding.uri` property. At the same time, we recommend avoiding providing PII.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateNewFactorParams) SetBindingAlg(BindingAlg string) *CreateNewFactorParams {
	params.BindingAlg = &BindingAlg
	return params
}
func (params *CreateNewFactorParams) SetBindingPublicKey(BindingPublicKey string) *CreateNewFactorParams {
	params.BindingPublicKey = &BindingPublicKey
	return params
}
func (params *CreateNewFactorParams) SetBindingSecret(BindingSecret string) *CreateNewFactorParams {
	params.BindingSecret = &BindingSecret
	return params
}
func (params *CreateNewFactorParams) SetConfigAlg(ConfigAlg string) *CreateNewFactorParams {
	params.ConfigAlg = &ConfigAlg
	return params
}
func (params *CreateNewFactorParams) SetConfigAppId(ConfigAppId string) *CreateNewFactorParams {
	params.ConfigAppId = &ConfigAppId
	return params
}
func (params *CreateNewFactorParams) SetConfigCodeLength(ConfigCodeLength int) *CreateNewFactorParams {
	params.ConfigCodeLength = &ConfigCodeLength
	return params
}
func (params *CreateNewFactorParams) SetConfigNotificationPlatform(ConfigNotificationPlatform string) *CreateNewFactorParams {
	params.ConfigNotificationPlatform = &ConfigNotificationPlatform
	return params
}
func (params *CreateNewFactorParams) SetConfigNotificationToken(ConfigNotificationToken string) *CreateNewFactorParams {
	params.ConfigNotificationToken = &ConfigNotificationToken
	return params
}
func (params *CreateNewFactorParams) SetConfigSdkVersion(ConfigSdkVersion string) *CreateNewFactorParams {
	params.ConfigSdkVersion = &ConfigSdkVersion
	return params
}
func (params *CreateNewFactorParams) SetConfigSkew(ConfigSkew int) *CreateNewFactorParams {
	params.ConfigSkew = &ConfigSkew
	return params
}
func (params *CreateNewFactorParams) SetConfigTimeStep(ConfigTimeStep int) *CreateNewFactorParams {
	params.ConfigTimeStep = &ConfigTimeStep
	return params
}
func (params *CreateNewFactorParams) SetFactorType(FactorType string) *CreateNewFactorParams {
	params.FactorType = &FactorType
	return params
}
func (params *CreateNewFactorParams) SetFriendlyName(FriendlyName string) *CreateNewFactorParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new Factor for the Entity
func (c *ApiService) CreateNewFactor(ServiceSid string, Identity string, params *CreateNewFactorParams) (*VerifyV2ServiceEntityNewFactor, error) {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}/Factors"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BindingAlg != nil {
		data.Set("Binding.Alg", *params.BindingAlg)
	}
	if params != nil && params.BindingPublicKey != nil {
		data.Set("Binding.PublicKey", *params.BindingPublicKey)
	}
	if params != nil && params.BindingSecret != nil {
		data.Set("Binding.Secret", *params.BindingSecret)
	}
	if params != nil && params.ConfigAlg != nil {
		data.Set("Config.Alg", *params.ConfigAlg)
	}
	if params != nil && params.ConfigAppId != nil {
		data.Set("Config.AppId", *params.ConfigAppId)
	}
	if params != nil && params.ConfigCodeLength != nil {
		data.Set("Config.CodeLength", fmt.Sprint(*params.ConfigCodeLength))
	}
	if params != nil && params.ConfigNotificationPlatform != nil {
		data.Set("Config.NotificationPlatform", *params.ConfigNotificationPlatform)
	}
	if params != nil && params.ConfigNotificationToken != nil {
		data.Set("Config.NotificationToken", *params.ConfigNotificationToken)
	}
	if params != nil && params.ConfigSdkVersion != nil {
		data.Set("Config.SdkVersion", *params.ConfigSdkVersion)
	}
	if params != nil && params.ConfigSkew != nil {
		data.Set("Config.Skew", fmt.Sprint(*params.ConfigSkew))
	}
	if params != nil && params.ConfigTimeStep != nil {
		data.Set("Config.TimeStep", fmt.Sprint(*params.ConfigTimeStep))
	}
	if params != nil && params.FactorType != nil {
		data.Set("FactorType", *params.FactorType)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2ServiceEntityNewFactor{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Factor.
func (c *ApiService) DeleteFactor(ServiceSid string, Identity string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)
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

// Fetch a specific Factor.
func (c *ApiService) FetchFactor(ServiceSid string, Identity string, Sid string) (*VerifyV2ServiceEntityFactor, error) {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid}"
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

	ps := &VerifyV2ServiceEntityFactor{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFactor'
type ListFactorParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListFactorParams) SetPageSize(PageSize int) *ListFactorParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Factors for an Entity.
func (c *ApiService) ListFactor(ServiceSid string, Identity string, params *ListFactorParams) (*ListFactorResponse, error) {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}/Factors"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

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

	ps := &ListFactorResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateFactor'
type UpdateFactorParams struct {
	// The optional payload needed to verify the Factor for the first time. E.g. for a TOTP, the numeric code.
	AuthPayload *string `json:"AuthPayload,omitempty"`
	// The algorithm used to derive the TOTP codes. Can be `sha1`, `sha256` or `sha512`
	ConfigAlg *string `json:"Config.Alg,omitempty"`
	// Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive
	ConfigCodeLength *int `json:"Config.CodeLength,omitempty"`
	// For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when `factor_type` is `push`. If specified, this value must be between 32 and 255 characters long.
	ConfigNotificationToken *string `json:"Config.NotificationToken,omitempty"`
	// The Verify Push SDK version used to configure the factor
	ConfigSdkVersion *string `json:"Config.SdkVersion,omitempty"`
	// The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive
	ConfigSkew *int `json:"Config.Skew,omitempty"`
	// Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive
	ConfigTimeStep *int `json:"Config.TimeStep,omitempty"`
	// The new friendly name of this Factor. It can be up to 64 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateFactorParams) SetAuthPayload(AuthPayload string) *UpdateFactorParams {
	params.AuthPayload = &AuthPayload
	return params
}
func (params *UpdateFactorParams) SetConfigAlg(ConfigAlg string) *UpdateFactorParams {
	params.ConfigAlg = &ConfigAlg
	return params
}
func (params *UpdateFactorParams) SetConfigCodeLength(ConfigCodeLength int) *UpdateFactorParams {
	params.ConfigCodeLength = &ConfigCodeLength
	return params
}
func (params *UpdateFactorParams) SetConfigNotificationToken(ConfigNotificationToken string) *UpdateFactorParams {
	params.ConfigNotificationToken = &ConfigNotificationToken
	return params
}
func (params *UpdateFactorParams) SetConfigSdkVersion(ConfigSdkVersion string) *UpdateFactorParams {
	params.ConfigSdkVersion = &ConfigSdkVersion
	return params
}
func (params *UpdateFactorParams) SetConfigSkew(ConfigSkew int) *UpdateFactorParams {
	params.ConfigSkew = &ConfigSkew
	return params
}
func (params *UpdateFactorParams) SetConfigTimeStep(ConfigTimeStep int) *UpdateFactorParams {
	params.ConfigTimeStep = &ConfigTimeStep
	return params
}
func (params *UpdateFactorParams) SetFriendlyName(FriendlyName string) *UpdateFactorParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Update a specific Factor. This endpoint can be used to Verify a Factor if passed an &#x60;AuthPayload&#x60; param.
func (c *ApiService) UpdateFactor(ServiceSid string, Identity string, Sid string, params *UpdateFactorParams) (*VerifyV2ServiceEntityFactor, error) {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AuthPayload != nil {
		data.Set("AuthPayload", *params.AuthPayload)
	}
	if params != nil && params.ConfigAlg != nil {
		data.Set("Config.Alg", *params.ConfigAlg)
	}
	if params != nil && params.ConfigCodeLength != nil {
		data.Set("Config.CodeLength", fmt.Sprint(*params.ConfigCodeLength))
	}
	if params != nil && params.ConfigNotificationToken != nil {
		data.Set("Config.NotificationToken", *params.ConfigNotificationToken)
	}
	if params != nil && params.ConfigSdkVersion != nil {
		data.Set("Config.SdkVersion", *params.ConfigSdkVersion)
	}
	if params != nil && params.ConfigSkew != nil {
		data.Set("Config.Skew", fmt.Sprint(*params.ConfigSkew))
	}
	if params != nil && params.ConfigTimeStep != nil {
		data.Set("Config.TimeStep", fmt.Sprint(*params.ConfigTimeStep))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2ServiceEntityFactor{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
