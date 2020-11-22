# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArchiveJob**](ArchiveManagementApi.md#CreateArchiveJob) | **Post** /v1/archive/{sourceId}/jobs | Create an ingestion job.
[**DeleteArchiveJob**](ArchiveManagementApi.md#DeleteArchiveJob) | **Delete** /v1/archive/{sourceId}/jobs/{id} | Delete an ingestion job.
[**ListArchiveJobsBySourceId**](ArchiveManagementApi.md#ListArchiveJobsBySourceId) | **Get** /v1/archive/{sourceId}/jobs | Get ingestion jobs for an Archive Source.
[**ListArchiveJobsCountPerSource**](ArchiveManagementApi.md#ListArchiveJobsCountPerSource) | **Get** /v1/archive/jobs/count | List ingestion jobs for all Archive Sources.

# **CreateArchiveJob**
> ArchiveJob CreateArchiveJob(ctx, body, sourceId)
Create an ingestion job.

Create an ingestion job to pull data from your S3 bucket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateArchiveJobRequest**](CreateArchiveJobRequest.md)| The definition of the ingestion job to create. | 
  **sourceId** | **string**| The identifier of the Archive Source for which the job is to be added. | 

### Return type

[**ArchiveJob**](ArchiveJob.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArchiveJob**
> DeleteArchiveJob(ctx, sourceId, id)
Delete an ingestion job.

Delete an ingestion job with the given identifier from the organization. The delete operation is only possible for jobs with a Succeeded or Failed status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceId** | **string**| The identifier of the Archive Source. | 
  **id** | **string**| The identifier of the ingestion job to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListArchiveJobsBySourceId**
> ListArchiveJobsResponse ListArchiveJobsBySourceId(ctx, sourceId, optional)
Get ingestion jobs for an Archive Source.

Get a list of all the ingestion jobs created on an Archive Source. The response is paginated with a default limit of 10 jobs per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceId** | **string**| The identifier of an Archive Source. | 
 **optional** | ***ArchiveManagementApiListArchiveJobsBySourceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArchiveManagementApiListArchiveJobsBySourceIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Limit the number of jobs returned in the response. The number of jobs returned may be less than the &#x60;limit&#x60;. | [default to 10]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListArchiveJobsResponse**](ListArchiveJobsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListArchiveJobsCountPerSource**
> ListArchiveJobsCount ListArchiveJobsCountPerSource(ctx, )
List ingestion jobs for all Archive Sources.

Get a list of all Archive Sources with the count and status of ingestion jobs.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListArchiveJobsCount**](ListArchiveJobsCount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

