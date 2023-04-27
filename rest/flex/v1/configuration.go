/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

    "github.com/twilio/twilio-go/client"
)


// Optional parameters for the method 'FetchConfiguration'
type FetchConfigurationParams struct {
    // The Pinned UI version of the Configuration resource to fetch.
    UiVersion *string `json:"UiVersion,omitempty"`
}

func (params *FetchConfigurationParams) SetUiVersion(UiVersion string) (*FetchConfigurationParams){
    params.UiVersion = &UiVersion
    return params
}

// 
func (c *ApiService) FetchConfiguration(params *FetchConfigurationParams) (*FlexV1Configuration, error) {
    path := "/v1/Configuration"
    
data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.UiVersion != nil {
    data.Set("UiVersion", *params.UiVersion)
}



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &FlexV1Configuration{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
