# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncCopyStatus**](ContentManagementApi.md#AsyncCopyStatus) | **Get** /v2/content/{id}/copy/{jobId}/status | Content copy job status.
[**BeginAsyncCopy**](ContentManagementApi.md#BeginAsyncCopy) | **Post** /v2/content/{id}/copy | Start a content copy job.
[**BeginAsyncDelete**](ContentManagementApi.md#BeginAsyncDelete) | **Delete** /v2/content/{id}/delete | Start a content deletion job.
[**BeginAsyncExport**](ContentManagementApi.md#BeginAsyncExport) | **Post** /v2/content/{id}/export | Start a content export job.
[**BeginAsyncImport**](ContentManagementApi.md#BeginAsyncImport) | **Post** /v2/content/folders/{folderId}/import | Start a content import job.
[**GetAsyncDeleteStatus**](ContentManagementApi.md#GetAsyncDeleteStatus) | **Get** /v2/content/{id}/delete/{jobId}/status | Content deletion job status.
[**GetAsyncExportResult**](ContentManagementApi.md#GetAsyncExportResult) | **Get** /v2/content/{contentId}/export/{jobId}/result | Content export job result.
[**GetAsyncExportStatus**](ContentManagementApi.md#GetAsyncExportStatus) | **Get** /v2/content/{contentId}/export/{jobId}/status | Content export job status.
[**GetAsyncImportStatus**](ContentManagementApi.md#GetAsyncImportStatus) | **Get** /v2/content/folders/{folderId}/import/{jobId}/status | Content import job status.
[**GetItemByPath**](ContentManagementApi.md#GetItemByPath) | **Get** /v2/content/path | Get content item by path.
[**GetPathById**](ContentManagementApi.md#GetPathById) | **Get** /v2/content/{contentId}/path | Get path of an item.
[**MoveItem**](ContentManagementApi.md#MoveItem) | **Post** /v2/content/{id}/move | Move an item.

# **AsyncCopyStatus**
> AsyncJobStatus AsyncCopyStatus(ctx, id, jobId, optional)
Content copy job status.

Get the status of the copy request with the given job identifier. On success, field `statusMessage` will contain identifier of the newly copied content in format: `id: {hexIdentifier}`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The identifier of the content which was copied. | 
  **jobId** | **string**| The identifier of the asynchronous copy request job. | 
 **optional** | ***ContentManagementApiAsyncCopyStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiAsyncCopyStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BeginAsyncCopy**
> BeginAsyncJobResponse BeginAsyncCopy(ctx, id, destinationFolder, optional)
Start a content copy job.

Start an asynchronous content copy job with the given identifier to the destination folder. If the content item is a folder, everything under the folder is copied recursively.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The identifier of the content item to copy. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format. | 
  **destinationFolder** | **string**| The identifier of the destination folder. | 
 **optional** | ***ContentManagementApiBeginAsyncCopyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiBeginAsyncCopyOpts struct
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

# **BeginAsyncDelete**
> BeginAsyncJobResponse BeginAsyncDelete(ctx, id, optional)
Start a content deletion job.

Start an asynchronous content deletion job with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the content to delete. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format. | 
 **optional** | ***ContentManagementApiBeginAsyncDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiBeginAsyncDeleteOpts struct
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

# **BeginAsyncExport**
> BeginAsyncJobResponse BeginAsyncExport(ctx, id, optional)
Start a content export job.

Schedule an _asynchronous_ export of content with the given identifier. You will get back an asynchronous job identifier on success. Use the [getAsyncExportStatus](#operation/getAsyncExportStatus) endpoint and the job identifier you got back in the response to track the status of an asynchronous export job. If the content item is a folder, everything under the folder is exported recursively. Keep in mind when exporting large folders that there is a limit of 1000 content objects that can be exported at once. If you want to import more than 1000 content objects, then be sure to split the import into batches of 1000 objects or less. The results from the export are compatible with the Library import feature in the Sumo Logic user interface as well as the API content import job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The identifier of the content item to export. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format. | 
 **optional** | ***ContentManagementApiBeginAsyncExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiBeginAsyncExportOpts struct
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

# **BeginAsyncImport**
> BeginAsyncJobResponse BeginAsyncImport(ctx, body, folderId, optional)
Start a content import job.

Schedule an asynchronous import of content inside an existing folder with the given identifier. Import requests can be used to create or update content within a folder. Content items need to have a unique name within their folder. If there is already a content item with the same name in the folder, you can set the `overwrite` parameter to `true` to overwrite existing content items. By default, the `overwrite` parameter is set to `false`, where the import will fail if a content item with the same name already exist. Keep in mind when importing large folders that there is a limit of 1000 content objects that can be imported at once. If you want to import more than 1000 content objects, then be sure to split the import into batches of 1000 objects or less.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ContentSyncDefinition**](ContentSyncDefinition.md)| The content to import. | 
  **folderId** | **string**| The identifier of the folder to import into. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format. | 
 **optional** | ***ContentManagementApiBeginAsyncImportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiBeginAsyncImportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 
 **overwrite** | **optional.**| Set this to \&quot;true\&quot; to overwrite a content item if the name already exists. | [default to false]

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncDeleteStatus**
> AsyncJobStatus GetAsyncDeleteStatus(ctx, id, jobId, optional)
Content deletion job status.

Get the status of an asynchronous content deletion job request for the given job identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the content to delete. | 
  **jobId** | **string**| The identifier of the asynchronous job. | 
 **optional** | ***ContentManagementApiGetAsyncDeleteStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiGetAsyncDeleteStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncExportResult**
> ContentSyncDefinition GetAsyncExportResult(ctx, contentId, jobId, optional)
Content export job result.

Get results from content export job for the given job identifier. The results from this export are incompatible with the Library import feature in the Sumo user interface.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentId** | **string**| The identifier of the exported content item. | 
  **jobId** | **string**| The identifier of the asynchronous job. | 
 **optional** | ***ContentManagementApiGetAsyncExportResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiGetAsyncExportResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**ContentSyncDefinition**](ContentSyncDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncExportStatus**
> AsyncJobStatus GetAsyncExportStatus(ctx, contentId, jobId, optional)
Content export job status.

Get the status of an asynchronous content export request for the given job identifier. On success, use the [getExportResult](#operation/getAsyncExportResult) endpoint to get the result of the export job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentId** | **string**| The identifier of the exported content item. | 
  **jobId** | **string**| The identifier of the asynchronous export job. | 
 **optional** | ***ContentManagementApiGetAsyncExportStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiGetAsyncExportStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncImportStatus**
> AsyncJobStatus GetAsyncImportStatus(ctx, folderId, jobId, optional)
Content import job status.

Get the status of a content import job for the given job identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The identifier of the folder to import into. | 
  **jobId** | **string**| The identifier of the import request. | 
 **optional** | ***ContentManagementApiGetAsyncImportStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiGetAsyncImportStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetItemByPath**
> Content GetItemByPath(ctx, path)
Get content item by path.

Get a content item corresponding to the given path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Path of the content item to retrieve. | 

### Return type

[**Content**](Content.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPathById**
> ContentPath GetPathById(ctx, contentId)
Get path of an item.

Get full path of a content item with the given identifier. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentId** | **string**| Identifier of the content item to get the path. | 

### Return type

[**ContentPath**](ContentPath.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveItem**
> MoveItem(ctx, destinationFolderId, id, optional)
Move an item.

Moves an item from its current location to another folder. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinationFolderId** | **string**| Identifier of the destination folder. | 
  **id** | **string**| Identifier of the item the user wants to move. | 
 **optional** | ***ContentManagementApiMoveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentManagementApiMoveItemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

