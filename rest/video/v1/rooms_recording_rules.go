/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
 * This is the public Twilio REST API.
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
)

// Returns a list of Recording Rules for the Room.
func (c *ApiService) FetchRoomRecordingRule(RoomSid string) (*VideoV1RoomRecordingRule, error) {
	path := "/v1/Rooms/{RoomSid}/RecordingRules"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomRecordingRule{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateRoomRecordingRule'
type UpdateRoomRecordingRuleParams struct {
	// A JSON-encoded array of recording rules.
	Rules *map[string]interface{} `json:"Rules,omitempty"`
}

func (params *UpdateRoomRecordingRuleParams) SetRules(Rules map[string]interface{}) *UpdateRoomRecordingRuleParams {
	params.Rules = &Rules
	return params
}

// Update the Recording Rules for the Room
func (c *ApiService) UpdateRoomRecordingRule(RoomSid string, params *UpdateRoomRecordingRuleParams) (*VideoV1RoomRecordingRule, error) {
	path := "/v1/Rooms/{RoomSid}/RecordingRules"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Rules != nil {
		v, err := json.Marshal(params.Rules)

		if err != nil {
			return nil, err
		}

		data.Set("Rules", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomRecordingRule{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
