# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FolderManagementApi.md#CreateFolder) | **Post** /v2/content/folders | Create a new folder.
[**GetAdminRecommendedFolderAsync**](FolderManagementApi.md#GetAdminRecommendedFolderAsync) | **Get** /v2/content/folders/adminRecommended | Get Admin Recommended folder.
[**GetAdminRecommendedFolderAsyncResult**](FolderManagementApi.md#GetAdminRecommendedFolderAsyncResult) | **Get** /v2/content/folders/adminRecommended/{jobId}/result | Admin Recommended folder job result.
[**GetAdminRecommendedFolderAsyncStatus**](FolderManagementApi.md#GetAdminRecommendedFolderAsyncStatus) | **Get** /v2/content/folders/adminRecommended/{jobId}/status | Admin Recommended folder job status.
[**GetFolder**](FolderManagementApi.md#GetFolder) | **Get** /v2/content/folders/{id} | Get a folder.
[**GetGlobalFolderAsync**](FolderManagementApi.md#GetGlobalFolderAsync) | **Get** /v2/content/folders/global | Get global folder.
[**GetGlobalFolderAsyncResult**](FolderManagementApi.md#GetGlobalFolderAsyncResult) | **Get** /v2/content/folders/global/{jobId}/result | Global folder job result.
[**GetGlobalFolderAsyncStatus**](FolderManagementApi.md#GetGlobalFolderAsyncStatus) | **Get** /v2/content/folders/global/{jobId}/status | Global folder job status.
[**GetPersonalFolder**](FolderManagementApi.md#GetPersonalFolder) | **Get** /v2/content/folders/personal | Get personal folder.
[**UpdateFolder**](FolderManagementApi.md#UpdateFolder) | **Put** /v2/content/folders/{id} | Update a folder.

# **CreateFolder**
> Folder CreateFolder(ctx, body, optional)
Create a new folder.

Creates a new folder under the given parent folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FolderDefinition**](FolderDefinition.md)| Information about the new folder. | 
 **optional** | ***FolderManagementApiCreateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FolderManagementApiCreateFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isAdminMode** | **optional.**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdminRecommendedFolderAsync**
> BeginAsyncJobResponse GetAdminRecommendedFolderAsync(ctx, optional)
Get Admin Recommended folder.

Schedule an asynchronous job to get the top-level Admin Recommended content items.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FolderManagementApiGetAdminRecommendedFolderAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FolderManagementApiGetAdminRecommendedFolderAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdminRecommendedFolderAsyncResult**
> Folder GetAdminRecommendedFolderAsyncResult(ctx, jobId)
Admin Recommended folder job result.

Get results from Admin Recommended job for the given job identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The identifier of the asynchronous Admin Recommended folder job. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdminRecommendedFolderAsyncStatus**
> AsyncJobStatus GetAdminRecommendedFolderAsyncStatus(ctx, jobId)
Admin Recommended folder job status.

Get the status of an asynchronous Admin Recommended folder job for the given job identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The identifier of the asynchronous Admin Recommended folder job. | 

### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolder**
> Folder GetFolder(ctx, id, optional)
Get a folder.

Get a folder with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the folder to fetch. | 
 **optional** | ***FolderManagementApiGetFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FolderManagementApiGetFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalFolderAsync**
> BeginAsyncJobResponse GetGlobalFolderAsync(ctx, optional)
Get global folder.

Schedule an asynchronous job to get global folder. Global folder contains all content items that a user has permissions to view in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FolderManagementApiGetGlobalFolderAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FolderManagementApiGetGlobalFolderAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalFolderAsyncResult**
> ContentList GetGlobalFolderAsyncResult(ctx, jobId)
Global folder job result.

Get results from global folder job for the given job identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The identifier of the asynchronous global folder job. | 

### Return type

[**ContentList**](ContentList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalFolderAsyncStatus**
> AsyncJobStatus GetGlobalFolderAsyncStatus(ctx, jobId)
Global folder job status.

Get the status of an asynchronous global folder job for the given job identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The identifier of the asynchronous global folder job. | 

### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPersonalFolder**
> Folder GetPersonalFolder(ctx, )
Get personal folder.

Get the personal folder of the current user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFolder**
> Folder UpdateFolder(ctx, body, id, optional)
Update a folder.

Update an existing folder with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFolderRequest**](UpdateFolderRequest.md)| Information to update about the folder. | 
  **id** | **string**| Identifier of the folder to update. | 
 **optional** | ***FolderManagementApiUpdateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FolderManagementApiUpdateFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

