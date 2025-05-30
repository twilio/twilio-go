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

All URIs are relative to *https://supersim.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ESimProfilesApi* | [**CreateEsimProfile**](docs/ESimProfilesApi.md#createesimprofile) | **Post** /v1/ESimProfiles | Order an eSIM Profile.
*ESimProfilesApi* | [**FetchEsimProfile**](docs/ESimProfilesApi.md#fetchesimprofile) | **Get** /v1/ESimProfiles/{Sid} | Fetch an eSIM Profile.
*ESimProfilesApi* | [**ListEsimProfile**](docs/ESimProfilesApi.md#listesimprofile) | **Get** /v1/ESimProfiles | Retrieve a list of eSIM Profiles.
*FleetsApi* | [**CreateFleet**](docs/FleetsApi.md#createfleet) | **Post** /v1/Fleets | Create a Fleet
*FleetsApi* | [**FetchFleet**](docs/FleetsApi.md#fetchfleet) | **Get** /v1/Fleets/{Sid} | Fetch a Fleet instance from your account.
*FleetsApi* | [**ListFleet**](docs/FleetsApi.md#listfleet) | **Get** /v1/Fleets | Retrieve a list of Fleets from your account.
*FleetsApi* | [**UpdateFleet**](docs/FleetsApi.md#updatefleet) | **Post** /v1/Fleets/{Sid} | Updates the given properties of a Super SIM Fleet instance from your account.
*IpCommandsApi* | [**CreateIpCommand**](docs/IpCommandsApi.md#createipcommand) | **Post** /v1/IpCommands | Send an IP Command to a Super SIM.
*IpCommandsApi* | [**FetchIpCommand**](docs/IpCommandsApi.md#fetchipcommand) | **Get** /v1/IpCommands/{Sid} | Fetch IP Command instance from your account.
*IpCommandsApi* | [**ListIpCommand**](docs/IpCommandsApi.md#listipcommand) | **Get** /v1/IpCommands | Retrieve a list of IP Commands from your account.
*NetworkAccessProfilesApi* | [**CreateNetworkAccessProfile**](docs/NetworkAccessProfilesApi.md#createnetworkaccessprofile) | **Post** /v1/NetworkAccessProfiles | Create a new Network Access Profile
*NetworkAccessProfilesApi* | [**FetchNetworkAccessProfile**](docs/NetworkAccessProfilesApi.md#fetchnetworkaccessprofile) | **Get** /v1/NetworkAccessProfiles/{Sid} | Fetch a Network Access Profile instance from your account.
*NetworkAccessProfilesApi* | [**ListNetworkAccessProfile**](docs/NetworkAccessProfilesApi.md#listnetworkaccessprofile) | **Get** /v1/NetworkAccessProfiles | Retrieve a list of Network Access Profiles from your account.
*NetworkAccessProfilesApi* | [**UpdateNetworkAccessProfile**](docs/NetworkAccessProfilesApi.md#updatenetworkaccessprofile) | **Post** /v1/NetworkAccessProfiles/{Sid} | Updates the given properties of a Network Access Profile in your account.
*NetworkAccessProfilesNetworksApi* | [**CreateNetworkAccessProfileNetwork**](docs/NetworkAccessProfilesNetworksApi.md#createnetworkaccessprofilenetwork) | **Post** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | Add a Network resource to the Network Access Profile resource.
*NetworkAccessProfilesNetworksApi* | [**DeleteNetworkAccessProfileNetwork**](docs/NetworkAccessProfilesNetworksApi.md#deletenetworkaccessprofilenetwork) | **Delete** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | Remove a Network resource from the Network Access Profile resource&#39;s.
*NetworkAccessProfilesNetworksApi* | [**FetchNetworkAccessProfileNetwork**](docs/NetworkAccessProfilesNetworksApi.md#fetchnetworkaccessprofilenetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | Fetch a Network Access Profile resource&#39;s Network resource.
*NetworkAccessProfilesNetworksApi* | [**ListNetworkAccessProfileNetwork**](docs/NetworkAccessProfilesNetworksApi.md#listnetworkaccessprofilenetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | Retrieve a list of Network Access Profile resource&#39;s Network resource.
*NetworksApi* | [**FetchNetwork**](docs/NetworksApi.md#fetchnetwork) | **Get** /v1/Networks/{Sid} | Fetch a Network resource.
*NetworksApi* | [**ListNetwork**](docs/NetworksApi.md#listnetwork) | **Get** /v1/Networks | Retrieve a list of Network resources.
*SettingsUpdatesApi* | [**ListSettingsUpdate**](docs/SettingsUpdatesApi.md#listsettingsupdate) | **Get** /v1/SettingsUpdates | Retrieve a list of Settings Updates.
*SimsApi* | [**CreateSim**](docs/SimsApi.md#createsim) | **Post** /v1/Sims | Register a Super SIM to your Account
*SimsApi* | [**FetchSim**](docs/SimsApi.md#fetchsim) | **Get** /v1/Sims/{Sid} | Fetch a Super SIM instance from your account.
*SimsApi* | [**ListSim**](docs/SimsApi.md#listsim) | **Get** /v1/Sims | Retrieve a list of Super SIMs from your account.
*SimsApi* | [**UpdateSim**](docs/SimsApi.md#updatesim) | **Post** /v1/Sims/{Sid} | Updates the given properties of a Super SIM instance from your account.
*SimsBillingPeriodsApi* | [**ListBillingPeriod**](docs/SimsBillingPeriodsApi.md#listbillingperiod) | **Get** /v1/Sims/{SimSid}/BillingPeriods | Retrieve a list of Billing Periods for a Super SIM.
*SimsIpAddressesApi* | [**ListSimIpAddress**](docs/SimsIpAddressesApi.md#listsimipaddress) | **Get** /v1/Sims/{SimSid}/IpAddresses | Retrieve a list of IP Addresses for the given Super SIM.
*SmsCommandsApi* | [**CreateSmsCommand**](docs/SmsCommandsApi.md#createsmscommand) | **Post** /v1/SmsCommands | Send SMS Command to a Sim.
*SmsCommandsApi* | [**FetchSmsCommand**](docs/SmsCommandsApi.md#fetchsmscommand) | **Get** /v1/SmsCommands/{Sid} | Fetch SMS Command instance from your account.
*SmsCommandsApi* | [**ListSmsCommand**](docs/SmsCommandsApi.md#listsmscommand) | **Get** /v1/SmsCommands | Retrieve a list of SMS Commands from your account.
*UsageRecordsApi* | [**ListUsageRecord**](docs/UsageRecordsApi.md#listusagerecord) | **Get** /v1/UsageRecords | List UsageRecords


## Documentation For Models

 - [SupersimV1Network](docs/SupersimV1Network.md)
 - [SupersimV1NetworkAccessProfileNetwork](docs/SupersimV1NetworkAccessProfileNetwork.md)
 - [ListFleetResponse](docs/ListFleetResponse.md)
 - [ListIpCommandResponse](docs/ListIpCommandResponse.md)
 - [ListNetworkAccessProfileNetworkResponse](docs/ListNetworkAccessProfileNetworkResponse.md)
 - [SupersimV1SmsCommand](docs/SupersimV1SmsCommand.md)
 - [SupersimV1EsimProfile](docs/SupersimV1EsimProfile.md)
 - [ListBillingPeriodResponse](docs/ListBillingPeriodResponse.md)
 - [ListNetworkResponse](docs/ListNetworkResponse.md)
 - [ListEsimProfileResponse](docs/ListEsimProfileResponse.md)
 - [ListNetworkAccessProfileResponse](docs/ListNetworkAccessProfileResponse.md)
 - [SupersimV1NetworkAccessProfile](docs/SupersimV1NetworkAccessProfile.md)
 - [ListSimIpAddressResponse](docs/ListSimIpAddressResponse.md)
 - [ListUsageRecordResponse](docs/ListUsageRecordResponse.md)
 - [ListSettingsUpdateResponse](docs/ListSettingsUpdateResponse.md)
 - [SupersimV1SimIpAddress](docs/SupersimV1SimIpAddress.md)
 - [ListSimResponse](docs/ListSimResponse.md)
 - [SupersimV1Fleet](docs/SupersimV1Fleet.md)
 - [SupersimV1SettingsUpdate](docs/SupersimV1SettingsUpdate.md)
 - [SupersimV1IpCommand](docs/SupersimV1IpCommand.md)
 - [ListSmsCommandResponse](docs/ListSmsCommandResponse.md)
 - [SupersimV1Sim](docs/SupersimV1Sim.md)
 - [ListBillingPeriodResponseMeta](docs/ListBillingPeriodResponseMeta.md)
 - [SupersimV1UsageRecord](docs/SupersimV1UsageRecord.md)
 - [SupersimV1BillingPeriod](docs/SupersimV1BillingPeriod.md)


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

