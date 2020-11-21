# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMonitorUsageInfo**](MonitorsLibraryManagementApi.md#GetMonitorUsageInfo) | **Get** /v1/monitors/usageInfo | Usage info of monitors.
[**GetMonitorsFullPath**](MonitorsLibraryManagementApi.md#GetMonitorsFullPath) | **Get** /v1/monitors/{id}/path | Get the path of a monitor or folder.
[**GetMonitorsLibraryRoot**](MonitorsLibraryManagementApi.md#GetMonitorsLibraryRoot) | **Get** /v1/monitors/root | Get the root monitors folder.
[**MonitorsCopy**](MonitorsLibraryManagementApi.md#MonitorsCopy) | **Post** /v1/monitors/{id}/copy | Copy a monitor or folder.
[**MonitorsCreate**](MonitorsLibraryManagementApi.md#MonitorsCreate) | **Post** /v1/monitors | Create a monitor or folder. 
[**MonitorsDeleteById**](MonitorsLibraryManagementApi.md#MonitorsDeleteById) | **Delete** /v1/monitors/{id} | Delete a monitor or folder. 
[**MonitorsDeleteByIds**](MonitorsLibraryManagementApi.md#MonitorsDeleteByIds) | **Delete** /v1/monitors | Bulk delete a monitor or folder. 
[**MonitorsExportItem**](MonitorsLibraryManagementApi.md#MonitorsExportItem) | **Get** /v1/monitors/{id}/export | Export a monitor or folder.
[**MonitorsGetByPath**](MonitorsLibraryManagementApi.md#MonitorsGetByPath) | **Get** /v1/monitors/path | Read a monitor or folder by its path.
[**MonitorsImportItem**](MonitorsLibraryManagementApi.md#MonitorsImportItem) | **Post** /v1/monitors/{parentId}/import | Import a monitor or folder.
[**MonitorsMove**](MonitorsLibraryManagementApi.md#MonitorsMove) | **Post** /v1/monitors/{id}/move | Move a monitor or folder.
[**MonitorsReadById**](MonitorsLibraryManagementApi.md#MonitorsReadById) | **Get** /v1/monitors/{id} | Get a monitor or folder.
[**MonitorsReadByIds**](MonitorsLibraryManagementApi.md#MonitorsReadByIds) | **Get** /v1/monitors | Bulk read a monitor or folder.
[**MonitorsSearch**](MonitorsLibraryManagementApi.md#MonitorsSearch) | **Get** /v1/monitors/search | Search for a monitor or folder.
[**MonitorsUpdateById**](MonitorsLibraryManagementApi.md#MonitorsUpdateById) | **Put** /v1/monitors/{id} | Update a monitor or folder. 

# **GetMonitorUsageInfo**
> MonitorUsageInfo GetMonitorUsageInfo(ctx, )
Usage info of monitors.

Get the current number and the allowed number of log and metrics monitors.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonitorUsageInfo**](MonitorUsageInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitorsFullPath**
> Path GetMonitorsFullPath(ctx, id)
Get the path of a monitor or folder.

Get the full path of the monitor or folder in the monitors library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the monitor or folder. | 

### Return type

[**Path**](Path.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitorsLibraryRoot**
> MonitorsLibraryFolderResponse GetMonitorsLibraryRoot(ctx, )
Get the root monitors folder.

Get the root folder in the monitors library.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonitorsLibraryFolderResponse**](MonitorsLibraryFolderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsCopy**
> MonitorsLibraryBaseResponse MonitorsCopy(ctx, body, id)
Copy a monitor or folder.

Copy a monitor or folder in the monitors library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ContentCopyParams**](ContentCopyParams.md)| Fields include:
  1) Identifier of the parent folder to copy to.
  2) Optionally provide a new name.
  3) Optionally provide a new description.
  4) Optionally set to true if you want to copy and preserve the locked status. Requires &#x60;LockMonitors&#x60; capability. | 
  **id** | **string**| Identifier of the monitor or folder to copy. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsCreate**
> MonitorsLibraryBaseResponse MonitorsCreate(ctx, body, parentId)
Create a monitor or folder. 

Create a monitor or folder in the monitors library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MonitorsLibraryBase**](MonitorsLibraryBase.md)| The monitor or folder to create. | 
  **parentId** | **string**| Identifier of the parent folder in which to create the monitor or folder. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsDeleteById**
> MonitorsDeleteById(ctx, id)
Delete a monitor or folder. 

Delete a monitor or folder from the monitors library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the monitor or folder to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsDeleteByIds**
> map[string]MonitorsLibraryBaseResponse MonitorsDeleteByIds(ctx, ids)
Bulk delete a monitor or folder. 

Bulk delete a monitor or folder by the given identifiers in the monitors library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]string**](string.md)| A comma-separated list of identifiers. | 

### Return type

[**map[string]MonitorsLibraryBaseResponse**](map.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsExportItem**
> MonitorsLibraryBaseExport MonitorsExportItem(ctx, id)
Export a monitor or folder.

Export a monitor or folder. If the given identifier is a folder, everything under the folder is exported recursively with folder as the root.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the monitor or folder to export. | 

### Return type

[**MonitorsLibraryBaseExport**](MonitorsLibraryBaseExport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsGetByPath**
> MonitorsLibraryBaseResponse MonitorsGetByPath(ctx, path)
Read a monitor or folder by its path.

Read a monitor or folder by its path in the monitors library structure.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The path of the monitor or folder. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsImportItem**
> MonitorsLibraryBaseResponse MonitorsImportItem(ctx, body, parentId)
Import a monitor or folder.

Import a monitor or folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MonitorsLibraryBaseExport**](MonitorsLibraryBaseExport.md)| The monitor or folder to be imported. | 
  **parentId** | **string**| Identifier of the parent folder in which to import the monitor or folder. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsMove**
> MonitorsLibraryBaseResponse MonitorsMove(ctx, id, parentId)
Move a monitor or folder.

Move a monitor or folder to a different location in the monitors library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the monitor or folder to move. | 
  **parentId** | **string**| Identifier of the parent folder to move the monitor or folder to. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsReadById**
> MonitorsLibraryBaseResponse MonitorsReadById(ctx, id)
Get a monitor or folder.

Get a monitor or folder from the monitors library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the monitor or folder to read. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsReadByIds**
> map[string]MonitorsLibraryBaseResponse MonitorsReadByIds(ctx, ids)
Bulk read a monitor or folder.

Bulk read a monitor or folder by the given identifiers from the monitors library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]string**](string.md)| A comma-separated list of identifiers. | 

### Return type

[**map[string]MonitorsLibraryBaseResponse**](map.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsSearch**
> []MonitorsLibraryItemWithPath MonitorsSearch(ctx, query, optional)
Search for a monitor or folder.

Search for a monitor or folder in the monitors library structure.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| The search query to find monitor or folder. Below is the list of different filters with examples:   - **createdBy** : Filter by the user&#x27;s identifier who created the content. Example: &#x60;createdBy:000000000000968B&#x60;.   - **createdBefore** : Filter by the content objects created before the given timestamp(in milliseconds). Example: &#x60;createdBefore:1457997222&#x60;.   - **createdAfter** : Filter by the content objects created after the given timestamp(in milliseconds). Example: &#x60;createdAfter:1457997111&#x60;.   - **modifiedBefore** : Filter by the content objects modified before the given timestamp(in milliseconds). Example: &#x60;modifiedBefore:1457997222&#x60;.   - **modifiedAfter** : Filter by the content objects modified after the given timestamp(in milliseconds). Example: &#x60;modifiedAfter:1457997111&#x60;.   - **type** : Filter by the type of the content object. Example: &#x60;type:folder&#x60;.   - **monitorStatus** : Filter by the status of the monitor: Normal, Critical, Warning, MissingData, Disabled, AllTriggered. Example: &#x60;monitorStatus:Normal&#x60;.  You can also use multiple filters in one query. For example to search for all content objects created by user with identifier 000000000000968B with creation timestamp after 1457997222 containing the text Test, the query would look like:    &#x60;createdBy:000000000000968B createdAfter:1457997222 Test&#x60; | 
 **optional** | ***MonitorsLibraryManagementApiMonitorsSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorsLibraryManagementApiMonitorsSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of items you want in the response. | [default to 100]
 **offset** | **optional.Int32**| The position or row from where to start the search operation. | [default to 0]

### Return type

[**[]MonitorsLibraryItemWithPath**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsUpdateById**
> MonitorsLibraryBaseResponse MonitorsUpdateById(ctx, body, id)
Update a monitor or folder. 

Update a monitor or folder in the monitors library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MonitorsLibraryBaseUpdate**](MonitorsLibraryBaseUpdate.md)| The monitor or folder to update. The content version must match its latest version number in the monitors library. If the version does not match it will not be updated. | 
  **id** | **string**| Identifier of the monitor or folder to update. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

