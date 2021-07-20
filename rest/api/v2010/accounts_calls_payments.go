/*
 * Twilio - Api
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

// Optional parameters for the method 'CreatePayments'
type CreatePaymentsParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Type of bank account if payment source is ACH. One of `consumer-checking`, `consumer-savings`, or `commercial-checking`. The default value is `consumer-checking`.
	BankAccountType *string `json:"BankAccountType,omitempty"`
	// A positive decimal value less than 1,000,000 to charge against the credit card or bank account. Default currency can be overwritten with `currency` field. Leave blank or set to 0 to tokenize.
	ChargeAmount *float32 `json:"ChargeAmount,omitempty"`
	// The currency of the `charge_amount`, formatted as [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format. The default value is `USD` and all values allowed from the <Pay> Connector are accepted.
	Currency *string `json:"Currency,omitempty"`
	// The description can be used to provide more details regarding the transaction. This information is submitted along with the payment details to the Payment Connector which are then posted on the transactions.
	Description *string `json:"Description,omitempty"`
	// A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated.
	IdempotencyKey *string `json:"IdempotencyKey,omitempty"`
	// A list of inputs that should be accepted. Currently only `dtmf` is supported. All digits captured during a pay session are redacted from the logs.
	Input *string `json:"Input,omitempty"`
	// A positive integer that is used to validate the length of the `PostalCode` inputted by the user. User must enter this many digits.
	MinPostalCodeLength *int `json:"MinPostalCodeLength,omitempty"`
	// A single level JSON string that is required when accepting certain information specific only to ACH payments. The information that has to be included here depends on the <Pay> Connector. [Read more](https://www.twilio.com/console/voice/pay-connectors).
	Parameter *map[string]interface{} `json:"Parameter,omitempty"`
	// This is the unique name corresponding to the Payment Gateway Connector installed in the Twilio Add-ons. Learn more about [<Pay> Connectors](https://www.twilio.com/console/voice/pay-connectors). The default value is `Default`.
	PaymentConnector *string `json:"PaymentConnector,omitempty"`
	// Type of payment being captured. One of `credit-card` or `ach-debit`. The default value is `credit-card`.
	PaymentMethod *string `json:"PaymentMethod,omitempty"`
	// Indicates whether the credit card postal code (zip code) is a required piece of payment information that must be provided by the caller. The default is `true`.
	PostalCode *bool `json:"PostalCode,omitempty"`
	// Indicates whether the credit card security code is a required piece of payment information that must be provided by the caller. The default is `true`.
	SecurityCode *bool `json:"SecurityCode,omitempty"`
	// Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [expected StatusCallback values](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback)
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The number of seconds that <Pay> should wait for the caller to press a digit between each subsequent digit, after the first one, before moving on to validate the digits captured. The default is `5`, maximum is `600`.
	Timeout *int `json:"Timeout,omitempty"`
	// Indicates whether the payment method should be tokenized as a `one-time` or `reusable` token. The default value is `reusable`. Do not enter a charge amount when tokenizing. If a charge amount is entered, the payment method will be charged and not tokenized.
	TokenType *string `json:"TokenType,omitempty"`
	// Credit card types separated by space that Pay should accept. The default value is `visa mastercard amex`
	ValidCardTypes *string `json:"ValidCardTypes,omitempty"`
}

func (params *CreatePaymentsParams) SetPathAccountSid(PathAccountSid string) *CreatePaymentsParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreatePaymentsParams) SetBankAccountType(BankAccountType string) *CreatePaymentsParams {
	params.BankAccountType = &BankAccountType
	return params
}
func (params *CreatePaymentsParams) SetChargeAmount(ChargeAmount float32) *CreatePaymentsParams {
	params.ChargeAmount = &ChargeAmount
	return params
}
func (params *CreatePaymentsParams) SetCurrency(Currency string) *CreatePaymentsParams {
	params.Currency = &Currency
	return params
}
func (params *CreatePaymentsParams) SetDescription(Description string) *CreatePaymentsParams {
	params.Description = &Description
	return params
}
func (params *CreatePaymentsParams) SetIdempotencyKey(IdempotencyKey string) *CreatePaymentsParams {
	params.IdempotencyKey = &IdempotencyKey
	return params
}
func (params *CreatePaymentsParams) SetInput(Input string) *CreatePaymentsParams {
	params.Input = &Input
	return params
}
func (params *CreatePaymentsParams) SetMinPostalCodeLength(MinPostalCodeLength int) *CreatePaymentsParams {
	params.MinPostalCodeLength = &MinPostalCodeLength
	return params
}
func (params *CreatePaymentsParams) SetParameter(Parameter map[string]interface{}) *CreatePaymentsParams {
	params.Parameter = &Parameter
	return params
}
func (params *CreatePaymentsParams) SetPaymentConnector(PaymentConnector string) *CreatePaymentsParams {
	params.PaymentConnector = &PaymentConnector
	return params
}
func (params *CreatePaymentsParams) SetPaymentMethod(PaymentMethod string) *CreatePaymentsParams {
	params.PaymentMethod = &PaymentMethod
	return params
}
func (params *CreatePaymentsParams) SetPostalCode(PostalCode bool) *CreatePaymentsParams {
	params.PostalCode = &PostalCode
	return params
}
func (params *CreatePaymentsParams) SetSecurityCode(SecurityCode bool) *CreatePaymentsParams {
	params.SecurityCode = &SecurityCode
	return params
}
func (params *CreatePaymentsParams) SetStatusCallback(StatusCallback string) *CreatePaymentsParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreatePaymentsParams) SetTimeout(Timeout int) *CreatePaymentsParams {
	params.Timeout = &Timeout
	return params
}
func (params *CreatePaymentsParams) SetTokenType(TokenType string) *CreatePaymentsParams {
	params.TokenType = &TokenType
	return params
}
func (params *CreatePaymentsParams) SetValidCardTypes(ValidCardTypes string) *CreatePaymentsParams {
	params.ValidCardTypes = &ValidCardTypes
	return params
}

// create an instance of payments. This will start a new payments session
func (c *ApiService) CreatePayments(CallSid string, params *CreatePaymentsParams) (*ApiV2010AccountCallPayments, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Payments.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	if params != nil && params.BankAccountType != nil {
		data.Set("BankAccountType", *params.BankAccountType)
	}
	if params != nil && params.ChargeAmount != nil {
		data.Set("ChargeAmount", fmt.Sprint(*params.ChargeAmount))
	}
	if params != nil && params.Currency != nil {
		data.Set("Currency", *params.Currency)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.IdempotencyKey != nil {
		data.Set("IdempotencyKey", *params.IdempotencyKey)
	}
	if params != nil && params.Input != nil {
		data.Set("Input", *params.Input)
	}
	if params != nil && params.MinPostalCodeLength != nil {
		data.Set("MinPostalCodeLength", fmt.Sprint(*params.MinPostalCodeLength))
	}
	if params != nil && params.Parameter != nil {
		v, err := json.Marshal(params.Parameter)

		if err != nil {
			return nil, err
		}

		data.Set("Parameter", string(v))
	}
	if params != nil && params.PaymentConnector != nil {
		data.Set("PaymentConnector", *params.PaymentConnector)
	}
	if params != nil && params.PaymentMethod != nil {
		data.Set("PaymentMethod", *params.PaymentMethod)
	}
	if params != nil && params.PostalCode != nil {
		data.Set("PostalCode", fmt.Sprint(*params.PostalCode))
	}
	if params != nil && params.SecurityCode != nil {
		data.Set("SecurityCode", fmt.Sprint(*params.SecurityCode))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.Timeout != nil {
		data.Set("Timeout", fmt.Sprint(*params.Timeout))
	}
	if params != nil && params.TokenType != nil {
		data.Set("TokenType", *params.TokenType)
	}
	if params != nil && params.ValidCardTypes != nil {
		data.Set("ValidCardTypes", *params.ValidCardTypes)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountCallPayments{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdatePayments'
type UpdatePaymentsParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will update the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The piece of payment information that you wish the caller to enter. Must be one of `payment-card-number`, `expiration-date`, `security-code`, `postal-code`, `bank-routing-number`, or `bank-account-number`.
	Capture *string `json:"Capture,omitempty"`
	// A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated.
	IdempotencyKey *string `json:"IdempotencyKey,omitempty"`
	// Indicates whether the current payment session should be cancelled or completed. When `cancel` the payment session is cancelled. When `complete`, Twilio sends the payment information to the selected <Pay> connector for processing.
	Status *string `json:"Status,omitempty"`
	// Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [Update](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-update) and [Complete/Cancel](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-cancelcomplete) POST requests.
	StatusCallback *string `json:"StatusCallback,omitempty"`
}

func (params *UpdatePaymentsParams) SetPathAccountSid(PathAccountSid string) *UpdatePaymentsParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdatePaymentsParams) SetCapture(Capture string) *UpdatePaymentsParams {
	params.Capture = &Capture
	return params
}
func (params *UpdatePaymentsParams) SetIdempotencyKey(IdempotencyKey string) *UpdatePaymentsParams {
	params.IdempotencyKey = &IdempotencyKey
	return params
}
func (params *UpdatePaymentsParams) SetStatus(Status string) *UpdatePaymentsParams {
	params.Status = &Status
	return params
}
func (params *UpdatePaymentsParams) SetStatusCallback(StatusCallback string) *UpdatePaymentsParams {
	params.StatusCallback = &StatusCallback
	return params
}

// update an instance of payments with different phases of payment flows.
func (c *ApiService) UpdatePayments(CallSid string, Sid string, params *UpdatePaymentsParams) (*ApiV2010AccountCallPayments, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Payments/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	if params != nil && params.Capture != nil {
		data.Set("Capture", *params.Capture)
	}
	if params != nil && params.IdempotencyKey != nil {
		data.Set("IdempotencyKey", *params.IdempotencyKey)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountCallPayments{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
