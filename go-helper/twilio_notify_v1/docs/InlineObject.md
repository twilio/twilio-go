# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | [GCM only] The &#x60;Server key&#x60; of your project from Firebase console under Settings / Cloud messaging. | [optional] 
**Certificate** | **string** | [APN only] The URL-encoded representation of the certificate. Strip everything outside of the headers, e.g. &#x60;-----BEGIN CERTIFICATE-----MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A&#x3D;&#x3D;-----END CERTIFICATE-----&#x60; | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | [optional] 
**PrivateKey** | **string** | [APN only] The URL-encoded representation of the private key. Strip everything outside of the headers, e.g. &#x60;-----BEGIN RSA PRIVATE KEY-----MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR\\n.-----END RSA PRIVATE KEY-----&#x60; | [optional] 
**Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production. | [optional] 
**Secret** | **string** | [FCM only] The &#x60;Server key&#x60; of your project from Firebase console under Settings / Cloud messaging. | [optional] 
**Type** | **string** | The Credential type. Can be: &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


