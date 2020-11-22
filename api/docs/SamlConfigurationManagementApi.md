# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAllowlistedUser**](SamlConfigurationManagementApi.md#CreateAllowlistedUser) | **Post** /v1/saml/allowlistedUsers/{userId} | Allowlist a user.
[**CreateIdentityProvider**](SamlConfigurationManagementApi.md#CreateIdentityProvider) | **Post** /v1/saml/identityProviders | Create a new SAML configuration.
[**CreateWhitelistedUser**](SamlConfigurationManagementApi.md#CreateWhitelistedUser) | **Post** /v1/saml/whitelistedUsers/{userId} | Whitelist a user.
[**DeleteAllowlistedUser**](SamlConfigurationManagementApi.md#DeleteAllowlistedUser) | **Delete** /v1/saml/allowlistedUsers/{userId} | Remove a allowlisted user.
[**DeleteIdentityProvider**](SamlConfigurationManagementApi.md#DeleteIdentityProvider) | **Delete** /v1/saml/identityProviders/{id} | Delete a SAML configuration.
[**DeleteWhitelistedUser**](SamlConfigurationManagementApi.md#DeleteWhitelistedUser) | **Delete** /v1/saml/whitelistedUsers/{userId} | Remove a whitelisted user.
[**DisableSamlLockdown**](SamlConfigurationManagementApi.md#DisableSamlLockdown) | **Post** /v1/saml/lockdown/disable | Disable SAML lockdown.
[**EnableSamlLockdown**](SamlConfigurationManagementApi.md#EnableSamlLockdown) | **Post** /v1/saml/lockdown/enable | Require SAML for sign-in.
[**GetAllowlistedUsers**](SamlConfigurationManagementApi.md#GetAllowlistedUsers) | **Get** /v1/saml/allowlistedUsers | Get list of allowlisted users.
[**GetIdentityProviders**](SamlConfigurationManagementApi.md#GetIdentityProviders) | **Get** /v1/saml/identityProviders | Get a list of SAML configurations.
[**GetWhitelistedUsers**](SamlConfigurationManagementApi.md#GetWhitelistedUsers) | **Get** /v1/saml/whitelistedUsers | Get list of whitelisted users.
[**UpdateIdentityProvider**](SamlConfigurationManagementApi.md#UpdateIdentityProvider) | **Put** /v1/saml/identityProviders/{id} | Update a SAML configuration.

# **CreateAllowlistedUser**
> AllowlistedUserResult CreateAllowlistedUser(ctx, userId)
Allowlist a user.

Allowlist a user from SAML lockdown allowing them to sign in using a password in addition to SAML.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Identifier of the user. | 

### Return type

[**AllowlistedUserResult**](AllowlistedUserResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIdentityProvider**
> SamlIdentityProvider CreateIdentityProvider(ctx, body)
Create a new SAML configuration.

Create a new SAML configuration in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SamlIdentityProviderRequest**](SamlIdentityProviderRequest.md)| The configuration of the SAML identity provider. | 

### Return type

[**SamlIdentityProvider**](SamlIdentityProvider.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWhitelistedUser**
> AllowlistedUserResult CreateWhitelistedUser(ctx, userId)
Whitelist a user.

Whitelist a user from SAML lockdown allowing them to sign in using a password in addition to SAML.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Identifier of the user. | 

### Return type

[**AllowlistedUserResult**](AllowlistedUserResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllowlistedUser**
> DeleteAllowlistedUser(ctx, userId)
Remove a allowlisted user.

Remove a allowlisted user requiring them to sign in using SAML.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Identifier of user that will no longer be allowlisted from SAML Lockdown. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdentityProvider**
> DeleteIdentityProvider(ctx, id)
Delete a SAML configuration.

Delete a SAML configuration with the given identifier from the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the SAML configuration to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWhitelistedUser**
> DeleteWhitelistedUser(ctx, userId)
Remove a whitelisted user.

Remove a whitelisted user requiring them to sign in using SAML.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Identifier of user that will no longer be whitelisted from SAML Lockdown. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableSamlLockdown**
> DisableSamlLockdown(ctx, )
Disable SAML lockdown.

Disable SAML lockdown for the organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableSamlLockdown**
> EnableSamlLockdown(ctx, )
Require SAML for sign-in.

Enabling SAML lockdown requires users to sign in using SAML preventing them from logging in with an email and password.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllowlistedUsers**
> []AllowlistedUserResult GetAllowlistedUsers(ctx, )
Get list of allowlisted users.

Get a list of allowlisted users.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AllowlistedUserResult**](AllowlistedUserResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityProviders**
> []SamlIdentityProvider GetIdentityProviders(ctx, )
Get a list of SAML configurations.

Get a list of all SAML configurations in the organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SamlIdentityProvider**](SamlIdentityProvider.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWhitelistedUsers**
> []AllowlistedUserResult GetWhitelistedUsers(ctx, )
Get list of whitelisted users.

Get a list of whitelisted users.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AllowlistedUserResult**](AllowlistedUserResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIdentityProvider**
> SamlIdentityProvider UpdateIdentityProvider(ctx, body, id)
Update a SAML configuration.

Update an existing SAML configuration in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SamlIdentityProviderRequest**](SamlIdentityProviderRequest.md)| Information to update in the SAML configuration. | 
  **id** | **string**| Identifier of the SAML configuration to update. | 

### Return type

[**SamlIdentityProvider**](SamlIdentityProvider.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

