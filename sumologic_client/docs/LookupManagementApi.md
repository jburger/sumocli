# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTable**](LookupManagementApi.md#CreateTable) | **Post** /v1/lookupTables | Create a lookup table.
[**DeleteTable**](LookupManagementApi.md#DeleteTable) | **Delete** /v1/lookupTables/{id} | Delete a lookup table.
[**DeleteTableRow**](LookupManagementApi.md#DeleteTableRow) | **Put** /v1/lookupTables/{id}/deleteTableRow | Delete a lookup table row.
[**LookupTableById**](LookupManagementApi.md#LookupTableById) | **Get** /v1/lookupTables/{id} | Get a lookup table.
[**RequestJobStatus**](LookupManagementApi.md#RequestJobStatus) | **Get** /v1/lookupTables/jobs/{jobId}/status | Get the status of an async job.
[**TruncateTable**](LookupManagementApi.md#TruncateTable) | **Post** /v1/lookupTables/{id}/truncate | Empty a lookup table.
[**UpdateTable**](LookupManagementApi.md#UpdateTable) | **Put** /v1/lookupTables/{id} | Edit a lookup table.
[**UpdateTableRow**](LookupManagementApi.md#UpdateTableRow) | **Put** /v1/lookupTables/{id}/row | Insert or Update a lookup table row.
[**UploadFile**](LookupManagementApi.md#UploadFile) | **Post** /v1/lookupTables/{id}/upload | Upload a CSV file.

# **CreateTable**
> LookupTable CreateTable(ctx, body)
Create a lookup table.

Create a new lookup table by providing a schema and specifying its configuration. Providing parentFolderId  is mandatory. Use the [getItemByPath](#operation/getItemByPath) endpoint to get content id of a path. Please check [Content management API](#tag/contentManagement) and [Folder management API](#tag/folderManagement) for all available options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LookupTableDefinition**](LookupTableDefinition.md)| The schema and configuration for the lookup table. | 

### Return type

[**LookupTable**](LookupTable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTable**
> DeleteTable(ctx, id)
Delete a lookup table.

Delete a lookup table completely.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the lookup table. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTableRow**
> DeleteTableRow(ctx, body, id)
Delete a lookup table row.

Delete a row from lookup table by giving primary key. The complete set of primary key fields of the lookup table should be provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RowDeleteDefinition**](RowDeleteDefinition.md)| Lookup table row delete definition. | 
  **id** | **string**| Identifier of the lookup table. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LookupTableById**
> LookupTable LookupTableById(ctx, id)
Get a lookup table.

Get a lookup table for the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the lookup table. | 

### Return type

[**LookupTable**](LookupTable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestJobStatus**
> LookupAsyncJobStatus RequestJobStatus(ctx, jobId)
Get the status of an async job.

Retrieve the status of a previously made request. If the request was successful, the status of the response object will be `Success`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| An identifier returned in response to an asynchronous request. | 

### Return type

[**LookupAsyncJobStatus**](LookupAsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TruncateTable**
> LookupRequestToken TruncateTable(ctx, id)
Empty a lookup table.

Delete all data from a lookup table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the table to clear. | 

### Return type

[**LookupRequestToken**](LookupRequestToken.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTable**
> LookupTable UpdateTable(ctx, body, id)
Edit a lookup table.

Edit the lookup table data. All the fields are mandatory in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LookupUpdateDefinition**](LookupUpdateDefinition.md)| The configuration changes for the lookup table. | 
  **id** | **string**| Identifier of the lookup table. | 

### Return type

[**LookupTable**](LookupTable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTableRow**
> UpdateTableRow(ctx, body, id)
Insert or Update a lookup table row.

Insert or update a row of a lookup table with the given identifier. A new row is inserted if the primary key does not exist already, otherwise the existing row with the specified primary key is updated. All the fields of the lookup table are required and will be updated to the given values. In case a field is not specified then it will be assumed to be set to null. If the table size exceeds the maximum limit of 100MB then based on the size limit action of the table the update will be processed or discarded.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RowUpdateDefinition**](RowUpdateDefinition.md)| Lookup table row update definition. | 
  **id** | **string**| Identifier of the lookup table. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFile**
> LookupRequestToken UploadFile(ctx, file, id, optional)
Upload a CSV file.

Create a request to populate a lookup table with a CSV file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 
  **id** | **string**| Identifier of the lookup table to populate. | 
 **optional** | ***LookupManagementApiUploadFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LookupManagementApiUploadFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **merge** | **optional.**| This indicates whether the file contents will be merged with existing data in the lookup table or not. If this is true then data with the same primary keys will be updated while the rest of the rows will be appended. By default, merge is false. The response includes a request identifier that you need to use in the [Request Status API](#operation/requestStatus) to track the status of the upload request. | [default to false]
 **fileEncoding** | **optional.**| File encoding of file being uploaded. | [default to UTF-8]

### Return type

[**LookupRequestToken**](LookupRequestToken.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

