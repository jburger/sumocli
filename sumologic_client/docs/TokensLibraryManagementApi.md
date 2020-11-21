# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokensLibraryManagementApi.md#CreateToken) | **Post** /v1/tokens | Create a token.
[**DeleteToken**](TokensLibraryManagementApi.md#DeleteToken) | **Delete** /v1/tokens/{id} | Delete a token.
[**GetToken**](TokensLibraryManagementApi.md#GetToken) | **Get** /v1/tokens/{id} | Get a token.
[**ListTokens**](TokensLibraryManagementApi.md#ListTokens) | **Get** /v1/tokens | Get a list of tokens.
[**UpdateToken**](TokensLibraryManagementApi.md#UpdateToken) | **Put** /v1/tokens/{id} | Update a token.

# **CreateToken**
> TokenBaseResponse CreateToken(ctx, body)
Create a token.

Create a token in the token library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TokenBaseDefinition**](TokenBaseDefinition.md)| Information about the token to create. | 

### Return type

[**TokenBaseResponse**](TokenBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteToken**
> DeleteToken(ctx, id)
Delete a token.

Delete a token with the given identifier in the token library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the token to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetToken**
> TokenBaseResponse GetToken(ctx, id)
Get a token.

Get a token with the given identifier in the token library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the token to return. | 

### Return type

[**TokenBaseResponse**](TokenBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTokens**
> ListTokensBaseResponse ListTokens(ctx, )
Get a list of tokens.

Get a list of all tokens in the token library.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListTokensBaseResponse**](ListTokensBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateToken**
> TokenBaseResponse UpdateToken(ctx, body, id)
Update a token.

Update a token with the given identifier in the token library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TokenBaseDefinitionUpdate**](TokenBaseDefinitionUpdate.md)| The token to update. | 
  **id** | **string**| Identifier of the token to update. | 

### Return type

[**TokenBaseResponse**](TokenBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

