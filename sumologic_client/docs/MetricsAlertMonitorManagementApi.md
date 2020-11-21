# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMonitor**](MetricsAlertMonitorManagementApi.md#CreateMonitor) | **Post** /v1/metricsAlertMonitors | Create a new metrics monitor.
[**DeleteMonitor**](MetricsAlertMonitorManagementApi.md#DeleteMonitor) | **Delete** /v1/metricsAlertMonitors/{id} | Delete a metrics monitor.
[**GetMonitor**](MetricsAlertMonitorManagementApi.md#GetMonitor) | **Get** /v1/metricsAlertMonitors/{id} | Get a metrics monitor.
[**GetMonitors**](MetricsAlertMonitorManagementApi.md#GetMonitors) | **Get** /v1/metricsAlertMonitors | Get a list of metrics monitors.
[**Mute**](MetricsAlertMonitorManagementApi.md#Mute) | **Post** /v1/metricsAlertMonitors/{id}/mute | Mute a metrics monitor.
[**Unmute**](MetricsAlertMonitorManagementApi.md#Unmute) | **Post** /v1/metricsAlertMonitors/{id}/unmute | Unmute a metrics monitor.
[**UpdateMonitor**](MetricsAlertMonitorManagementApi.md#UpdateMonitor) | **Put** /v1/metricsAlertMonitors/{id} | Update a metrics monitor.

# **CreateMonitor**
> MetricsMonitorInstance CreateMonitor(ctx, body)
Create a new metrics monitor.

Create a new metrics monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MetricsMonitorDefinition**](MetricsMonitorDefinition.md)| Data to create a new metrics monitor. | 

### Return type

[**MetricsMonitorInstance**](MetricsMonitorInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMonitor**
> DeleteMonitor(ctx, id)
Delete a metrics monitor.

Delete an existing metrics monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the monitor to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitor**
> MetricsMonitorInstance GetMonitor(ctx, id)
Get a metrics monitor.

Get a metrics monitor with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the monitor to return. | 

### Return type

[**MetricsMonitorInstance**](MetricsMonitorInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitors**
> MetricsMonitorsResponse GetMonitors(ctx, optional)
Get a list of metrics monitors.

Get a list of all metrics monitors in the organization. The response is paginated with a default limit of 100 monitors per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetricsAlertMonitorManagementApiGetMonitorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsAlertMonitorManagementApiGetMonitorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of monitors returned in the response. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**MetricsMonitorsResponse**](MetricsMonitorsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Mute**
> MetricsMonitorMuteStatus Mute(ctx, id, optional)
Mute a metrics monitor.

Mute a metrics monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the monitor to mute. | 
 **optional** | ***MetricsAlertMonitorManagementApiMuteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsAlertMonitorManagementApiMuteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MuteMetricsMonitorRequest**](MuteMetricsMonitorRequest.md)|  | 

### Return type

[**MetricsMonitorMuteStatus**](MetricsMonitorMuteStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unmute**
> MetricsMonitorMuteStatus Unmute(ctx, id)
Unmute a metrics monitor.

Unmute a metrics monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the monitor to unmute. | 

### Return type

[**MetricsMonitorMuteStatus**](MetricsMonitorMuteStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMonitor**
> MetricsMonitorInstance UpdateMonitor(ctx, body, id)
Update a metrics monitor.

Update an existing metrics monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MetricsMonitorDefinition**](MetricsMonitorDefinition.md)| Information to update the existing metrics monitor. | 
  **id** | **string**| Identifier of the monitor to update. | 

### Return type

[**MetricsMonitorInstance**](MetricsMonitorInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

