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

All URIs are relative to *https://flex-api.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*InstancesUsersApi* | [**FetchFlexUser**](docs/InstancesUsersApi.md#fetchflexuser) | **Get** /v2/Instances/{InstanceSid}/Users/{FlexUserSid} | Fetch flex user for the given flex user sid
*InstancesUsersApi* | [**UpdateFlexUser**](docs/InstancesUsersApi.md#updateflexuser) | **Post** /v2/Instances/{InstanceSid}/Users/{FlexUserSid} | Update flex user for the given flex user sid
*WebChatsApi* | [**CreateWebChannel**](docs/WebChatsApi.md#createwebchannel) | **Post** /v2/WebChats | 


## Documentation For Models

 - [FlexV2FlexUser](docs/FlexV2FlexUser.md)
 - [FlexV2WebChannel](docs/FlexV2WebChannel.md)


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

