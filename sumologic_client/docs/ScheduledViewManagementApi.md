# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScheduledView**](ScheduledViewManagementApi.md#CreateScheduledView) | **Post** /v1/scheduledViews | Create a new scheduled view.
[**DisableScheduledView**](ScheduledViewManagementApi.md#DisableScheduledView) | **Delete** /v1/scheduledViews/{id}/disable | Disable a scheduled view.
[**GetScheduledView**](ScheduledViewManagementApi.md#GetScheduledView) | **Get** /v1/scheduledViews/{id} | Get a scheduled view.
[**ListScheduledViews**](ScheduledViewManagementApi.md#ListScheduledViews) | **Get** /v1/scheduledViews | Get a list of scheduled views.
[**PauseScheduledView**](ScheduledViewManagementApi.md#PauseScheduledView) | **Post** /v1/scheduledViews/{id}/pause | Pause a scheduled view.
[**StartScheduledView**](ScheduledViewManagementApi.md#StartScheduledView) | **Post** /v1/scheduledViews/{id}/start | Start a scheduled view.
[**UpdateScheduledView**](ScheduledViewManagementApi.md#UpdateScheduledView) | **Put** /v1/scheduledViews/{id} | Update a scheduled view.

# **CreateScheduledView**
> ScheduledView CreateScheduledView(ctx, body)
Create a new scheduled view.

Creates a new scheduled view in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateScheduledViewDefinition**](CreateScheduledViewDefinition.md)| Information about the new scheduled view. | 

### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableScheduledView**
> DisableScheduledView(ctx, id)
Disable a scheduled view.

Disable a scheduled view with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the scheduled view to disable. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScheduledView**
> ScheduledView GetScheduledView(ctx, id)
Get a scheduled view.

Get a scheduled view with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the scheduled view to fetch. | 

### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListScheduledViews**
> ListScheduledViewsResponse ListScheduledViews(ctx, optional)
Get a list of scheduled views.

Get a list of all scheduled views in the organization. The response is paginated with a default limit of 100 scheduled views per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScheduledViewManagementApiListScheduledViewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduledViewManagementApiListScheduledViewsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of scheduled views returned in the response. The number of scheduled views returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListScheduledViewsResponse**](ListScheduledViewsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PauseScheduledView**
> ScheduledView PauseScheduledView(ctx, id)
Pause a scheduled view.

Pause a scheduled view with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the scheduled view to pause. | 

### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartScheduledView**
> ScheduledView StartScheduledView(ctx, id)
Start a scheduled view.

Start a scheduled view with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the scheduled view to start. | 

### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateScheduledView**
> ScheduledView UpdateScheduledView(ctx, body, id)
Update a scheduled view.

Update an existing scheduled view.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateScheduledViewDefinition**](UpdateScheduledViewDefinition.md)| Information to update about the scheduled view. | 
  **id** | **string**| Identifier of the scheduled view to update. | 

### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

