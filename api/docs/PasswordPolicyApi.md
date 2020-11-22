# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPasswordPolicy**](PasswordPolicyApi.md#GetPasswordPolicy) | **Get** /v1/passwordPolicy | Get the current password policy.
[**SetPasswordPolicy**](PasswordPolicyApi.md#SetPasswordPolicy) | **Put** /v1/passwordPolicy | Update password policy.

# **GetPasswordPolicy**
> PasswordPolicy GetPasswordPolicy(ctx, )
Get the current password policy.

Get the current password policy.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPasswordPolicy**
> PasswordPolicy SetPasswordPolicy(ctx, body)
Update password policy.

Update the current password policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PasswordPolicy**](PasswordPolicy.md)|  | 

### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

