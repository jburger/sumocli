# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserManagementApi.md#CreateUser) | **Post** /v1/users | Create a new user.
[**DeleteUser**](UserManagementApi.md#DeleteUser) | **Delete** /v1/users/{id} | Delete a user.
[**DisableMfa**](UserManagementApi.md#DisableMfa) | **Put** /v1/users/{id}/mfa/disable | Disable MFA for user.
[**GetUser**](UserManagementApi.md#GetUser) | **Get** /v1/users/{id} | Get a user.
[**ListUsers**](UserManagementApi.md#ListUsers) | **Get** /v1/users | Get a list of users.
[**RequestChangeEmail**](UserManagementApi.md#RequestChangeEmail) | **Post** /v1/users/{id}/email/requestChange | Change email address.
[**ResetPassword**](UserManagementApi.md#ResetPassword) | **Post** /v1/users/{id}/password/reset | Reset password.
[**UnlockUser**](UserManagementApi.md#UnlockUser) | **Post** /v1/users/{id}/unlock | Unlock a user.
[**UpdateUser**](UserManagementApi.md#UpdateUser) | **Put** /v1/users/{id} | Update a user.

# **CreateUser**
> UserModel CreateUser(ctx, body)
Create a new user.

Create a new user in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserDefinition**](CreateUserDefinition.md)| Information about the new user. | 

### Return type

[**UserModel**](UserModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, id, optional)
Delete a user.

Delete a user with the given identifier from the organization and transfer their content to the user with the identifier specified in \"transferTo\". If \"transferTo\" is not specified the contents are deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the user to delete. | 
 **optional** | ***UserManagementApiDeleteUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserManagementApiDeleteUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferTo** | **optional.String**| Identifier of the user to receive the transfer of content from the deleted user. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableMfa**
> DisableMfa(ctx, body, id)
Disable MFA for user.

Disable multi-factor authentication for given user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DisableMfaRequest**](DisableMfaRequest.md)| Email and Password of the user to disable MFA for. | 
  **id** | **string**| Identifier of the user to disable MFA for. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> UserModel GetUser(ctx, id)
Get a user.

Get a user with the given identifier from the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of user to return. | 

### Return type

[**UserModel**](UserModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsers**
> ListUserModelsResponse ListUsers(ctx, optional)
Get a list of users.

Get a list of all users in the organization. The response is paginated with a default limit of 100 users per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserManagementApiListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserManagementApiListUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of users returned in the response. The number of users returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 
 **sortBy** | **optional.String**| Sort the list of users by the &#x60;firstName&#x60;, &#x60;lastName&#x60;, or &#x60;email&#x60; field. | 
 **email** | **optional.String**| Find user with the given email address. | 

### Return type

[**ListUserModelsResponse**](ListUserModelsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestChangeEmail**
> RequestChangeEmail(ctx, body, id)
Change email address.

An email with an activation link is sent to the user’s new email address. The user must click the link in the email within seven days to complete the email address change, or the link will expire.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeEmailRequest**](ChangeEmailRequest.md)| New email address of the user. | 
  **id** | **string**| Identifier of the user to change email address. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPassword**
> ResetPassword(ctx, id)
Reset password.

Reset a user's password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the user to reset password. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlockUser**
> UnlockUser(ctx, id)
Unlock a user.

Unlock another user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the user that needs to be unlocked. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UserModel UpdateUser(ctx, body, id)
Update a user.

Update an existing user in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserDefinition**](UpdateUserDefinition.md)| Information to update about the user. | 
  **id** | **string**| Identifier of the user to update. | 

### Return type

[**UserModel**](UserModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

