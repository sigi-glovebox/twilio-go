# ServicesAccessTokensApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](ServicesAccessTokensApi.md#CreateAccessToken) | **Post** /v2/Services/{ServiceSid}/AccessTokens | Create a new enrollment Access Token for the Entity
[**FetchAccessToken**](ServicesAccessTokensApi.md#FetchAccessToken) | **Get** /v2/Services/{ServiceSid}/AccessTokens/{Sid} | Fetch an Access Token for the Entity



## CreateAccessToken

> VerifyV2AccessToken CreateAccessToken(ctx, ServiceSidoptional)

Create a new enrollment Access Token for the Entity

Create a new enrollment Access Token for the Entity

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a CreateAccessTokenParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **string** | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user's UUID, GUID, or SID.
**FactorType** | [**string**](string.md) | 
**FactorFriendlyName** | **string** | The friendly name of the factor that is going to be created with this access token
**Ttl** | **int** | How long, in seconds, the access token is valid. Can be an integer between 60 and 300. Default is 60.

### Return type

[**VerifyV2AccessToken**](VerifyV2AccessToken.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccessToken

> VerifyV2AccessToken FetchAccessToken(ctx, ServiceSidSid)

Fetch an Access Token for the Entity

Fetch an Access Token for the Entity

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Sid** | **string** | A 34 character string that uniquely identifies this Access Token.

### Other Parameters

Other parameters are passed through a pointer to a FetchAccessTokenParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2AccessToken**](VerifyV2AccessToken.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

