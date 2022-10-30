/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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

// Optional parameters for the method 'CreateMessageFeedback'
type CreateMessageFeedbackParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	//
	Outcome *string `json:"Outcome,omitempty"`
}

func (params *CreateMessageFeedbackParams) SetPathAccountSid(PathAccountSid string) *CreateMessageFeedbackParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateMessageFeedbackParams) SetOutcome(Outcome string) *CreateMessageFeedbackParams {
	params.Outcome = &Outcome
	return params
}

//
func (c *ApiService) CreateMessageFeedback(MessageSid string, params *CreateMessageFeedbackParams) (*ApiV2010MessageFeedback, error) {
	return c.CreateMessageFeedbackWithCtx(context.TODO(), MessageSid, params)
}

//
func (c *ApiService) CreateMessageFeedbackWithCtx(ctx context.Context, MessageSid string, params *CreateMessageFeedbackParams) (*ApiV2010MessageFeedback, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Feedback.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Outcome != nil {
		data.Set("Outcome", *params.Outcome)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010MessageFeedback{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
