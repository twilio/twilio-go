/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Pricing
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
)

func (c *ApiService) FetchVoiceNumber(Number string) (*PricingV1VoiceNumber, error) {
	return c.FetchVoiceNumberWithCtx(context.TODO(), Number)
}

func (c *ApiService) FetchVoiceNumberWithCtx(ctx context.Context, Number string) (*PricingV1VoiceNumber, error) {
	path := "/v1/Voice/Numbers/{Number}"
	path = strings.Replace(path, "{"+"Number"+"}", Number, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1VoiceNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
