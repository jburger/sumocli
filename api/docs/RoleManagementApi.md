# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToUser**](RoleManagementApi.md#AssignRoleToUser) | **Put** /v1/roles/{roleId}/users/{userId} | Assign a role to a user.
[**CreateRole**](RoleManagementApi.md#CreateRole) | **Post** /v1/roles | Create a new role.
[**DeleteRole**](RoleManagementApi.md#DeleteRole) | **Delete** /v1/roles/{id} | Delete a role.
[**GetRole**](RoleManagementApi.md#GetRole) | **Get** /v1/roles/{id} | Get a role.
[**ListRoles**](RoleManagementApi.md#ListRoles) | **Get** /v1/roles | Get a list of roles.
[**RemoveRoleFromUser**](RoleManagementApi.md#RemoveRoleFromUser) | **Delete** /v1/roles/{roleId}/users/{userId} | Remove role from a user.
[**UpdateRole**](RoleManagementApi.md#UpdateRole) | **Put** /v1/roles/{id} | Update a role.

# **AssignRoleToUser**
> RoleModel AssignRoleToUser(ctx, roleId, userId)
Assign a role to a user.

Assign a role to a user in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**| Identifier of the role to assign. | 
  **userId** | **string**| Identifier of the user to assign the role to. | 

### Return type

[**RoleModel**](RoleModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRole**
> RoleModel CreateRole(ctx, body)
Create a new role.

Create a new role in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleDefinition**](CreateRoleDefinition.md)| Information about the new role. | 

### Return type

[**RoleModel**](RoleModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> DeleteRole(ctx, id)
Delete a role.

Delete a role with the given identifier from the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the role to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> RoleModel GetRole(ctx, id)
Get a role.

Get a role with the given identifier in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the role to fetch. | 

### Return type

[**RoleModel**](RoleModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRoles**
> ListRoleModelsResponse ListRoles(ctx, optional)
Get a list of roles.

Get a list of all the roles in the organization. The response is paginated with a default limit of 100 roles per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoleManagementApiListRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleManagementApiListRolesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of roles returned in the response. The number of roles returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 
 **sortBy** | **optional.String**| Sort the list of roles by the &#x60;name&#x60; field. | 
 **name** | **optional.String**| Only return roles matching the given name. | 

### Return type

[**ListRoleModelsResponse**](ListRoleModelsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoleFromUser**
> RemoveRoleFromUser(ctx, roleId, userId)
Remove role from a user.

Remove a role from a user in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**| Identifier of the role to delete. | 
  **userId** | **string**| Identifier of the user to remove the role from. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRole**
> RoleModel UpdateRole(ctx, body, id)
Update a role.

Update an existing role in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateRoleDefinition**](UpdateRoleDefinition.md)| Information to update about the role. | 
  **id** | **string**| Identifier of the role to update. | 

### Return type

[**RoleModel**](RoleModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

