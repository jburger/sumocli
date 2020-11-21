# CreatePartitionDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the partition. | [default to null]
**RoutingExpression** | **string** | The query that defines the data to be included in the partition. | [default to null]
**DataForwardingId** | **string** | An optional ID of a data forwarding configuration to be used by the partition. | [optional] [default to null]
**AnalyticsTier** | **string** | The Data Tier where the data in the partition will reside. Possible values are:               1. &#x60;continuous&#x60;               2. &#x60;frequent&#x60;               3. &#x60;infrequent&#x60;               4. &#x60;security&#x60; Note: The \&quot;infrequent\&quot; and \&quot;frequent\&quot; tiers are only to available Cloud Flex Credits Enterprise Suite accounts. Security tier is in private beta and is not available until given access. To participate in the beta program contact your Sumo Logic account representative. The terms for data tiers, \&quot;basic\&quot;, \&quot;enhanced\&quot;, \&quot;cold\&quot;, will be deprecated soon, and replaced by the terms, \&quot;continuous\&quot;, \&quot;infrequent\&quot;, \&quot;frequent\&quot;, respectively. Going forward, use the new terms. | [optional] [default to continuous]
**RetentionPeriod** | **int32** | The number of days to retain data in the partition, or -1 to use the default value for your account.  Only relevant if your account has variable retention enabled. | [optional] [default to -1]
**IsCompliant** | **bool** | Whether the partition is compliant or not. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition is marked compliant. A partition once marked compliant, cannot be marked non-compliant later. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

