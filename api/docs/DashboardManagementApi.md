# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardManagementApi.md#CreateDashboard) | **Post** /v2/dashboards | Create a new dashboard.
[**DeleteDashboard**](DashboardManagementApi.md#DeleteDashboard) | **Delete** /v2/dashboards/{id} | Delete a dashboard.
[**GetDashboard**](DashboardManagementApi.md#GetDashboard) | **Get** /v2/dashboards/{id} | Get a dashboard.
[**UpdateDashboard**](DashboardManagementApi.md#UpdateDashboard) | **Put** /v2/dashboards/{id} | Update a dashboard.

# **CreateDashboard**
> Dashboard CreateDashboard(ctx, body)
Create a new dashboard.

Creates a new dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DashboardRequest**](DashboardRequest.md)| Information to create the new dashboard. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDashboard**
> DeleteDashboard(ctx, id)
Delete a dashboard.

Delete a dashboard by the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the dashboard to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboard**
> Dashboard GetDashboard(ctx, id)
Get a dashboard.

Get a dashboard by the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| UUID of the dashboard to return. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDashboard**
> Dashboard UpdateDashboard(ctx, body, id)
Update a dashboard.

Update a dashboard by the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DashboardRequest**](DashboardRequest.md)| Information to update on the dashboard. | 
  **id** | **string**| Identifier of the dashboard to update. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

