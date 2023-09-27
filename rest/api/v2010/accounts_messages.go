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
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateMessage'
type CreateMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) creating the Message resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The recipient's phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (for SMS/MMS) or [channel address](https://www.twilio.com/docs/messaging/channels), e.g. `whatsapp:+15552229999`.
	To *string `json:"To,omitempty"`
	// The URL of the endpoint to which Twilio sends [Message status callback requests](https://www.twilio.com/docs/sms/api/message-resource#twilios-request-to-the-statuscallback-url). URL must contain a valid hostname and underscores are not allowed. If you include this parameter with the `messaging_service_sid`, Twilio uses this URL instead of the Status Callback URL of the [Messaging Service](https://www.twilio.com/docs/messaging/api/service-resource).
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The SID of the associated [TwiML Application](https://www.twilio.com/docs/usage/api/applications). If this parameter is provided, the `status_callback` parameter of this request is ignored; [Message status callback requests](https://www.twilio.com/docs/sms/api/message-resource#twilios-request-to-the-statuscallback-url) are sent to the TwiML App's `message_status_callback` URL.
	ApplicationSid *string `json:"ApplicationSid,omitempty"`
	// The maximum price in US dollars that you are willing to pay for this Message's delivery. The value can have up to four decimal places. When the `max_price` parameter is provided, the cost of a message is checked before it is sent. If the cost exceeds `max_price`, the message is not sent and the Message `status` is `failed`.
	MaxPrice *float32 `json:"MaxPrice,omitempty"`
	// Boolean indicating whether or not you intend to provide delivery confirmation feedback to Twilio (used in conjunction with the [Message Feedback subresource](https://www.twilio.com/docs/sms/api/message-feedback-resource)). Default value is `false`.
	ProvideFeedback *bool `json:"ProvideFeedback,omitempty"`
	// Total number of attempts made (including this request) to send the message regardless of the provider used
	Attempt *int `json:"Attempt,omitempty"`
	// The maximum length in seconds that the Message can remain in Twilio's outgoing message queue. If a queued Message exceeds the `validity_period`, the Message is not sent. Accepted values are integers from `1` to `14400`. Default value is `14400`. A `validity_period` greater than `5` is recommended. [Learn more about the validity period](https://www.twilio.com/blog/take-more-control-of-outbound-messages-using-validity-period-html)
	ValidityPeriod *int `json:"ValidityPeriod,omitempty"`
	// Reserved
	ForceDelivery *bool `json:"ForceDelivery,omitempty"`
	//
	ContentRetention *string `json:"ContentRetention,omitempty"`
	//
	AddressRetention *string `json:"AddressRetention,omitempty"`
	// Whether to detect Unicode characters that have a similar GSM-7 character and replace them. Can be: `true` or `false`.
	SmartEncoded *bool `json:"SmartEncoded,omitempty"`
	// Rich actions for non-SMS/MMS channels. Used for [sending location in WhatsApp messages](https://www.twilio.com/docs/whatsapp/message-features#location-messages-with-whatsapp).
	PersistentAction *[]string `json:"PersistentAction,omitempty"`
	// For Messaging Services with [Link Shortening configured](https://www.twilio.com/docs/messaging/features/how-to-configure-link-shortening) only: A Boolean indicating whether or not Twilio should shorten links in the `body` of the Message. Default value is `false`. If `true`, the `messaging_service_sid` parameter must also be provided.
	ShortenUrls *bool `json:"ShortenUrls,omitempty"`
	//
	ScheduleType *string `json:"ScheduleType,omitempty"`
	// The time that Twilio will send the message. Must be in ISO 8601 format.
	SendAt *time.Time `json:"SendAt,omitempty"`
	// If set to `true`, Twilio delivers the message as a single MMS message, regardless of the presence of media.
	SendAsMms *bool `json:"SendAsMms,omitempty"`
	// For [Content Editor/API](https://www.twilio.com/docs/content) only: Key-value pairs of [Template variables](https://www.twilio.com/docs/content/using-variables-with-content-api) and their substitution values. `content_sid` parameter must also be provided. If values are not defined in the `content_variables` parameter, the [Template's default placeholder values](https://www.twilio.com/docs/content/content-api-resources#create-templates) are used.
	ContentVariables *string `json:"ContentVariables,omitempty"`
	// A string containing a JSON map of key value pairs of tags to be recorded as metadata for the message. The object may contain up to 10 tags. Keys and values can each be up to 128 characters in length.
	Tags *string `json:"Tags,omitempty"`
	//
	RiskCheck *string `json:"RiskCheck,omitempty"`
	// The sender's Twilio phone number (in [E.164](https://en.wikipedia.org/wiki/E.164) format), [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), [Wireless SIM](https://www.twilio.com/docs/iot/wireless/programmable-wireless-send-machine-machine-sms-commands), [short code](https://www.twilio.com/docs/sms/api/short-code), or [channel address](https://www.twilio.com/docs/messaging/channels) (e.g., `whatsapp:+15554449999`). The value of the `from` parameter must be a sender that is hosted within Twilio and belongs to the Account creating the Message. If you are using `messaging_service_sid`, this parameter can be empty (Twilio assigns a `from` value from the Messaging Service's Sender Pool) or you can provide a specific sender from your Sender Pool.
	From *string `json:"From,omitempty"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services) you want to associate with the Message. When this parameter is provided and the `from` parameter is omitted, Twilio selects the optimal sender from the Messaging Service's Sender Pool. You may also provide a `from` parameter if you want to use a specific Sender from the Sender Pool.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// The text content of the outgoing message. Can be up to 1,600 characters in length. SMS only: If the `body` contains more than 160 [GSM-7](https://www.twilio.com/docs/glossary/what-is-gsm-7-character-encoding) characters (or 70 [UCS-2](https://www.twilio.com/docs/glossary/what-is-ucs-2-character-encoding) characters), the message is segmented and charged accordingly. For long `body` text, consider using the [send_as_mms parameter](https://www.twilio.com/blog/mms-for-long-text-messages).
	Body *string `json:"Body,omitempty"`
	// The URL of media to include in the Message content. `jpeg`, `jpg`, `gif`, and `png` file types are fully supported by Twilio and content is formatted for delivery on destination devices. The media size limit is 5 MB for supported file types (`jpeg`, `jpg`, `png`, `gif`) and 500 KB for [other types](https://www.twilio.com/docs/sms/accepted-mime-types) of accepted media. To send more than one image in the message, provide multiple `media_url` parameters in the POST request. You can include up to ten `media_url` parameters per message. [International](https://support.twilio.com/hc/en-us/articles/223179808-Sending-and-receiving-MMS-messages) and [carrier](https://support.twilio.com/hc/en-us/articles/223133707-Is-MMS-supported-for-all-carriers-in-US-and-Canada-) limits apply.
	MediaUrl *[]string `json:"MediaUrl,omitempty"`
	// For [Content Editor/API](https://www.twilio.com/docs/content) only: The SID of the Content Template to be used with the Message, e.g., `HXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`. If this parameter is not provided, a Content Template is not used. Find the SID in the Console on the Content Editor page. For Content API users, the SID is found in Twilio's response when [creating the Template](https://www.twilio.com/docs/content/content-api-resources#create-templates) or by [fetching your Templates](https://www.twilio.com/docs/content/content-api-resources#fetch-all-content-resources).
	ContentSid *string `json:"ContentSid,omitempty"`
}

func (params *CreateMessageParams) SetPathAccountSid(PathAccountSid string) *CreateMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateMessageParams) SetTo(To string) *CreateMessageParams {
	params.To = &To
	return params
}
func (params *CreateMessageParams) SetStatusCallback(StatusCallback string) *CreateMessageParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateMessageParams) SetApplicationSid(ApplicationSid string) *CreateMessageParams {
	params.ApplicationSid = &ApplicationSid
	return params
}
func (params *CreateMessageParams) SetMaxPrice(MaxPrice float32) *CreateMessageParams {
	params.MaxPrice = &MaxPrice
	return params
}
func (params *CreateMessageParams) SetProvideFeedback(ProvideFeedback bool) *CreateMessageParams {
	params.ProvideFeedback = &ProvideFeedback
	return params
}
func (params *CreateMessageParams) SetAttempt(Attempt int) *CreateMessageParams {
	params.Attempt = &Attempt
	return params
}
func (params *CreateMessageParams) SetValidityPeriod(ValidityPeriod int) *CreateMessageParams {
	params.ValidityPeriod = &ValidityPeriod
	return params
}
func (params *CreateMessageParams) SetForceDelivery(ForceDelivery bool) *CreateMessageParams {
	params.ForceDelivery = &ForceDelivery
	return params
}
func (params *CreateMessageParams) SetContentRetention(ContentRetention string) *CreateMessageParams {
	params.ContentRetention = &ContentRetention
	return params
}
func (params *CreateMessageParams) SetAddressRetention(AddressRetention string) *CreateMessageParams {
	params.AddressRetention = &AddressRetention
	return params
}
func (params *CreateMessageParams) SetSmartEncoded(SmartEncoded bool) *CreateMessageParams {
	params.SmartEncoded = &SmartEncoded
	return params
}
func (params *CreateMessageParams) SetPersistentAction(PersistentAction []string) *CreateMessageParams {
	params.PersistentAction = &PersistentAction
	return params
}
func (params *CreateMessageParams) SetShortenUrls(ShortenUrls bool) *CreateMessageParams {
	params.ShortenUrls = &ShortenUrls
	return params
}
func (params *CreateMessageParams) SetScheduleType(ScheduleType string) *CreateMessageParams {
	params.ScheduleType = &ScheduleType
	return params
}
func (params *CreateMessageParams) SetSendAt(SendAt time.Time) *CreateMessageParams {
	params.SendAt = &SendAt
	return params
}
func (params *CreateMessageParams) SetSendAsMms(SendAsMms bool) *CreateMessageParams {
	params.SendAsMms = &SendAsMms
	return params
}
func (params *CreateMessageParams) SetContentVariables(ContentVariables string) *CreateMessageParams {
	params.ContentVariables = &ContentVariables
	return params
}
func (params *CreateMessageParams) SetTags(Tags string) *CreateMessageParams {
	params.Tags = &Tags
	return params
}
func (params *CreateMessageParams) SetRiskCheck(RiskCheck string) *CreateMessageParams {
	params.RiskCheck = &RiskCheck
	return params
}
func (params *CreateMessageParams) SetFrom(From string) *CreateMessageParams {
	params.From = &From
	return params
}
func (params *CreateMessageParams) SetMessagingServiceSid(MessagingServiceSid string) *CreateMessageParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}
func (params *CreateMessageParams) SetBody(Body string) *CreateMessageParams {
	params.Body = &Body
	return params
}
func (params *CreateMessageParams) SetMediaUrl(MediaUrl []string) *CreateMessageParams {
	params.MediaUrl = &MediaUrl
	return params
}
func (params *CreateMessageParams) SetContentSid(ContentSid string) *CreateMessageParams {
	params.ContentSid = &ContentSid
	return params
}

// Send a message
func (c *ApiService) CreateMessage(params *CreateMessageParams) (*ApiV2010Message, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.ApplicationSid != nil {
		data.Set("ApplicationSid", *params.ApplicationSid)
	}
	if params != nil && params.MaxPrice != nil {
		data.Set("MaxPrice", fmt.Sprint(*params.MaxPrice))
	}
	if params != nil && params.ProvideFeedback != nil {
		data.Set("ProvideFeedback", fmt.Sprint(*params.ProvideFeedback))
	}
	if params != nil && params.Attempt != nil {
		data.Set("Attempt", fmt.Sprint(*params.Attempt))
	}
	if params != nil && params.ValidityPeriod != nil {
		data.Set("ValidityPeriod", fmt.Sprint(*params.ValidityPeriod))
	}
	if params != nil && params.ForceDelivery != nil {
		data.Set("ForceDelivery", fmt.Sprint(*params.ForceDelivery))
	}
	if params != nil && params.ContentRetention != nil {
		data.Set("ContentRetention", *params.ContentRetention)
	}
	if params != nil && params.AddressRetention != nil {
		data.Set("AddressRetention", *params.AddressRetention)
	}
	if params != nil && params.SmartEncoded != nil {
		data.Set("SmartEncoded", fmt.Sprint(*params.SmartEncoded))
	}
	if params != nil && params.PersistentAction != nil {
		for _, item := range *params.PersistentAction {
			data.Add("PersistentAction", item)
		}
	}
	if params != nil && params.ShortenUrls != nil {
		data.Set("ShortenUrls", fmt.Sprint(*params.ShortenUrls))
	}
	if params != nil && params.ScheduleType != nil {
		data.Set("ScheduleType", *params.ScheduleType)
	}
	if params != nil && params.SendAt != nil {
		data.Set("SendAt", fmt.Sprint((*params.SendAt).Format(time.RFC3339)))
	}
	if params != nil && params.SendAsMms != nil {
		data.Set("SendAsMms", fmt.Sprint(*params.SendAsMms))
	}
	if params != nil && params.ContentVariables != nil {
		data.Set("ContentVariables", *params.ContentVariables)
	}
	if params != nil && params.Tags != nil {
		data.Set("Tags", *params.Tags)
	}
	if params != nil && params.RiskCheck != nil {
		data.Set("RiskCheck", *params.RiskCheck)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.MediaUrl != nil {
		for _, item := range *params.MediaUrl {
			data.Add("MediaUrl", item)
		}
	}
	if params != nil && params.ContentSid != nil {
		data.Set("ContentSid", *params.ContentSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteMessage'
type DeleteMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) associated with the Message resource
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteMessageParams) SetPathAccountSid(PathAccountSid string) *DeleteMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Deletes a Message resource from your account
func (c *ApiService) DeleteMessage(Sid string, params *DeleteMessageParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchMessage'
type FetchMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) associated with the Message resource
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchMessageParams) SetPathAccountSid(PathAccountSid string) *FetchMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a specific Message
func (c *ApiService) FetchMessage(Sid string, params *FetchMessageParams) (*ApiV2010Message, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMessage'
type ListMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) associated with the Message resources.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Filter by recipient. For example: Set this `to` parameter to `+15558881111` to retrieve a list of Message resources with `to` properties of `+15558881111`
	To *string `json:"To,omitempty"`
	// Filter by sender. For example: Set this `from` parameter to `+15552229999` to retrieve a list of Message resources with `from` properties of `+15552229999`
	From *string `json:"From,omitempty"`
	// Filter by Message `sent_date`. Accepts GMT dates in the following formats: `YYYY-MM-DD` (to find Messages with a specific `sent_date`), `<=YYYY-MM-DD` (to find Messages with `sent_date`s on and before a specific date), and `>=YYYY-MM-DD` (to find Messages with `sent_dates` on and after a specific date).
	DateSent *time.Time `json:"DateSent,omitempty"`
	// Filter by Message `sent_date`. Accepts GMT dates in the following formats: `YYYY-MM-DD` (to find Messages with a specific `sent_date`), `<=YYYY-MM-DD` (to find Messages with `sent_date`s on and before a specific date), and `>=YYYY-MM-DD` (to find Messages with `sent_dates` on and after a specific date).
	DateSentBefore *time.Time `json:"DateSent&lt;,omitempty"`
	// Filter by Message `sent_date`. Accepts GMT dates in the following formats: `YYYY-MM-DD` (to find Messages with a specific `sent_date`), `<=YYYY-MM-DD` (to find Messages with `sent_date`s on and before a specific date), and `>=YYYY-MM-DD` (to find Messages with `sent_dates` on and after a specific date).
	DateSentAfter *time.Time `json:"DateSent&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMessageParams) SetPathAccountSid(PathAccountSid string) *ListMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListMessageParams) SetTo(To string) *ListMessageParams {
	params.To = &To
	return params
}
func (params *ListMessageParams) SetFrom(From string) *ListMessageParams {
	params.From = &From
	return params
}
func (params *ListMessageParams) SetDateSent(DateSent time.Time) *ListMessageParams {
	params.DateSent = &DateSent
	return params
}
func (params *ListMessageParams) SetDateSentBefore(DateSentBefore time.Time) *ListMessageParams {
	params.DateSentBefore = &DateSentBefore
	return params
}
func (params *ListMessageParams) SetDateSentAfter(DateSentAfter time.Time) *ListMessageParams {
	params.DateSentAfter = &DateSentAfter
	return params
}
func (params *ListMessageParams) SetPageSize(PageSize int) *ListMessageParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMessageParams) SetLimit(Limit int) *ListMessageParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Message records from the API. Request is executed immediately.
func (c *ApiService) PageMessage(params *ListMessageParams, pageToken, pageNumber string) (*ListMessageResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.DateSent != nil {
		data.Set("DateSent", fmt.Sprint((*params.DateSent).Format(time.RFC3339)))
	}
	if params != nil && params.DateSentBefore != nil {
		data.Set("DateSent<", fmt.Sprint((*params.DateSentBefore).Format(time.RFC3339)))
	}
	if params != nil && params.DateSentAfter != nil {
		data.Set("DateSent>", fmt.Sprint((*params.DateSentAfter).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
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

	ps := &ListMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Message records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMessage(params *ListMessageParams) ([]ApiV2010Message, error) {
	response, errors := c.StreamMessage(params)

	records := make([]ApiV2010Message, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Message records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMessage(params *ListMessageParams) (chan ApiV2010Message, chan error) {
	if params == nil {
		params = &ListMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010Message, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageMessage(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamMessage(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamMessage(response *ListMessageResponse, params *ListMessageParams, recordChannel chan ApiV2010Message, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Messages
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListMessageResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListMessageResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListMessageResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateMessage'
type UpdateMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The new `body` of the Message resource. To redact the text content of a Message, this parameter's value must be an empty string
	Body *string `json:"Body,omitempty"`
	//
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateMessageParams) SetPathAccountSid(PathAccountSid string) *UpdateMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateMessageParams) SetBody(Body string) *UpdateMessageParams {
	params.Body = &Body
	return params
}
func (params *UpdateMessageParams) SetStatus(Status string) *UpdateMessageParams {
	params.Status = &Status
	return params
}

// Update a Message resource (used to redact Message `body` text and to cancel not-yet-sent messages)
func (c *ApiService) UpdateMessage(Sid string, params *UpdateMessageParams) (*ApiV2010Message, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
