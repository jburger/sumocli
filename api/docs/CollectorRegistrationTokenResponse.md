# CollectorRegistrationTokenResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the token. | [default to null]
**Name** | **string** | Name of the token. | [default to null]
**Description** | **string** | Description of the token. | [default to null]
**Status** | **string** | Status of the token. Can be &#x60;Active&#x60;, or &#x60;Inactive&#x60;. | [default to null]
**Type_** | **string** | Type of the token. Valid values: 1) CollectorRegistrationTokenResponse | [default to null]
**Version** | **int64** | Version of the token. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**CreatedBy** | **string** | Identifier of the user who created the resource. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Last modification timestamp in UTC. | [default to null]
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | [default to null]
**EncodedTokenAndUrl** | **string** | The token and URL used to register the Collector as an encoded string. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

