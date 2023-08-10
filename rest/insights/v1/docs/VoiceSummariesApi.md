# VoiceSummariesApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCallSummaries**](VoiceSummariesApi.md#ListCallSummaries) | **Get** /v1/Voice/Summaries | 



## ListCallSummaries

> []InsightsV1CallSummaries ListCallSummaries(ctx, optional)



Get a list of Call Summaries.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCallSummariesParams struct


Name | Type | Description
------------- | ------------- | -------------
**From** | **string** | A calling party. Could be an E.164 number, a SIP URI, or a Twilio Client registered name.
**To** | **string** | A called party. Could be an E.164 number, a SIP URI, or a Twilio Client registered name.
**FromCarrier** | **string** | An origination carrier.
**ToCarrier** | **string** | A destination carrier.
**FromCountryCode** | **string** | A source country code based on phone number in From.
**ToCountryCode** | **string** | A destination country code. Based on phone number in To.
**Branded** | **bool** | A boolean flag indicating whether or not the calls were branded using Twilio Branded Calls.
**VerifiedCaller** | **bool** | A boolean flag indicating whether or not the caller was verified using SHAKEN/STIR.
**HasTag** | **bool** | A boolean flag indicating the presence of one or more [Voice Insights Call Tags](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-tags).
**StartTime** | **string** | A Start time of the calls. xm (x minutes), xh (x hours), xd (x days), 1w, 30m, 3d, 4w or datetime-ISO. Defaults to 4h.
**EndTime** | **string** | An End Time of the calls. xm (x minutes), xh (x hours), xd (x days), 1w, 30m, 3d, 4w or datetime-ISO. Defaults to 0m.
**CallType** | **string** | A Call Type of the calls. One of `carrier`, `sip`, `trunking` or `client`.
**CallState** | **string** | A Call State of the calls. One of `ringing`, `completed`, `busy`, `fail`, `noanswer`, `canceled`, `answered`, `undialed`.
**Direction** | **string** | A Direction of the calls. One of `outbound_api`, `outbound_dial`, `inbound`, `trunking_originating`, `trunking_terminating`.
**ProcessingState** | **string** | A Processing State of the Call Summaries. One of `completed`, `partial` or `all`.
**SortBy** | **string** | A Sort By criterion for the returned list of Call Summaries. One of `start_time` or `end_time`.
**Subaccount** | **string** | A unique SID identifier of a Subaccount.
**AbnormalSession** | **bool** | A boolean flag indicating an abnormal session where the last SIP response was not 200 OK.
**AnsweredBy** | **string** | An Answered By value for the calls based on `Answering Machine Detection (AMD)`. One of `unknown`, `machine_start`, `machine_end_beep`, `machine_end_silence`, `machine_end_other`, `human` or `fax`.
**AnsweredByAnnotation** | **string** | Either machine or human.
**ConnectivityIssueAnnotation** | **string** | A Connectivity Issue with the calls. One of `no_connectivity_issue`, `invalid_number`, `caller_id`, `dropped_call`, or `number_reachability`.
**QualityIssueAnnotation** | **string** | A subjective Quality Issue with the calls. One of `no_quality_issue`, `low_volume`, `choppy_robotic`, `echo`, `dtmf`, `latency`, `owa`, `static_noise`.
**SpamAnnotation** | **bool** | A boolean flag indicating spam calls.
**CallScoreAnnotation** | **string** | A Call Score of the calls. Use a range of 1-5 to indicate the call experience score, with the following mapping as a reference for the rated call [5: Excellent, 4: Good, 3 : Fair, 2 : Poor, 1: Bad].
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]InsightsV1CallSummaries**](InsightsV1CallSummaries.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

