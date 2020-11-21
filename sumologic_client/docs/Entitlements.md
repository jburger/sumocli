# Entitlements

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractType** | **string** | Details of the contract type. &#x60;AnnualBucket&#x60; are contracts that buy and consume ingest on yearly basis. &#x60;Credits&#x60; are contracts that buy a single unit called credits for all our features. &#x60;DailyAverage&#x60; are contracts that buy and consume ingest on a monthly basis. | [default to null]
**EntitlementType** | **string** | Text denoting the type of entitlement. - &#x60;continuous&#x60; for Continuous Analytics, - &#x60;frequent&#x60; for Frequent Analytics, - &#x60;storage&#x60; for Total Storage, - &#x60;metrics&#x60; for Metrics. | [default to null]
**Label** | **string** | The label of an entitlement is the plan name displayed on the accounts page in our user interface. | [default to null]
**Capacity** | [***Capacity**](Capacity.md) |  | [default to null]
**Capacities** | [**[]Capacity**](Capacity.md) | Contains the capacities that were part of the contract. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

