/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://wireless.twilio.com",
	}
}

// Optional parameters for the method 'CreateCommand'
type CreateCommandParams struct {
	// The HTTP method we use to call `callback_url`. Can be: `POST` or `GET`, and the default is `POST`.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we call using the `callback_url` when the Command has finished sending, whether the command was delivered or it failed.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The message body of the Command. Can be plain text in text mode or a Base64 encoded byte string in binary mode.
	Command *string `json:"Command,omitempty"`
	// The mode to use when sending the SMS message. Can be: `text` or `binary`. The default SMS mode is `text`.
	CommandMode *string `json:"CommandMode,omitempty"`
	// Whether to request delivery receipt from the recipient. For Commands that request delivery receipt, the Command state transitions to 'delivered' once the server has received a delivery receipt from the device. The default value is `true`.
	DeliveryReceiptRequested *bool `json:"DeliveryReceiptRequested,omitempty"`
	// Whether to include the SID of the command in the message body. Can be: `none`, `start`, or `end`, and the default behavior is `none`. When sending a Command to a SIM in text mode, we can automatically include the SID of the Command in the message body, which could be used to ensure that the device does not process the same Command more than once.  A value of `start` will prepend the message with the Command SID, and `end` will append it to the end, separating the Command SID from the message body with a space. The length of the Command SID is included in the 160 character limit so the SMS body must be 128 characters or less before the Command SID is included.
	IncludeSid *string `json:"IncludeSid,omitempty"`
	// The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
	Sim *string `json:"Sim,omitempty"`
}

func (params *CreateCommandParams) SetCallbackMethod(CallbackMethod string) *CreateCommandParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *CreateCommandParams) SetCallbackUrl(CallbackUrl string) *CreateCommandParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateCommandParams) SetCommand(Command string) *CreateCommandParams {
	params.Command = &Command
	return params
}
func (params *CreateCommandParams) SetCommandMode(CommandMode string) *CreateCommandParams {
	params.CommandMode = &CommandMode
	return params
}
func (params *CreateCommandParams) SetDeliveryReceiptRequested(DeliveryReceiptRequested bool) *CreateCommandParams {
	params.DeliveryReceiptRequested = &DeliveryReceiptRequested
	return params
}
func (params *CreateCommandParams) SetIncludeSid(IncludeSid string) *CreateCommandParams {
	params.IncludeSid = &IncludeSid
	return params
}
func (params *CreateCommandParams) SetSim(Sim string) *CreateCommandParams {
	params.Sim = &Sim
	return params
}

// Send a Command to a Sim.
func (c *DefaultApiService) CreateCommand(params *CreateCommandParams) (*WirelessV1Command, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Command != nil {
		data.Set("Command", *params.Command)
	}
	if params != nil && params.CommandMode != nil {
		data.Set("CommandMode", *params.CommandMode)
	}
	if params != nil && params.DeliveryReceiptRequested != nil {
		data.Set("DeliveryReceiptRequested", fmt.Sprint(*params.DeliveryReceiptRequested))
	}
	if params != nil && params.IncludeSid != nil {
		data.Set("IncludeSid", *params.IncludeSid)
	}
	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'CreateRatePlan'
type CreateRatePlanParams struct {
	// Whether SIMs can use GPRS/3G/4G/LTE data connectivity.
	DataEnabled *bool `json:"DataEnabled,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month on the home network (T-Mobile USA). The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB and the default value is `1000`.
	DataLimit *int32 `json:"DataLimit,omitempty"`
	// The model used to meter data usage. Can be: `payg` and `quota-1`, `quota-10`, and `quota-50`. Learn more about the available [data metering models](https://www.twilio.com/docs/wireless/api/rateplan-resource#payg-vs-quota-data-plans).
	DataMetering *string `json:"DataMetering,omitempty"`
	// A descriptive string that you create to describe the resource. It does not have to be unique.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The list of services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States. Can be: `data`, `voice`, and `messaging`.
	InternationalRoaming *[]string `json:"InternationalRoaming,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States. Can be up to 2TB.
	InternationalRoamingDataLimit *int32 `json:"InternationalRoamingDataLimit,omitempty"`
	// Whether SIMs can make, send, and receive SMS using [Commands](https://www.twilio.com/docs/wireless/api/command-resource).
	MessagingEnabled *bool `json:"MessagingEnabled,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month on non-home networks in the United States. The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming) for more info.
	NationalRoamingDataLimit *int32 `json:"NationalRoamingDataLimit,omitempty"`
	// Whether SIMs can roam on networks other than the home network (T-Mobile USA) in the United States. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming).
	NationalRoamingEnabled *bool `json:"NationalRoamingEnabled,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
	// Whether SIMs can make and receive voice calls.
	VoiceEnabled *bool `json:"VoiceEnabled,omitempty"`
}

func (params *CreateRatePlanParams) SetDataEnabled(DataEnabled bool) *CreateRatePlanParams {
	params.DataEnabled = &DataEnabled
	return params
}
func (params *CreateRatePlanParams) SetDataLimit(DataLimit int32) *CreateRatePlanParams {
	params.DataLimit = &DataLimit
	return params
}
func (params *CreateRatePlanParams) SetDataMetering(DataMetering string) *CreateRatePlanParams {
	params.DataMetering = &DataMetering
	return params
}
func (params *CreateRatePlanParams) SetFriendlyName(FriendlyName string) *CreateRatePlanParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateRatePlanParams) SetInternationalRoaming(InternationalRoaming []string) *CreateRatePlanParams {
	params.InternationalRoaming = &InternationalRoaming
	return params
}
func (params *CreateRatePlanParams) SetInternationalRoamingDataLimit(InternationalRoamingDataLimit int32) *CreateRatePlanParams {
	params.InternationalRoamingDataLimit = &InternationalRoamingDataLimit
	return params
}
func (params *CreateRatePlanParams) SetMessagingEnabled(MessagingEnabled bool) *CreateRatePlanParams {
	params.MessagingEnabled = &MessagingEnabled
	return params
}
func (params *CreateRatePlanParams) SetNationalRoamingDataLimit(NationalRoamingDataLimit int32) *CreateRatePlanParams {
	params.NationalRoamingDataLimit = &NationalRoamingDataLimit
	return params
}
func (params *CreateRatePlanParams) SetNationalRoamingEnabled(NationalRoamingEnabled bool) *CreateRatePlanParams {
	params.NationalRoamingEnabled = &NationalRoamingEnabled
	return params
}
func (params *CreateRatePlanParams) SetUniqueName(UniqueName string) *CreateRatePlanParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateRatePlanParams) SetVoiceEnabled(VoiceEnabled bool) *CreateRatePlanParams {
	params.VoiceEnabled = &VoiceEnabled
	return params
}

func (c *DefaultApiService) CreateRatePlan(params *CreateRatePlanParams) (*WirelessV1RatePlan, error) {
	path := "/v1/RatePlans"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DataEnabled != nil {
		data.Set("DataEnabled", fmt.Sprint(*params.DataEnabled))
	}
	if params != nil && params.DataLimit != nil {
		data.Set("DataLimit", fmt.Sprint(*params.DataLimit))
	}
	if params != nil && params.DataMetering != nil {
		data.Set("DataMetering", *params.DataMetering)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.InternationalRoaming != nil {
		data.Set("InternationalRoaming", strings.Join(*params.InternationalRoaming, ","))
	}
	if params != nil && params.InternationalRoamingDataLimit != nil {
		data.Set("InternationalRoamingDataLimit", fmt.Sprint(*params.InternationalRoamingDataLimit))
	}
	if params != nil && params.MessagingEnabled != nil {
		data.Set("MessagingEnabled", fmt.Sprint(*params.MessagingEnabled))
	}
	if params != nil && params.NationalRoamingDataLimit != nil {
		data.Set("NationalRoamingDataLimit", fmt.Sprint(*params.NationalRoamingDataLimit))
	}
	if params != nil && params.NationalRoamingEnabled != nil {
		data.Set("NationalRoamingEnabled", fmt.Sprint(*params.NationalRoamingEnabled))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.VoiceEnabled != nil {
		data.Set("VoiceEnabled", fmt.Sprint(*params.VoiceEnabled))
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1RatePlan{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Command instance from your account.
func (c *DefaultApiService) DeleteCommand(Sid string) error {
	path := "/v1/Commands/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *DefaultApiService) DeleteRatePlan(Sid string) error {
	path := "/v1/RatePlans/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Delete a Sim resource on your Account.
func (c *DefaultApiService) DeleteSim(Sid string) error {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a Command instance from your account.
func (c *DefaultApiService) FetchCommand(Sid string) (*WirelessV1Command, error) {
	path := "/v1/Commands/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *DefaultApiService) FetchRatePlan(Sid string) (*WirelessV1RatePlan, error) {
	path := "/v1/RatePlans/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1RatePlan{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch a Sim resource on your Account.
func (c *DefaultApiService) FetchSim(Sid string) (*WirelessV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAccountUsageRecord'
type ListAccountUsageRecordParams struct {
	// Only include usage that has occurred on or before this date. Format is [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html).
	End *time.Time `json:"End,omitempty"`
	// Only include usage that has occurred on or after this date. Format is [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html).
	Start *time.Time `json:"Start,omitempty"`
	// How to summarize the usage by time. Can be: `daily`, `hourly`, or `all`. A value of `all` returns one Usage Record that describes the usage for the entire period.
	Granularity *string `json:"Granularity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListAccountUsageRecordParams) SetEnd(End time.Time) *ListAccountUsageRecordParams {
	params.End = &End
	return params
}
func (params *ListAccountUsageRecordParams) SetStart(Start time.Time) *ListAccountUsageRecordParams {
	params.Start = &Start
	return params
}
func (params *ListAccountUsageRecordParams) SetGranularity(Granularity string) *ListAccountUsageRecordParams {
	params.Granularity = &Granularity
	return params
}
func (params *ListAccountUsageRecordParams) SetPageSize(PageSize int32) *ListAccountUsageRecordParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListAccountUsageRecord(params *ListAccountUsageRecordParams) (*ListAccountUsageRecordResponse, error) {
	path := "/v1/UsageRecords"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.End != nil {
		data.Set("End", fmt.Sprint((*params.End).Format(time.RFC3339)))
	}
	if params != nil && params.Start != nil {
		data.Set("Start", fmt.Sprint((*params.Start).Format(time.RFC3339)))
	}
	if params != nil && params.Granularity != nil {
		data.Set("Granularity", *params.Granularity)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAccountUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCommand'
type ListCommandParams struct {
	// The `sid` or `unique_name` of the [Sim resources](https://www.twilio.com/docs/wireless/api/sim-resource) to read.
	Sim *string `json:"Sim,omitempty"`
	// The status of the resources to read. Can be: `queued`, `sent`, `delivered`, `received`, or `failed`.
	Status *string `json:"Status,omitempty"`
	// Only return Commands with this direction value.
	Direction *string `json:"Direction,omitempty"`
	// Only return Commands with this transport value. Can be: `sms` or `ip`.
	Transport *string `json:"Transport,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListCommandParams) SetSim(Sim string) *ListCommandParams {
	params.Sim = &Sim
	return params
}
func (params *ListCommandParams) SetStatus(Status string) *ListCommandParams {
	params.Status = &Status
	return params
}
func (params *ListCommandParams) SetDirection(Direction string) *ListCommandParams {
	params.Direction = &Direction
	return params
}
func (params *ListCommandParams) SetTransport(Transport string) *ListCommandParams {
	params.Transport = &Transport
	return params
}
func (params *ListCommandParams) SetPageSize(PageSize int32) *ListCommandParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of Commands from your account.
func (c *DefaultApiService) ListCommand(params *ListCommandParams) (*ListCommandResponse, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.Transport != nil {
		data.Set("Transport", *params.Transport)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListDataSession'
type ListDataSessionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListDataSessionParams) SetPageSize(PageSize int32) *ListDataSessionParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListDataSession(SimSid string, params *ListDataSessionParams) (*ListDataSessionResponse, error) {
	path := "/v1/Sims/{SimSid}/DataSessions"
	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDataSessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRatePlan'
type ListRatePlanParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListRatePlanParams) SetPageSize(PageSize int32) *ListRatePlanParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListRatePlan(params *ListRatePlanParams) (*ListRatePlanResponse, error) {
	path := "/v1/RatePlans"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRatePlanResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSim'
type ListSimParams struct {
	// Only return Sim resources with this status.
	Status *string `json:"Status,omitempty"`
	// Only return Sim resources with this ICCID. This will return a list with a maximum size of 1.
	Iccid *string `json:"Iccid,omitempty"`
	// The SID or unique name of a [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource). Only return Sim resources assigned to this RatePlan resource.
	RatePlan *string `json:"RatePlan,omitempty"`
	// Deprecated.
	EId *string `json:"EId,omitempty"`
	// Only return Sim resources with this registration code. This will return a list with a maximum size of 1.
	SimRegistrationCode *string `json:"SimRegistrationCode,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListSimParams) SetStatus(Status string) *ListSimParams {
	params.Status = &Status
	return params
}
func (params *ListSimParams) SetIccid(Iccid string) *ListSimParams {
	params.Iccid = &Iccid
	return params
}
func (params *ListSimParams) SetRatePlan(RatePlan string) *ListSimParams {
	params.RatePlan = &RatePlan
	return params
}
func (params *ListSimParams) SetEId(EId string) *ListSimParams {
	params.EId = &EId
	return params
}
func (params *ListSimParams) SetSimRegistrationCode(SimRegistrationCode string) *ListSimParams {
	params.SimRegistrationCode = &SimRegistrationCode
	return params
}
func (params *ListSimParams) SetPageSize(PageSize int32) *ListSimParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of Sim resources on your Account.
func (c *DefaultApiService) ListSim(params *ListSimParams) (*ListSimResponse, error) {
	path := "/v1/Sims"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Iccid != nil {
		data.Set("Iccid", *params.Iccid)
	}
	if params != nil && params.RatePlan != nil {
		data.Set("RatePlan", *params.RatePlan)
	}
	if params != nil && params.EId != nil {
		data.Set("EId", *params.EId)
	}
	if params != nil && params.SimRegistrationCode != nil {
		data.Set("SimRegistrationCode", *params.SimRegistrationCode)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSimResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUsageRecord'
type ListUsageRecordParams struct {
	// Only include usage that occurred on or before this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is the current time.
	End *time.Time `json:"End,omitempty"`
	// Only include usage that has occurred on or after this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is one month before the `end` parameter value.
	Start *time.Time `json:"Start,omitempty"`
	// How to summarize the usage by time. Can be: `daily`, `hourly`, or `all`. The default is `all`. A value of `all` returns one Usage Record that describes the usage for the entire period.
	Granularity *string `json:"Granularity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListUsageRecordParams) SetEnd(End time.Time) *ListUsageRecordParams {
	params.End = &End
	return params
}
func (params *ListUsageRecordParams) SetStart(Start time.Time) *ListUsageRecordParams {
	params.Start = &Start
	return params
}
func (params *ListUsageRecordParams) SetGranularity(Granularity string) *ListUsageRecordParams {
	params.Granularity = &Granularity
	return params
}
func (params *ListUsageRecordParams) SetPageSize(PageSize int32) *ListUsageRecordParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListUsageRecord(SimSid string, params *ListUsageRecordParams) (*ListUsageRecordResponse, error) {
	path := "/v1/Sims/{SimSid}/UsageRecords"
	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.End != nil {
		data.Set("End", fmt.Sprint((*params.End).Format(time.RFC3339)))
	}
	if params != nil && params.Start != nil {
		data.Set("Start", fmt.Sprint((*params.Start).Format(time.RFC3339)))
	}
	if params != nil && params.Granularity != nil {
		data.Set("Granularity", *params.Granularity)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateRatePlan'
type UpdateRatePlanParams struct {
	// A descriptive string that you create to describe the resource. It does not have to be unique.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateRatePlanParams) SetFriendlyName(FriendlyName string) *UpdateRatePlanParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateRatePlanParams) SetUniqueName(UniqueName string) *UpdateRatePlanParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *DefaultApiService) UpdateRatePlan(Sid string, params *UpdateRatePlanParams) (*WirelessV1RatePlan, error) {
	path := "/v1/RatePlans/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1RatePlan{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateSim'
type UpdateSimParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a [Subaccount](https://www.twilio.com/docs/iam/api/subaccounts) of the requesting Account. Only valid when the Sim resource's status is `new`. For more information, see the [Move SIMs between Subaccounts documentation](https://www.twilio.com/docs/wireless/api/sim-resource#move-sims-between-subaccounts).
	AccountSid *string `json:"AccountSid,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be: `POST` or `GET`. The default is `POST`.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_url` when the SIM has finished updating. When the SIM transitions from `new` to `ready` or from any status to `deactivated`, we call this URL when the status changes to an intermediate status (`ready` or `deactivated`) and again when the status changes to its final status (`active` or `canceled`).
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The HTTP method we should use to call `commands_callback_url`. Can be: `POST` or `GET`. The default is `POST`.
	CommandsCallbackMethod *string `json:"CommandsCallbackMethod,omitempty"`
	// The URL we should call using the `commands_callback_method` when the SIM sends a [Command](https://www.twilio.com/docs/wireless/api/command-resource). Your server should respond with an HTTP status code in the 200 range; any response body is ignored.
	CommandsCallbackUrl *string `json:"CommandsCallbackUrl,omitempty"`
	// A descriptive string that you create to describe the Sim resource. It does not need to be unique.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID or unique name of the [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource) to which the Sim resource should be assigned.
	RatePlan *string `json:"RatePlan,omitempty"`
	// Initiate a connectivity reset on the SIM. Set to `resetting` to initiate a connectivity reset on the SIM. No other value is valid.
	ResetStatus *string `json:"ResetStatus,omitempty"`
	// The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`. Default is `POST`.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// The URL we should call using the `sms_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `sms_url`.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`. Default is `POST`.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// The URL we should call using the `sms_method` when the SIM-connected device sends an SMS message that is not a [Command](https://www.twilio.com/docs/wireless/api/command-resource).
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The new status of the Sim resource. Can be: `ready`, `active`, `suspended`, or `deactivated`.
	Status *string `json:"Status,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL we should call using the `voice_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `voice_url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use when we call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The URL we should call using the `voice_method` when the SIM-connected device makes a voice call.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
}

func (params *UpdateSimParams) SetAccountSid(AccountSid string) *UpdateSimParams {
	params.AccountSid = &AccountSid
	return params
}
func (params *UpdateSimParams) SetCallbackMethod(CallbackMethod string) *UpdateSimParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *UpdateSimParams) SetCallbackUrl(CallbackUrl string) *UpdateSimParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *UpdateSimParams) SetCommandsCallbackMethod(CommandsCallbackMethod string) *UpdateSimParams {
	params.CommandsCallbackMethod = &CommandsCallbackMethod
	return params
}
func (params *UpdateSimParams) SetCommandsCallbackUrl(CommandsCallbackUrl string) *UpdateSimParams {
	params.CommandsCallbackUrl = &CommandsCallbackUrl
	return params
}
func (params *UpdateSimParams) SetFriendlyName(FriendlyName string) *UpdateSimParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateSimParams) SetRatePlan(RatePlan string) *UpdateSimParams {
	params.RatePlan = &RatePlan
	return params
}
func (params *UpdateSimParams) SetResetStatus(ResetStatus string) *UpdateSimParams {
	params.ResetStatus = &ResetStatus
	return params
}
func (params *UpdateSimParams) SetSmsFallbackMethod(SmsFallbackMethod string) *UpdateSimParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}
func (params *UpdateSimParams) SetSmsFallbackUrl(SmsFallbackUrl string) *UpdateSimParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *UpdateSimParams) SetSmsMethod(SmsMethod string) *UpdateSimParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *UpdateSimParams) SetSmsUrl(SmsUrl string) *UpdateSimParams {
	params.SmsUrl = &SmsUrl
	return params
}
func (params *UpdateSimParams) SetStatus(Status string) *UpdateSimParams {
	params.Status = &Status
	return params
}
func (params *UpdateSimParams) SetUniqueName(UniqueName string) *UpdateSimParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *UpdateSimParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *UpdateSimParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *UpdateSimParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *UpdateSimParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *UpdateSimParams) SetVoiceMethod(VoiceMethod string) *UpdateSimParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *UpdateSimParams) SetVoiceUrl(VoiceUrl string) *UpdateSimParams {
	params.VoiceUrl = &VoiceUrl
	return params
}

// Updates the given properties of a Sim resource on your Account.
func (c *DefaultApiService) UpdateSim(Sid string, params *UpdateSimParams) (*WirelessV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.CommandsCallbackMethod != nil {
		data.Set("CommandsCallbackMethod", *params.CommandsCallbackMethod)
	}
	if params != nil && params.CommandsCallbackUrl != nil {
		data.Set("CommandsCallbackUrl", *params.CommandsCallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.RatePlan != nil {
		data.Set("RatePlan", *params.RatePlan)
	}
	if params != nil && params.ResetStatus != nil {
		data.Set("ResetStatus", *params.ResetStatus)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
