# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIngestBudgetV2**](IngestBudgetManagementV2Api.md#CreateIngestBudgetV2) | **Post** /v2/ingestBudgets | Create a new ingest budget.
[**DeleteIngestBudgetV2**](IngestBudgetManagementV2Api.md#DeleteIngestBudgetV2) | **Delete** /v2/ingestBudgets/{id} | Delete an ingest budget.
[**GetIngestBudgetV2**](IngestBudgetManagementV2Api.md#GetIngestBudgetV2) | **Get** /v2/ingestBudgets/{id} | Get an ingest budget.
[**ListIngestBudgetsV2**](IngestBudgetManagementV2Api.md#ListIngestBudgetsV2) | **Get** /v2/ingestBudgets | Get a list of ingest budgets.
[**ResetUsageV2**](IngestBudgetManagementV2Api.md#ResetUsageV2) | **Post** /v2/ingestBudgets/{id}/usage/reset | Reset usage.
[**UpdateIngestBudgetV2**](IngestBudgetManagementV2Api.md#UpdateIngestBudgetV2) | **Put** /v2/ingestBudgets/{id} | Update an ingest budget.

# **CreateIngestBudgetV2**
> IngestBudgetV2 CreateIngestBudgetV2(ctx, body)
Create a new ingest budget.

Create a new ingest budget.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IngestBudgetDefinitionV2**](IngestBudgetDefinitionV2.md)| Information about the new ingest budget. | 

### Return type

[**IngestBudgetV2**](IngestBudgetV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIngestBudgetV2**
> DeleteIngestBudgetV2(ctx, id)
Delete an ingest budget.

Delete an ingest budget with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ingest budget to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIngestBudgetV2**
> IngestBudgetV2 GetIngestBudgetV2(ctx, id)
Get an ingest budget.

Get an ingest budget by the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of ingest budget to return. | 

### Return type

[**IngestBudgetV2**](IngestBudgetV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIngestBudgetsV2**
> ListIngestBudgetsResponseV2 ListIngestBudgetsV2(ctx, optional)
Get a list of ingest budgets.

Get a list of all ingest budgets. The response is paginated with a default limit of 100 budgets per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IngestBudgetManagementV2ApiListIngestBudgetsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IngestBudgetManagementV2ApiListIngestBudgetsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of budgets returned in the response. The number of budgets returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ListIngestBudgetsResponseV2**](ListIngestBudgetsResponseV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetUsageV2**
> ResetUsageV2(ctx, id)
Reset usage.

Reset ingest budget's current usage to 0 before the scheduled reset time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ingest budget to reset usage. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIngestBudgetV2**
> IngestBudgetV2 UpdateIngestBudgetV2(ctx, body, id)
Update an ingest budget.

Update an existing ingest budget. All properties specified in the request are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IngestBudgetDefinitionV2**](IngestBudgetDefinitionV2.md)| Information to update about the ingest budget. | 
  **id** | **string**| Identifier of the ingest budget to update. | 

### Return type

[**IngestBudgetV2**](IngestBudgetV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

