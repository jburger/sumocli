# CreateScheduledViewDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The query that defines the data to be included in the scheduled view. | [default to null]
**IndexName** | **string** | Name of the index for the scheduled view. | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | Start timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**RetentionPeriod** | **int32** | The number of days to retain data in the scheduled view, or -1 to use the default value for your account.  Only relevant if your account has multi-retention. enabled. | [optional] [default to -1]
**DataForwardingId** | **string** | An optional ID of a data forwarding configuration to be used by the scheduled view. | [optional] [default to null]
**ParsingMode** | **string** | Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;AutoParse&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to Manual]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

