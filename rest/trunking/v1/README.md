# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://trunking.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*TrunksApi* | [**CreateTrunk**](docs/TrunksApi.md#createtrunk) | **Post** /v1/Trunks | 
*TrunksApi* | [**DeleteTrunk**](docs/TrunksApi.md#deletetrunk) | **Delete** /v1/Trunks/{Sid} | 
*TrunksApi* | [**FetchTrunk**](docs/TrunksApi.md#fetchtrunk) | **Get** /v1/Trunks/{Sid} | 
*TrunksApi* | [**ListTrunk**](docs/TrunksApi.md#listtrunk) | **Get** /v1/Trunks | 
*TrunksApi* | [**UpdateTrunk**](docs/TrunksApi.md#updatetrunk) | **Post** /v1/Trunks/{Sid} | 
*TrunksCredentialListsApi* | [**CreateCredentialList**](docs/TrunksCredentialListsApi.md#createcredentiallist) | **Post** /v1/Trunks/{TrunkSid}/CredentialLists | 
*TrunksCredentialListsApi* | [**DeleteCredentialList**](docs/TrunksCredentialListsApi.md#deletecredentiallist) | **Delete** /v1/Trunks/{TrunkSid}/CredentialLists/{Sid} | 
*TrunksCredentialListsApi* | [**FetchCredentialList**](docs/TrunksCredentialListsApi.md#fetchcredentiallist) | **Get** /v1/Trunks/{TrunkSid}/CredentialLists/{Sid} | 
*TrunksCredentialListsApi* | [**ListCredentialList**](docs/TrunksCredentialListsApi.md#listcredentiallist) | **Get** /v1/Trunks/{TrunkSid}/CredentialLists | 
*TrunksIpAccessControlListsApi* | [**CreateIpAccessControlList**](docs/TrunksIpAccessControlListsApi.md#createipaccesscontrollist) | **Post** /v1/Trunks/{TrunkSid}/IpAccessControlLists | Associate an IP Access Control List with a Trunk
*TrunksIpAccessControlListsApi* | [**DeleteIpAccessControlList**](docs/TrunksIpAccessControlListsApi.md#deleteipaccesscontrollist) | **Delete** /v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid} | Remove an associated IP Access Control List from a Trunk
*TrunksIpAccessControlListsApi* | [**FetchIpAccessControlList**](docs/TrunksIpAccessControlListsApi.md#fetchipaccesscontrollist) | **Get** /v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid} | 
*TrunksIpAccessControlListsApi* | [**ListIpAccessControlList**](docs/TrunksIpAccessControlListsApi.md#listipaccesscontrollist) | **Get** /v1/Trunks/{TrunkSid}/IpAccessControlLists | List all IP Access Control Lists for a Trunk
*TrunksOriginationUrlsApi* | [**CreateOriginationUrl**](docs/TrunksOriginationUrlsApi.md#createoriginationurl) | **Post** /v1/Trunks/{TrunkSid}/OriginationUrls | 
*TrunksOriginationUrlsApi* | [**DeleteOriginationUrl**](docs/TrunksOriginationUrlsApi.md#deleteoriginationurl) | **Delete** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
*TrunksOriginationUrlsApi* | [**FetchOriginationUrl**](docs/TrunksOriginationUrlsApi.md#fetchoriginationurl) | **Get** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
*TrunksOriginationUrlsApi* | [**ListOriginationUrl**](docs/TrunksOriginationUrlsApi.md#listoriginationurl) | **Get** /v1/Trunks/{TrunkSid}/OriginationUrls | 
*TrunksOriginationUrlsApi* | [**UpdateOriginationUrl**](docs/TrunksOriginationUrlsApi.md#updateoriginationurl) | **Post** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
*TrunksPhoneNumbersApi* | [**CreatePhoneNumber**](docs/TrunksPhoneNumbersApi.md#createphonenumber) | **Post** /v1/Trunks/{TrunkSid}/PhoneNumbers | 
*TrunksPhoneNumbersApi* | [**DeletePhoneNumber**](docs/TrunksPhoneNumbersApi.md#deletephonenumber) | **Delete** /v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid} | 
*TrunksPhoneNumbersApi* | [**FetchPhoneNumber**](docs/TrunksPhoneNumbersApi.md#fetchphonenumber) | **Get** /v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid} | 
*TrunksPhoneNumbersApi* | [**ListPhoneNumber**](docs/TrunksPhoneNumbersApi.md#listphonenumber) | **Get** /v1/Trunks/{TrunkSid}/PhoneNumbers | 
*TrunksRecordingApi* | [**FetchRecording**](docs/TrunksRecordingApi.md#fetchrecording) | **Get** /v1/Trunks/{TrunkSid}/Recording | 
*TrunksRecordingApi* | [**UpdateRecording**](docs/TrunksRecordingApi.md#updaterecording) | **Post** /v1/Trunks/{TrunkSid}/Recording | 


## Documentation For Models

 - [ListCredentialListResponse](docs/ListCredentialListResponse.md)
 - [TrunkingV1CredentialList](docs/TrunkingV1CredentialList.md)
 - [TrunkingV1PhoneNumber](docs/TrunkingV1PhoneNumber.md)
 - [ListCredentialListResponseMeta](docs/ListCredentialListResponseMeta.md)
 - [TrunkingV1Trunk](docs/TrunkingV1Trunk.md)
 - [TrunkingV1IpAccessControlList](docs/TrunkingV1IpAccessControlList.md)
 - [ListIpAccessControlListResponse](docs/ListIpAccessControlListResponse.md)
 - [ListOriginationUrlResponse](docs/ListOriginationUrlResponse.md)
 - [ListPhoneNumberResponse](docs/ListPhoneNumberResponse.md)
 - [ListTrunkResponse](docs/ListTrunkResponse.md)
 - [TrunkingV1Recording](docs/TrunkingV1Recording.md)
 - [TrunkingV1OriginationUrl](docs/TrunkingV1OriginationUrl.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

