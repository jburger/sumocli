# SamlIdentityProviderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpInitiatedLoginPath** | **string** | The identifier used to generate a unique URL for user login. | [optional] 
**ConfigurationName** | **string** | Name of the SSO policy or another name used to describe the policy internally. | [default to null]
**Issuer** | **string** | The unique URL assigned to the organization by the SAML Identity Provider. | [default to null]
**SpInitiatedLoginEnabled** | **bool** | True if Sumo Logic redirects users to your identity provider with a SAML AuthnRequest when signing in. | [optional] [default to false]
**AuthnRequestUrl** | **string** | The URL that the identity provider has assigned for Sumo Logic to submit SAML authentication requests to the identity provider. | [optional] 
**X509cert1** | **string** | The certificate is used to verify the signature in SAML assertions. | [default to null]
**X509cert2** | **string** | The backup certificate used to verify the signature in SAML assertions when x509cert1 expires. | [optional] 
**X509cert3** | **string** | The backup certificate used to verify the signature in SAML assertions when x509cert1 expires and x509cert2 is empty. | [optional] 
**OnDemandProvisioningEnabled** | [***OnDemandProvisioningInfo**](OnDemandProvisioningInfo.md) |  | [optional] [default to null]
**RolesAttribute** | **string** | The role that Sumo Logic will assign to users when they sign in. | [optional] 
**LogoutEnabled** | **bool** | True if users are redirected to a URL after signing out of Sumo Logic. | [optional] [default to false]
**LogoutUrl** | **string** | The URL that users will be redirected to after signing out of Sumo Logic. | [optional] 
**EmailAttribute** | **string** | The email address of the new user account. | [optional] 
**DebugMode** | **bool** | True if additional details are included when a user fails to sign in. | [optional] [default to false]
**SignAuthnRequest** | **bool** | True if Sumo Logic will send signed Authn requests to the identity provider. | [optional] [default to false]
**DisableRequestedAuthnContext** | **bool** | True if Sumo Logic will include the RequestedAuthnContext element of the SAML AuthnRequests it sends to the identity provider. | [optional] [default to false]
**IsRedirectBinding** | **bool** | True if the SAML binding is of HTTP Redirect type. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

