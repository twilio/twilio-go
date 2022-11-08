/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Media
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
	"fmt"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreatePlayerStreamerPlaybackGrant'
type CreatePlayerStreamerPlaybackGrantParams struct {
	// The time to live of the PlaybackGrant. Default value is 15 seconds. Maximum value is 60 seconds.
	Ttl *int `json:"Ttl,omitempty"`
	// The full origin URL where the livestream can be streamed. If this is not provided, it can be streamed from any domain.
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty"`
}

func (params *CreatePlayerStreamerPlaybackGrantParams) SetTtl(Ttl int) *CreatePlayerStreamerPlaybackGrantParams {
	params.Ttl = &Ttl
	return params
}
func (params *CreatePlayerStreamerPlaybackGrantParams) SetAccessControlAllowOrigin(AccessControlAllowOrigin string) *CreatePlayerStreamerPlaybackGrantParams {
	params.AccessControlAllowOrigin = &AccessControlAllowOrigin
	return params
}

func (c *ApiService) CreatePlayerStreamerPlaybackGrant(Sid string, params *CreatePlayerStreamerPlaybackGrantParams) (*MediaV1PlayerStreamerPlaybackGrant, error) {
	return c.CreatePlayerStreamerPlaybackGrantWithCtx(context.TODO(), Sid, params)
}

func (c *ApiService) CreatePlayerStreamerPlaybackGrantWithCtx(ctx context.Context, Sid string, params *CreatePlayerStreamerPlaybackGrantParams) (*MediaV1PlayerStreamerPlaybackGrant, error) {
	path := "/v1/PlayerStreamers/{Sid}/PlaybackGrant"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.AccessControlAllowOrigin != nil {
		data.Set("AccessControlAllowOrigin", *params.AccessControlAllowOrigin)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MediaV1PlayerStreamerPlaybackGrant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// **This method is not enabled.** Returns a single PlaybackGrant resource identified by a SID.
func (c *ApiService) FetchPlayerStreamerPlaybackGrant(Sid string) (*MediaV1PlayerStreamerPlaybackGrant, error) {
	return c.FetchPlayerStreamerPlaybackGrantWithCtx(context.TODO(), Sid)
}

// **This method is not enabled.** Returns a single PlaybackGrant resource identified by a SID.
func (c *ApiService) FetchPlayerStreamerPlaybackGrantWithCtx(ctx context.Context, Sid string) (*MediaV1PlayerStreamerPlaybackGrant, error) {
	path := "/v1/PlayerStreamers/{Sid}/PlaybackGrant"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MediaV1PlayerStreamerPlaybackGrant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
