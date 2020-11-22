# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](ConnectionManagementApi.md#CreateConnection) | **Post** /v1/connections | Create a new connection.
[**DeleteConnection**](ConnectionManagementApi.md#DeleteConnection) | **Delete** /v1/connections/{id} | Delete a connection.
[**GetConnection**](ConnectionManagementApi.md#GetConnection) | **Get** /v1/connections/{id} | Get a connection.
[**ListConnections**](ConnectionManagementApi.md#ListConnections) | **Get** /v1/connections | Get a list of connections.
[**TestConnection**](ConnectionManagementApi.md#TestConnection) | **Post** /v1/connections/test | Test a new connection url.
[**UpdateConnection**](ConnectionManagementApi.md#UpdateConnection) | **Put** /v1/connections/{id} | Update a connection.

# **CreateConnection**
> Connection CreateConnection(ctx, body)
Create a new connection.

Create a new connection in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionDefinition**](ConnectionDefinition.md)| Information about the new connection. | 

### Return type

[**Connection**](Connection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnection**
> DeleteConnection(ctx, id, type_)
Delete a connection.

Delete a connection with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the connection to delete. | 
  **type_** | **string**| Type of connection to delete. Valid values are &#x60;WebhookConnection&#x60;, &#x60;ServiceNowConnection&#x60;. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnection**
> Connection GetConnection(ctx, id, type_)
Get a connection.

Get a connection with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of connection to return. | 
  **type_** | **string**| Type of connection to return. Valid values are &#x60;WebhookConnection&#x60;, &#x60;ServiceNowConnection&#x60;. | [default to WebhookConnection]

### Return type

[**Connection**](Connection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConnections**
> ListConnectionsResponse ListConnections(ctx, optional)
Get a list of connections.

Get a list of all connections in the organization. The response is paginated with a default limit of 100 connections per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectionManagementApiListConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionManagementApiListConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of connections returned in the response. The number of connections returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListConnectionsResponse**](ListConnectionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestConnection**
> TestConnectionResponse TestConnection(ctx, body)
Test a new connection url.

Test a new connection url is valid and can connect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionDefinition**](ConnectionDefinition.md)| Information about the new connection. | 

### Return type

[**TestConnectionResponse**](TestConnectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnection**
> Connection UpdateConnection(ctx, body, id)
Update a connection.

Update an existing connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionDefinition**](ConnectionDefinition.md)| Information to update about the connection. | 
  **id** | **string**| Identifier of the connection to update. | 

### Return type

[**Connection**](Connection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

