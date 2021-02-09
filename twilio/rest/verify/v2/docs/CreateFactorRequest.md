# CreateFactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingAlg** | **string** | The algorithm used when &#x60;factor_type&#x60; is &#x60;push&#x60;. Algorithm supported: &#x60;ES256&#x60; | [optional] 
**BindingPublicKey** | **string** | The Ecdsa public key in PKIX, ASN.1 DER format encoded in Base64 | [optional] 
**ConfigAppId** | **string** | The ID that uniquely identifies your app in the Google or Apple store, such as &#x60;com.example.myapp&#x60;. Required when &#x60;factor_type&#x60; is &#x60;push&#x60; | [optional] 
**ConfigNotificationPlatform** | **string** | The transport technology used to generate the Notification Token. Can be &#x60;apn&#x60; or &#x60;fcm&#x60;. Required when &#x60;factor_type&#x60; is &#x60;push&#x60; | [optional] 
**ConfigNotificationToken** | **string** | For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when &#x60;factor_type&#x60; is &#x60;push&#x60; | [optional] 
**ConfigSdkVersion** | **string** | The Verify Push SDK version used to configure the factor | [optional] 
**FactorType** | **string** | The Type of this Factor. Currently only &#x60;push&#x60; is supported | 
**FriendlyName** | **string** | The friendly name of this Factor | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


