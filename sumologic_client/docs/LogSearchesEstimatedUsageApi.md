# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogSearchEstimatedUsage**](LogSearchesEstimatedUsageApi.md#GetLogSearchEstimatedUsage) | **Post** /v1/logSearches/estimatedUsage | Gets estimated usage details.

# **GetLogSearchEstimatedUsage**
> LogSearchEstimatedUsageDefinition GetLogSearchEstimatedUsage(ctx, body)
Gets estimated usage details.

Gets the estimated volume of data that would be scanned for a given log search in the Infrequent data tier. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogSearchEstimatedUsageRequest**](LogSearchEstimatedUsageRequest.md)| The definition of the log search estimated usage. | 

### Return type

[**LogSearchEstimatedUsageDefinition**](LogSearchEstimatedUsageDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

