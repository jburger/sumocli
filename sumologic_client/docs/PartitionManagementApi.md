# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRetentionUpdate**](PartitionManagementApi.md#CancelRetentionUpdate) | **Post** /v1/partitions/{id}/cancelRetentionUpdate | Cancel a retention update for a partition
[**CreatePartition**](PartitionManagementApi.md#CreatePartition) | **Post** /v1/partitions | Create a new partition.
[**DecommissionPartition**](PartitionManagementApi.md#DecommissionPartition) | **Post** /v1/partitions/{id}/decommission | Decommission a partition.
[**GetPartition**](PartitionManagementApi.md#GetPartition) | **Get** /v1/partitions/{id} | Get a partition.
[**ListPartitions**](PartitionManagementApi.md#ListPartitions) | **Get** /v1/partitions | Get a list of partitions.
[**UpdatePartition**](PartitionManagementApi.md#UpdatePartition) | **Put** /v1/partitions/{id} | Update a partition.

# **CancelRetentionUpdate**
> CancelRetentionUpdate(ctx, id)
Cancel a retention update for a partition

Cancel update to retention of a partition for which retention was updated previously using `reduceRetentionPeriodImmediately` parameter as false

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the partition to cancel the retention update for. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePartition**
> Partition CreatePartition(ctx, body)
Create a new partition.

Create a new partition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePartitionDefinition**](CreatePartitionDefinition.md)| Information about the new partition. | 

### Return type

[**Partition**](Partition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DecommissionPartition**
> DecommissionPartition(ctx, id)
Decommission a partition.

Decommission a partition with the given identifier from the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the partition to decommission. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPartition**
> Partition GetPartition(ctx, id)
Get a partition.

Get a partition with the given identifier from the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of partition to return. | 

### Return type

[**Partition**](Partition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPartitions**
> ListPartitionsResponse ListPartitions(ctx, optional)
Get a list of partitions.

Get a list of all partitions in the organization. The response is paginated with a default limit of 100 partitions per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PartitionManagementApiListPartitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartitionManagementApiListPartitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of partitions returned in the response. The number of partitions returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListPartitionsResponse**](ListPartitionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePartition**
> Partition UpdatePartition(ctx, body, id)
Update a partition.

Update an existing partition in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePartitionDefinition**](UpdatePartitionDefinition.md)| Information to update about the partition. | 
  **id** | **string**| Identifier of the partition to update. | 

### Return type

[**Partition**](Partition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

