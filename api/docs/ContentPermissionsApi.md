# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContentPermissions**](ContentPermissionsApi.md#AddContentPermissions) | **Put** /v2/content/{id}/permissions/add | Add permissions to a content item.
[**GetContentPermissions**](ContentPermissionsApi.md#GetContentPermissions) | **Get** /v2/content/{id}/permissions | Get permissions of a content item
[**RemoveContentPermissions**](ContentPermissionsApi.md#RemoveContentPermissions) | **Put** /v2/content/{id}/permissions/remove | Remove permissions from a content item.

# **AddContentPermissions**
> ContentPermissionResult AddContentPermissions(ctx, body, id, optional)
Add permissions to a content item.

Add permissions to a content item with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ContentPermissionUpdateRequest**](ContentPermissionUpdateRequest.md)| New permissions to add to the content item with the given identifier. | 
  **id** | **string**| The identifier of the content item. | 
 **optional** | ***ContentPermissionsApiAddContentPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentPermissionsApiAddContentPermissionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**ContentPermissionResult**](ContentPermissionResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentPermissions**
> ContentPermissionResult GetContentPermissions(ctx, id, optional)
Get permissions of a content item

Returns content permissions of a content item with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The identifier of the content item. | 
 **optional** | ***ContentPermissionsApiGetContentPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentPermissionsApiGetContentPermissionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **explicitOnly** | **optional.Bool**| There are two permission types: explicit and implicit. Permissions specifically assigned to the content item are explicit. Permissions derived from a parent content item, like a folder are implicit. To return only explicit permissions set this to true. | [default to false]
 **isAdminMode** | **optional.String**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**ContentPermissionResult**](ContentPermissionResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveContentPermissions**
> ContentPermissionResult RemoveContentPermissions(ctx, body, id, optional)
Remove permissions from a content item.

Remove permissions from a content item with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ContentPermissionUpdateRequest**](ContentPermissionUpdateRequest.md)| Permissions to remove from a content item with the given identifier. | 
  **id** | **string**| The identifier of the content item. | 
 **optional** | ***ContentPermissionsApiRemoveContentPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentPermissionsApiRemoveContentPermissionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **optional.**| Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**ContentPermissionResult**](ContentPermissionResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

