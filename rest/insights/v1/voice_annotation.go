/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
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

// Fetch a specific Annotation.
func (c *ApiService) FetchAnnotation(CallSid string) (*InsightsV1Annotation, error) {
	return c.FetchAnnotationWithCtx(context.TODO(), CallSid)
}

// Fetch a specific Annotation.
func (c *ApiService) FetchAnnotationWithCtx(ctx context.Context, CallSid string) (*InsightsV1Annotation, error) {
	path := "/v1/Voice/{CallSid}/Annotation"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1Annotation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateAnnotation'
type UpdateAnnotationParams struct {
	//
	AnsweredBy *string `json:"AnsweredBy,omitempty"`
	//
	ConnectivityIssue *string `json:"ConnectivityIssue,omitempty"`
	// Specify if the call had any subjective quality issues. Possible values, one or more of:  no_quality_issue, low_volume, choppy_robotic, echo, dtmf, latency, owa, static_noise. Use comma separated values to indicate multiple quality issues for the same call
	QualityIssues *string `json:"QualityIssues,omitempty"`
	// Specify if the call was a spam call. Use this to provide feedback on whether calls placed from your account were marked as spam, or if inbound calls received by your account were unwanted spam. Is of type Boolean: true, false. Use true if the call was a spam call.
	Spam *bool `json:"Spam,omitempty"`
	// Specify the call score. This is of type integer. Use a range of 1-5 to indicate the call experience score, with the following mapping as a reference for rating the call [5: Excellent, 4: Good, 3 : Fair, 2 : Poor, 1: Bad].
	CallScore *int `json:"CallScore,omitempty"`
	// Specify any comments pertaining to the call. This of type string with a max limit of 100 characters. Twilio does not treat this field as PII, so don’t put any PII in here.
	Comment *string `json:"Comment,omitempty"`
	// Associate this call with an incident or support ticket. This is of type string with a max limit of 100 characters. Twilio does not treat this field as PII, so don’t put any PII in here.
	Incident *string `json:"Incident,omitempty"`
}

func (params *UpdateAnnotationParams) SetAnsweredBy(AnsweredBy string) *UpdateAnnotationParams {
	params.AnsweredBy = &AnsweredBy
	return params
}
func (params *UpdateAnnotationParams) SetConnectivityIssue(ConnectivityIssue string) *UpdateAnnotationParams {
	params.ConnectivityIssue = &ConnectivityIssue
	return params
}
func (params *UpdateAnnotationParams) SetQualityIssues(QualityIssues string) *UpdateAnnotationParams {
	params.QualityIssues = &QualityIssues
	return params
}
func (params *UpdateAnnotationParams) SetSpam(Spam bool) *UpdateAnnotationParams {
	params.Spam = &Spam
	return params
}
func (params *UpdateAnnotationParams) SetCallScore(CallScore int) *UpdateAnnotationParams {
	params.CallScore = &CallScore
	return params
}
func (params *UpdateAnnotationParams) SetComment(Comment string) *UpdateAnnotationParams {
	params.Comment = &Comment
	return params
}
func (params *UpdateAnnotationParams) SetIncident(Incident string) *UpdateAnnotationParams {
	params.Incident = &Incident
	return params
}

// Create/Update the annotation for the call
func (c *ApiService) UpdateAnnotation(CallSid string, params *UpdateAnnotationParams) (*InsightsV1Annotation, error) {
	return c.UpdateAnnotationWithCtx(context.TODO(), CallSid, params)
}

// Create/Update the annotation for the call
func (c *ApiService) UpdateAnnotationWithCtx(ctx context.Context, CallSid string, params *UpdateAnnotationParams) (*InsightsV1Annotation, error) {
	path := "/v1/Voice/{CallSid}/Annotation"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AnsweredBy != nil {
		data.Set("AnsweredBy", *params.AnsweredBy)
	}
	if params != nil && params.ConnectivityIssue != nil {
		data.Set("ConnectivityIssue", *params.ConnectivityIssue)
	}
	if params != nil && params.QualityIssues != nil {
		data.Set("QualityIssues", *params.QualityIssues)
	}
	if params != nil && params.Spam != nil {
		data.Set("Spam", fmt.Sprint(*params.Spam))
	}
	if params != nil && params.CallScore != nil {
		data.Set("CallScore", fmt.Sprint(*params.CallScore))
	}
	if params != nil && params.Comment != nil {
		data.Set("Comment", *params.Comment)
	}
	if params != nil && params.Incident != nil {
		data.Set("Incident", *params.Incident)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1Annotation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
