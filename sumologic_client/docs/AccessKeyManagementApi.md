# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessKey**](AccessKeyManagementApi.md#CreateAccessKey) | **Post** /v1/accessKeys | Create an access key.
[**DeleteAccessKey**](AccessKeyManagementApi.md#DeleteAccessKey) | **Delete** /v1/accessKeys/{id} | Delete an access key.
[**ListAccessKeys**](AccessKeyManagementApi.md#ListAccessKeys) | **Get** /v1/accessKeys | List all access keys.
[**ListPersonalAccessKeys**](AccessKeyManagementApi.md#ListPersonalAccessKeys) | **Get** /v1/accessKeys/personal | List personal keys.
[**UpdateAccessKey**](AccessKeyManagementApi.md#UpdateAccessKey) | **Put** /v1/accessKeys/{id} | Update an access key.

# **CreateAccessKey**
> AccessKey CreateAccessKey(ctx, body)
Create an access key.

Creates a new access ID and key pair. The new access key can be used from the domains specified in corsHeaders field. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don't match any entry in the allowlist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccessKeyCreateRequest**](AccessKeyCreateRequest.md)|  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccessKey**
> DeleteAccessKey(ctx, id)
Delete an access key.

Deletes the access key with the given accessId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The accessId of the access key to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccessKeys**
> PaginatedListAccessKeysResult ListAccessKeys(ctx, optional)
List all access keys.

List all access keys in your account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccessKeyManagementApiListAccessKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccessKeyManagementApiListAccessKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of access keys returned in the response. The number of access keys returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**PaginatedListAccessKeysResult**](PaginatedListAccessKeysResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPersonalAccessKeys**
> ListAccessKeysResult ListPersonalAccessKeys(ctx, )
List personal keys.

List all access keys that belong to your user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListAccessKeysResult**](ListAccessKeysResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccessKey**
> AccessKeyPublic UpdateAccessKey(ctx, body, id)
Update an access key.

Updates the properties of existing accessKey by accessId. It can be used to enable or disable the access key and to update the corsHeaders list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccessKeyUpdateRequest**](AccessKeyUpdateRequest.md)|  | 
  **id** | **string**| The accessId of the access key to update. | 

### Return type

[**AccessKeyPublic**](AccessKeyPublic.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

