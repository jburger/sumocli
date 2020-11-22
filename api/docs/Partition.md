# Partition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**CreatedBy** | **string** | Identifier of the user who created the resource. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Last modification timestamp in UTC. | [default to null]
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | [default to null]
**Name** | **string** | The name of the partition. | [default to null]
**RoutingExpression** | **string** | The query that defines the data to be included in the partition. | [default to null]
**DataForwardingId** | **string** | An optional ID of a data forwarding configuration to be used by the partition. | [optional] [default to null]
**AnalyticsTier** | **string** | The Data Tier where the data in the partition will reside. Possible values are:               1. &#x60;continuous&#x60;               2. &#x60;frequent&#x60;               3. &#x60;infrequent&#x60;               4. &#x60;security&#x60; Note: The \&quot;infrequent\&quot; and \&quot;frequent\&quot; tiers are only to available Cloud Flex Credits Enterprise Suite accounts. Security tier is in private beta and is not available until given access. To participate in the beta program contact your Sumo Logic account representative. The terms for data tiers, \&quot;basic\&quot;, \&quot;enhanced\&quot;, \&quot;cold\&quot;, will be deprecated soon, and replaced by the terms, \&quot;continuous\&quot;, \&quot;infrequent\&quot;, \&quot;frequent\&quot;, respectively. Going forward, use the new terms. | [optional] [default to continuous]
**RetentionPeriod** | **int32** | The number of days to retain data in the partition, or -1 to use the default value for your account.  Only relevant if your account has variable retention enabled. | [optional] [default to -1]
**IsCompliant** | **bool** | Whether the partition is compliant or not. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition is marked compliant. A partition once marked compliant, cannot be marked non-compliant later. | [optional] [default to false]
**Id** | **string** | Unique identifier for the partition. | [default to null]
**TotalBytes** | **int64** | Size of data in partition in bytes. | [default to null]
**IsActive** | **bool** | This has the value &#x60;true&#x60; if the partition is active and &#x60;false&#x60; if it has been decommissioned. | [optional] [default to null]
**NewRetentionPeriod** | **int32** | If the retentionPeriod is scheduled to be updated in the future (i.e., if retentionPeriod is previously reduced with value of reduceRetentionPeriodImmediately as false), this property gives the future value of retentionPeriod while retentionPeriod gives the current value. retentionPeriod will take up the value of newRetentionPeriod after the scheduled time. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

