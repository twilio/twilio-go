# InlineObject11

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential. | [optional] 
**Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. | [optional] 
**PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60;. | [optional] 
**Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production. | [optional] 
**Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging. | [optional] 
**Type** | **string** | The type of push-notification service the credential is for. Can be: &#x60;fcm&#x60;, &#x60;gcm&#x60;, or &#x60;apn&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


