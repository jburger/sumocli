# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricsSearch**](MetricsSearchesManagementApi.md#CreateMetricsSearch) | **Post** /v1/metricsSearches | Save a metrics search.
[**DeleteMetricsSearch**](MetricsSearchesManagementApi.md#DeleteMetricsSearch) | **Delete** /v1/metricsSearches/{id} | Deletes a metrics search.
[**GetMetricsSearch**](MetricsSearchesManagementApi.md#GetMetricsSearch) | **Get** /v1/metricsSearches/{id} | Get a metrics search.
[**UpdateMetricsSearch**](MetricsSearchesManagementApi.md#UpdateMetricsSearch) | **Put** /v1/metricsSearches/{id} | Updates a metrics search.

# **CreateMetricsSearch**
> MetricsSearchInstance CreateMetricsSearch(ctx, body)
Save a metrics search.

Saves a metrics search in the content library. Metrics search consists of one or more queries, a time range, a quantization period and a set of chart properties like line width.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveMetricsSearchRequest**](SaveMetricsSearchRequest.md)| The definition of the metrics search. | 

### Return type

[**MetricsSearchInstance**](MetricsSearchInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMetricsSearch**
> DeleteMetricsSearch(ctx, id)
Deletes a metrics search.

Deletes a metrics search from the content library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the metrics search. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetricsSearch**
> MetricsSearchInstance GetMetricsSearch(ctx, id)
Get a metrics search.

Returns a metrics search with the specified identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the metrics search. | 

### Return type

[**MetricsSearchInstance**](MetricsSearchInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMetricsSearch**
> MetricsSearchInstance UpdateMetricsSearch(ctx, body, id)
Updates a metrics search.

Updates a metrics search with the specified identifier. Partial updates are not supported, you must provide values for all fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MetricsSearchV1**](MetricsSearchV1.md)| An updated metrics search definition. | 
  **id** | **string**| Identifier of the metrics search. | 

### Return type

[**MetricsSearchInstance**](MetricsSearchInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

