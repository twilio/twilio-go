/*
    * Twilio - Numbers
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	twilio "github.com/twilio/twilio-go/client"
	"net/url"
    "strings"
)

type DefaultApiService struct {
	baseURL string
	client  *twilio.Client
}

func NewDefaultApiService(client *twilio.Client) *DefaultApiService {
	return &DefaultApiService {
		client: client,
		baseURL: "https://numbers.twilio.com",
	}
}

// CreateBundleParams Optional parameters for the method 'CreateBundle'
type CreateBundleParams struct {
	Email *string `json:"Email,omitempty"`
	EndUserType *string `json:"EndUserType,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
	IsoCountry *string `json:"IsoCountry,omitempty"`
	NumberType *string `json:"NumberType,omitempty"`
	RegulationSid *string `json:"RegulationSid,omitempty"`
	StatusCallback *string `json:"StatusCallback,omitempty"`
}

/*
* CreateBundle Method for CreateBundle
* Create a new Bundle.
* @param optional nil or *CreateBundleParams - Optional Parameters:
* @param "Email" (string) - The email address that will receive updates when the Bundle resource changes status.
* @param "EndUserType" (string) - The type of End User of the Bundle resource.
* @param "FriendlyName" (string) - The string that you assigned to describe the resource.
* @param "IsoCountry" (string) - The ISO country code of the Bundle's phone number country ownership request.
* @param "NumberType" (string) - The type of phone number of the Bundle's ownership request.
* @param "RegulationSid" (string) - The unique string of a regulation that is associated to the Bundle resource.
* @param "StatusCallback" (string) - The URL we call to inform your application of status changes.
* @return NumbersV2RegulatoryComplianceBundle
*/
func (c *DefaultApiService) CreateBundle(params *CreateBundleParams) (*NumbersV2RegulatoryComplianceBundle, error) {
    path := "/v2/RegulatoryCompliance/Bundles"

    data := url.Values{}
    headers := 0

    if params != nil && params.Email != nil {
	data.Set("Email", *params.Email) 
}
    if params != nil && params.EndUserType != nil {
	data.Set("EndUserType", *params.EndUserType) 
}
    if params != nil && params.FriendlyName != nil {
	data.Set("FriendlyName", *params.FriendlyName) 
}
    if params != nil && params.IsoCountry != nil {
	data.Set("IsoCountry", *params.IsoCountry) 
}
    if params != nil && params.NumberType != nil {
	data.Set("NumberType", *params.NumberType) 
}
    if params != nil && params.RegulationSid != nil {
	data.Set("RegulationSid", *params.RegulationSid) 
}
    if params != nil && params.StatusCallback != nil {
	data.Set("StatusCallback", *params.StatusCallback) 
}


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceBundle{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// CreateEndUserParams Optional parameters for the method 'CreateEndUser'
type CreateEndUserParams struct {
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
	Type *string `json:"Type,omitempty"`
}

/*
* CreateEndUser Method for CreateEndUser
* Create a new End User.
* @param optional nil or *CreateEndUserParams - Optional Parameters:
* @param "Attributes" (map[string]interface{}) - The set of parameters that are the attributes of the End User resource which are derived End User Types.
* @param "FriendlyName" (string) - The string that you assigned to describe the resource.
* @param "Type" (string) - The type of end user of the Bundle resource - can be `individual` or `business`.
* @return NumbersV2RegulatoryComplianceEndUser
*/
func (c *DefaultApiService) CreateEndUser(params *CreateEndUserParams) (*NumbersV2RegulatoryComplianceEndUser, error) {
    path := "/v2/RegulatoryCompliance/EndUsers"

    data := url.Values{}
    headers := 0

    if params != nil && params.Attributes != nil {
	v, err := json.Marshal(params.Attributes)

	if err != nil {
	    return nil, err
	}

	data.Set("Attributes", fmt.Sprint(v))
}
    if params != nil && params.FriendlyName != nil {
	data.Set("FriendlyName", *params.FriendlyName) 
}
    if params != nil && params.Type != nil {
	data.Set("Type", *params.Type) 
}


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceEndUser{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
* CreateEvaluation Method for CreateEvaluation
* @param BundleSid
* @return NumbersV2RegulatoryComplianceBundleEvaluation
*/
func (c *DefaultApiService) CreateEvaluation(BundleSid string, ) (*NumbersV2RegulatoryComplianceBundleEvaluation, error) {
    path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceBundleEvaluation{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// CreateItemAssignmentParams Optional parameters for the method 'CreateItemAssignment'
type CreateItemAssignmentParams struct {
	ObjectSid *string `json:"ObjectSid,omitempty"`
}

/*
* CreateItemAssignment Method for CreateItemAssignment
* Create a new Assigned Item.
* @param BundleSid The unique string that we created to identify the Bundle resource.
* @param optional nil or *CreateItemAssignmentParams - Optional Parameters:
* @param "ObjectSid" (string) - The SID of an object bag that holds information of the different items.
* @return NumbersV2RegulatoryComplianceBundleItemAssignment
*/
func (c *DefaultApiService) CreateItemAssignment(BundleSid string, params *CreateItemAssignmentParams) (*NumbersV2RegulatoryComplianceBundleItemAssignment, error) {
    path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.ObjectSid != nil {
	data.Set("ObjectSid", *params.ObjectSid) 
}


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceBundleItemAssignment{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// CreateSupportingDocumentParams Optional parameters for the method 'CreateSupportingDocument'
type CreateSupportingDocumentParams struct {
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
	Type *string `json:"Type,omitempty"`
}

/*
* CreateSupportingDocument Method for CreateSupportingDocument
* Create a new Supporting Document.
* @param optional nil or *CreateSupportingDocumentParams - Optional Parameters:
* @param "Attributes" (map[string]interface{}) - The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.
* @param "FriendlyName" (string) - The string that you assigned to describe the resource.
* @param "Type" (string) - The type of the Supporting Document.
* @return NumbersV2RegulatoryComplianceSupportingDocument
*/
func (c *DefaultApiService) CreateSupportingDocument(params *CreateSupportingDocumentParams) (*NumbersV2RegulatoryComplianceSupportingDocument, error) {
    path := "/v2/RegulatoryCompliance/SupportingDocuments"

    data := url.Values{}
    headers := 0

    if params != nil && params.Attributes != nil {
	v, err := json.Marshal(params.Attributes)

	if err != nil {
	    return nil, err
	}

	data.Set("Attributes", fmt.Sprint(v))
}
    if params != nil && params.FriendlyName != nil {
	data.Set("FriendlyName", *params.FriendlyName) 
}
    if params != nil && params.Type != nil {
	data.Set("Type", *params.Type) 
}


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceSupportingDocument{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
* DeleteBundle Method for DeleteBundle
* Delete a specific Bundle.
* @param Sid The unique string that we created to identify the Bundle resource.
*/
func (c *DefaultApiService) DeleteBundle(Sid string, ) (error) {
    path := "/v2/RegulatoryCompliance/Bundles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

/*
* DeleteEndUser Method for DeleteEndUser
* Delete a specific End User.
* @param Sid The unique string created by Twilio to identify the End User resource.
*/
func (c *DefaultApiService) DeleteEndUser(Sid string, ) (error) {
    path := "/v2/RegulatoryCompliance/EndUsers/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

/*
* DeleteItemAssignment Method for DeleteItemAssignment
* Remove an Assignment Item Instance.
* @param BundleSid The unique string that we created to identify the Bundle resource.
* @param Sid The unique string that we created to identify the Identity resource.
*/
func (c *DefaultApiService) DeleteItemAssignment(BundleSid string, Sid string, ) (error) {
    path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

/*
* DeleteSupportingDocument Method for DeleteSupportingDocument
* Delete a specific Supporting Document.
* @param Sid The unique string created by Twilio to identify the Supporting Document resource.
*/
func (c *DefaultApiService) DeleteSupportingDocument(Sid string, ) (error) {
    path := "/v2/RegulatoryCompliance/SupportingDocuments/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

/*
* FetchBundle Method for FetchBundle
* Fetch a specific Bundle instance.
* @param Sid The unique string that we created to identify the Bundle resource.
* @return NumbersV2RegulatoryComplianceBundle
*/
func (c *DefaultApiService) FetchBundle(Sid string, ) (*NumbersV2RegulatoryComplianceBundle, error) {
    path := "/v2/RegulatoryCompliance/Bundles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceBundle{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
* FetchEndUser Method for FetchEndUser
* Fetch specific End User Instance.
* @param Sid The unique string created by Twilio to identify the End User resource.
* @return NumbersV2RegulatoryComplianceEndUser
*/
func (c *DefaultApiService) FetchEndUser(Sid string, ) (*NumbersV2RegulatoryComplianceEndUser, error) {
    path := "/v2/RegulatoryCompliance/EndUsers/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceEndUser{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
* FetchEndUserType Method for FetchEndUserType
* Fetch a specific End-User Type Instance.
* @param Sid The unique string that identifies the End-User Type resource.
* @return NumbersV2RegulatoryComplianceEndUserType
*/
func (c *DefaultApiService) FetchEndUserType(Sid string, ) (*NumbersV2RegulatoryComplianceEndUserType, error) {
    path := "/v2/RegulatoryCompliance/EndUserTypes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceEndUserType{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
* FetchEvaluation Method for FetchEvaluation
* Fetch specific Evaluation Instance.
* @param BundleSid The unique string that we created to identify the Bundle resource.
* @param Sid The unique string that identifies the Evaluation resource.
* @return NumbersV2RegulatoryComplianceBundleEvaluation
*/
func (c *DefaultApiService) FetchEvaluation(BundleSid string, Sid string, ) (*NumbersV2RegulatoryComplianceBundleEvaluation, error) {
    path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceBundleEvaluation{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
* FetchItemAssignment Method for FetchItemAssignment
* Fetch specific Assigned Item Instance.
* @param BundleSid The unique string that we created to identify the Bundle resource.
* @param Sid The unique string that we created to identify the Identity resource.
* @return NumbersV2RegulatoryComplianceBundleItemAssignment
*/
func (c *DefaultApiService) FetchItemAssignment(BundleSid string, Sid string, ) (*NumbersV2RegulatoryComplianceBundleItemAssignment, error) {
    path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceBundleItemAssignment{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
* FetchRegulation Method for FetchRegulation
* Fetch specific Regulation Instance.
* @param Sid The unique string that identifies the Regulation resource.
* @return NumbersV2RegulatoryComplianceRegulation
*/
func (c *DefaultApiService) FetchRegulation(Sid string, ) (*NumbersV2RegulatoryComplianceRegulation, error) {
    path := "/v2/RegulatoryCompliance/Regulations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceRegulation{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
* FetchSupportingDocument Method for FetchSupportingDocument
* Fetch specific Supporting Document Instance.
* @param Sid The unique string created by Twilio to identify the Supporting Document resource.
* @return NumbersV2RegulatoryComplianceSupportingDocument
*/
func (c *DefaultApiService) FetchSupportingDocument(Sid string, ) (*NumbersV2RegulatoryComplianceSupportingDocument, error) {
    path := "/v2/RegulatoryCompliance/SupportingDocuments/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceSupportingDocument{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
* FetchSupportingDocumentType Method for FetchSupportingDocumentType
* Fetch a specific Supporting Document Type Instance.
* @param Sid The unique string that identifies the Supporting Document Type resource.
* @return NumbersV2RegulatoryComplianceSupportingDocumentType
*/
func (c *DefaultApiService) FetchSupportingDocumentType(Sid string, ) (*NumbersV2RegulatoryComplianceSupportingDocumentType, error) {
    path := "/v2/RegulatoryCompliance/SupportingDocumentTypes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceSupportingDocumentType{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListBundleParams Optional parameters for the method 'ListBundle'
type ListBundleParams struct {
	Status *string `json:"Status,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
	RegulationSid *string `json:"RegulationSid,omitempty"`
	IsoCountry *string `json:"IsoCountry,omitempty"`
	NumberType *string `json:"NumberType,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListBundle Method for ListBundle
* Retrieve a list of all Bundles for an account.
* @param optional nil or *ListBundleParams - Optional Parameters:
* @param "Status" (string) - The verification status of the Bundle resource.
* @param "FriendlyName" (string) - The string that you assigned to describe the resource.
* @param "RegulationSid" (string) - The unique string of a regulation that is associated to the Bundle resource.
* @param "IsoCountry" (string) - The ISO country code of the Bundle's phone number country ownership request.
* @param "NumberType" (string) - The type of phone number of the Bundle's ownership request.
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListBundleResponse
*/
func (c *DefaultApiService) ListBundle(params *ListBundleParams) (*ListBundleResponse, error) {
    path := "/v2/RegulatoryCompliance/Bundles"

    data := url.Values{}
    headers := 0

    if params != nil && params.Status != nil {
	data.Set("Status", *params.Status) 
}
    if params != nil && params.FriendlyName != nil {
	data.Set("FriendlyName", *params.FriendlyName) 
}
    if params != nil && params.RegulationSid != nil {
	data.Set("RegulationSid", *params.RegulationSid) 
}
    if params != nil && params.IsoCountry != nil {
	data.Set("IsoCountry", *params.IsoCountry) 
}
    if params != nil && params.NumberType != nil {
	data.Set("NumberType", *params.NumberType) 
}
    if params != nil && params.PageSize != nil {
	data.Set("PageSize", fmt.Sprint(*params.PageSize)) 
}


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListBundleResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListEndUserParams Optional parameters for the method 'ListEndUser'
type ListEndUserParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListEndUser Method for ListEndUser
* Retrieve a list of all End User for an account.
* @param optional nil or *ListEndUserParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListEndUserResponse
*/
func (c *DefaultApiService) ListEndUser(params *ListEndUserParams) (*ListEndUserResponse, error) {
    path := "/v2/RegulatoryCompliance/EndUsers"

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
	data.Set("PageSize", fmt.Sprint(*params.PageSize)) 
}


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListEndUserResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListEndUserTypeParams Optional parameters for the method 'ListEndUserType'
type ListEndUserTypeParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListEndUserType Method for ListEndUserType
* Retrieve a list of all End-User Types.
* @param optional nil or *ListEndUserTypeParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListEndUserTypeResponse
*/
func (c *DefaultApiService) ListEndUserType(params *ListEndUserTypeParams) (*ListEndUserTypeResponse, error) {
    path := "/v2/RegulatoryCompliance/EndUserTypes"

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
	data.Set("PageSize", fmt.Sprint(*params.PageSize)) 
}


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListEndUserTypeResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListEvaluationParams Optional parameters for the method 'ListEvaluation'
type ListEvaluationParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListEvaluation Method for ListEvaluation
* Retrieve a list of Evaluations associated to the Bundle resource.
* @param BundleSid
* @param optional nil or *ListEvaluationParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListEvaluationResponse
*/
func (c *DefaultApiService) ListEvaluation(BundleSid string, params *ListEvaluationParams) (*ListEvaluationResponse, error) {
    path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
	data.Set("PageSize", fmt.Sprint(*params.PageSize)) 
}


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListEvaluationResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListItemAssignmentParams Optional parameters for the method 'ListItemAssignment'
type ListItemAssignmentParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListItemAssignment Method for ListItemAssignment
* Retrieve a list of all Assigned Items for an account.
* @param BundleSid The unique string that we created to identify the Bundle resource.
* @param optional nil or *ListItemAssignmentParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListItemAssignmentResponse
*/
func (c *DefaultApiService) ListItemAssignment(BundleSid string, params *ListItemAssignmentParams) (*ListItemAssignmentResponse, error) {
    path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
	data.Set("PageSize", fmt.Sprint(*params.PageSize)) 
}


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListItemAssignmentResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListRegulationParams Optional parameters for the method 'ListRegulation'
type ListRegulationParams struct {
	EndUserType *string `json:"EndUserType,omitempty"`
	IsoCountry *string `json:"IsoCountry,omitempty"`
	NumberType *string `json:"NumberType,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListRegulation Method for ListRegulation
* Retrieve a list of all Regulations.
* @param optional nil or *ListRegulationParams - Optional Parameters:
* @param "EndUserType" (string) - The type of End User the regulation requires - can be `individual` or `business`.
* @param "IsoCountry" (string) - The ISO country code of the phone number's country.
* @param "NumberType" (string) - The type of phone number that the regulatory requiremnt is restricting.
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListRegulationResponse
*/
func (c *DefaultApiService) ListRegulation(params *ListRegulationParams) (*ListRegulationResponse, error) {
    path := "/v2/RegulatoryCompliance/Regulations"

    data := url.Values{}
    headers := 0

    if params != nil && params.EndUserType != nil {
	data.Set("EndUserType", *params.EndUserType) 
}
    if params != nil && params.IsoCountry != nil {
	data.Set("IsoCountry", *params.IsoCountry) 
}
    if params != nil && params.NumberType != nil {
	data.Set("NumberType", *params.NumberType) 
}
    if params != nil && params.PageSize != nil {
	data.Set("PageSize", fmt.Sprint(*params.PageSize)) 
}


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListRegulationResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListSupportingDocumentParams Optional parameters for the method 'ListSupportingDocument'
type ListSupportingDocumentParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListSupportingDocument Method for ListSupportingDocument
* Retrieve a list of all Supporting Document for an account.
* @param optional nil or *ListSupportingDocumentParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListSupportingDocumentResponse
*/
func (c *DefaultApiService) ListSupportingDocument(params *ListSupportingDocumentParams) (*ListSupportingDocumentResponse, error) {
    path := "/v2/RegulatoryCompliance/SupportingDocuments"

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
	data.Set("PageSize", fmt.Sprint(*params.PageSize)) 
}


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListSupportingDocumentResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListSupportingDocumentTypeParams Optional parameters for the method 'ListSupportingDocumentType'
type ListSupportingDocumentTypeParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListSupportingDocumentType Method for ListSupportingDocumentType
* Retrieve a list of all Supporting Document Types.
* @param optional nil or *ListSupportingDocumentTypeParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListSupportingDocumentTypeResponse
*/
func (c *DefaultApiService) ListSupportingDocumentType(params *ListSupportingDocumentTypeParams) (*ListSupportingDocumentTypeResponse, error) {
    path := "/v2/RegulatoryCompliance/SupportingDocumentTypes"

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
	data.Set("PageSize", fmt.Sprint(*params.PageSize)) 
}


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListSupportingDocumentTypeResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// UpdateBundleParams Optional parameters for the method 'UpdateBundle'
type UpdateBundleParams struct {
	Email *string `json:"Email,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
	Status *string `json:"Status,omitempty"`
	StatusCallback *string `json:"StatusCallback,omitempty"`
}

/*
* UpdateBundle Method for UpdateBundle
* Updates a Bundle in an account.
* @param Sid The unique string that we created to identify the Bundle resource.
* @param optional nil or *UpdateBundleParams - Optional Parameters:
* @param "Email" (string) - The email address that will receive updates when the Bundle resource changes status.
* @param "FriendlyName" (string) - The string that you assigned to describe the resource.
* @param "Status" (string) - The verification status of the Bundle resource.
* @param "StatusCallback" (string) - The URL we call to inform your application of status changes.
* @return NumbersV2RegulatoryComplianceBundle
*/
func (c *DefaultApiService) UpdateBundle(Sid string, params *UpdateBundleParams) (*NumbersV2RegulatoryComplianceBundle, error) {
    path := "/v2/RegulatoryCompliance/Bundles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.Email != nil {
	data.Set("Email", *params.Email) 
}
    if params != nil && params.FriendlyName != nil {
	data.Set("FriendlyName", *params.FriendlyName) 
}
    if params != nil && params.Status != nil {
	data.Set("Status", *params.Status) 
}
    if params != nil && params.StatusCallback != nil {
	data.Set("StatusCallback", *params.StatusCallback) 
}


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceBundle{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// UpdateEndUserParams Optional parameters for the method 'UpdateEndUser'
type UpdateEndUserParams struct {
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

/*
* UpdateEndUser Method for UpdateEndUser
* Update an existing End User.
* @param Sid The unique string created by Twilio to identify the End User resource.
* @param optional nil or *UpdateEndUserParams - Optional Parameters:
* @param "Attributes" (map[string]interface{}) - The set of parameters that are the attributes of the End User resource which are derived End User Types.
* @param "FriendlyName" (string) - The string that you assigned to describe the resource.
* @return NumbersV2RegulatoryComplianceEndUser
*/
func (c *DefaultApiService) UpdateEndUser(Sid string, params *UpdateEndUserParams) (*NumbersV2RegulatoryComplianceEndUser, error) {
    path := "/v2/RegulatoryCompliance/EndUsers/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.Attributes != nil {
	v, err := json.Marshal(params.Attributes)

	if err != nil {
	    return nil, err
	}

	data.Set("Attributes", fmt.Sprint(v))
}
    if params != nil && params.FriendlyName != nil {
	data.Set("FriendlyName", *params.FriendlyName) 
}


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceEndUser{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// UpdateSupportingDocumentParams Optional parameters for the method 'UpdateSupportingDocument'
type UpdateSupportingDocumentParams struct {
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

/*
* UpdateSupportingDocument Method for UpdateSupportingDocument
* Update an existing Supporting Document.
* @param Sid The unique string created by Twilio to identify the Supporting Document resource.
* @param optional nil or *UpdateSupportingDocumentParams - Optional Parameters:
* @param "Attributes" (map[string]interface{}) - The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types.
* @param "FriendlyName" (string) - The string that you assigned to describe the resource.
* @return NumbersV2RegulatoryComplianceSupportingDocument
*/
func (c *DefaultApiService) UpdateSupportingDocument(Sid string, params *UpdateSupportingDocumentParams) (*NumbersV2RegulatoryComplianceSupportingDocument, error) {
    path := "/v2/RegulatoryCompliance/SupportingDocuments/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.Attributes != nil {
	v, err := json.Marshal(params.Attributes)

	if err != nil {
	    return nil, err
	}

	data.Set("Attributes", fmt.Sprint(v))
}
    if params != nil && params.FriendlyName != nil {
	data.Set("FriendlyName", *params.FriendlyName) 
}


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &NumbersV2RegulatoryComplianceSupportingDocument{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
