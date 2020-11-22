# AccessKeyCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A name for the access key to be created. | [default to null]
**CorsHeaders** | **[]string** | An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don&#x27;t match any entry in the allowlist. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

