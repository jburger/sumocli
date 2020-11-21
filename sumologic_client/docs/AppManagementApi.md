# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApp**](AppManagementApi.md#GetApp) | **Get** /v1/apps/{uuid} | Get an app by UUID.
[**GetAsyncInstallStatus**](AppManagementApi.md#GetAsyncInstallStatus) | **Get** /v1/apps/install/{jobId}/status | App install job status.
[**InstallApp**](AppManagementApi.md#InstallApp) | **Post** /v1/apps/{uuid}/install | Install an app by UUID.
[**ListApps**](AppManagementApi.md#ListApps) | **Get** /v1/apps | List available apps.

# **GetApp**
> App GetApp(ctx, uuid)
Get an app by UUID.

Gets the app with the given universally unique identifier (UUID).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)| The identifier of the app to retrieve. | 

### Return type

[**App**](App.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncInstallStatus**
> AsyncJobStatus GetAsyncInstallStatus(ctx, jobId)
App install job status.

Get the status of an asynchronous app install request for the given job identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The identifier of the asynchronous install job. | 

### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallApp**
> BeginAsyncJobResponse InstallApp(ctx, body, uuid)
Install an app by UUID.

Installs the app with given UUID in the folder specified using destinationFolderId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AppInstallRequest**](AppInstallRequest.md)|  | 
  **uuid** | [**string**](.md)| UUID of the app to install. | 

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApps**
> ListAppsResult ListApps(ctx, )
List available apps.

Lists all available apps from the App Catalog.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListAppsResult**](ListAppsResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

