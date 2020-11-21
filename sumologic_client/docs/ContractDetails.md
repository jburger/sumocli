# ContractDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | Organization identifier of the account. | [default to null]
**PlanType** | **string** | Plan name of the account. | [default to null]
**Entitlements** | [**[]Entitlements**](Entitlements.md) | List of the entitlements of the account. Entitlements of the account are the list of products subscribed by the user. | [default to null]
**SharedBuckets** | [**[]SharedBucket**](SharedBucket.md) | Contains list of buckets. Bucket means shared pool from which multiple entitlements can use capacity. | [optional] [default to null]
**ContractPeriod** | [***ContractPeriod**](ContractPeriod.md) |  | [default to null]
**CurrentBillingPeriod** | [***CurrentBillingPeriod**](CurrentBillingPeriod.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

