# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignCollectorToBudget**](IngestBudgetManagementV1Api.md#AssignCollectorToBudget) | **Put** /v1/ingestBudgets/{id}/collectors/{collectorId} | Assign a Collector to a budget.
[**CreateIngestBudget**](IngestBudgetManagementV1Api.md#CreateIngestBudget) | **Post** /v1/ingestBudgets | Create a new ingest budget.
[**DeleteIngestBudget**](IngestBudgetManagementV1Api.md#DeleteIngestBudget) | **Delete** /v1/ingestBudgets/{id} | Delete an ingest budget.
[**GetAssignedCollectors**](IngestBudgetManagementV1Api.md#GetAssignedCollectors) | **Get** /v1/ingestBudgets/{id}/collectors | Get a list of Collectors.
[**GetIngestBudget**](IngestBudgetManagementV1Api.md#GetIngestBudget) | **Get** /v1/ingestBudgets/{id} | Get an ingest budget.
[**ListIngestBudgets**](IngestBudgetManagementV1Api.md#ListIngestBudgets) | **Get** /v1/ingestBudgets | Get a list of ingest budgets.
[**RemoveCollectorFromBudget**](IngestBudgetManagementV1Api.md#RemoveCollectorFromBudget) | **Delete** /v1/ingestBudgets/{id}/collectors/{collectorId} | Remove Collector from a budget.
[**ResetUsage**](IngestBudgetManagementV1Api.md#ResetUsage) | **Post** /v1/ingestBudgets/{id}/usage/reset | Reset usage.
[**UpdateIngestBudget**](IngestBudgetManagementV1Api.md#UpdateIngestBudget) | **Put** /v1/ingestBudgets/{id} | Update an ingest budget.

# **AssignCollectorToBudget**
> IngestBudget AssignCollectorToBudget(ctx, id, collectorId)
Assign a Collector to a budget.

Assign a Collector to a budget.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ingest budget to assign to the Collector. | 
  **collectorId** | **string**| Identifier of the Collector to assign. | 

### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIngestBudget**
> IngestBudget CreateIngestBudget(ctx, body)
Create a new ingest budget.

Create a new ingest budget.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IngestBudgetDefinition**](IngestBudgetDefinition.md)| Information about the new ingest budget. | 

### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIngestBudget**
> DeleteIngestBudget(ctx, id)
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

# **GetAssignedCollectors**
> ListCollectorIdentitiesResponse GetAssignedCollectors(ctx, id, optional)
Get a list of Collectors.

Get a list of Collectors assigned to an ingest budget. The response is paginated with a default limit of 100 Collectors per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of ingest budget to which Collectors are assigned. | 
 **optional** | ***IngestBudgetManagementV1ApiGetAssignedCollectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IngestBudgetManagementV1ApiGetAssignedCollectorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Limit the number of Collectors returned in the response. The number of Collectors returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ListCollectorIdentitiesResponse**](ListCollectorIdentitiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIngestBudget**
> IngestBudget GetIngestBudget(ctx, id)
Get an ingest budget.

Get an ingest budget by the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of ingest budget to return. | 

### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIngestBudgets**
> ListIngestBudgetsResponse ListIngestBudgets(ctx, optional)
Get a list of ingest budgets.

Get a list of all ingest budgets. The response is paginated with a default limit of 100 budgets per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IngestBudgetManagementV1ApiListIngestBudgetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IngestBudgetManagementV1ApiListIngestBudgetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of budgets returned in the response. The number of budgets returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ListIngestBudgetsResponse**](ListIngestBudgetsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCollectorFromBudget**
> IngestBudget RemoveCollectorFromBudget(ctx, id, collectorId)
Remove Collector from a budget.

Remove Collector from a budget.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ingest budget to unassign from the Collector. | 
  **collectorId** | **string**| Identifier of the Collector to unassign. | 

### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetUsage**
> ResetUsage(ctx, id)
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

# **UpdateIngestBudget**
> IngestBudget UpdateIngestBudget(ctx, body, id)
Update an ingest budget.

Update an existing ingest budget. All properties specified in the request are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IngestBudgetDefinition**](IngestBudgetDefinition.md)| Information to update about the ingest budget. | 
  **id** | **string**| Identifier of the ingest budget to update. | 

### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

