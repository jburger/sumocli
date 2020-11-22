# CurrentPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | Unique identifier of the product in current plan. Valid values are: 1. &#x60;Free&#x60; 2. &#x60;Trial&#x60; 3. &#x60;Essentials&#x60; 4. &#x60;EnterpriseOps&#x60; 5. &#x60;EnterpriseSec&#x60; 6. &#x60;EnterpriseSuite&#x60;  | [default to null]
**PlanCost** | **int64** | Cost incurred for the current plan. | [default to null]
**BillingFrequency** | **string** | Billing frequency for the current plan. Valid values are: 1. &#x60;monthly&#x60;  | [default to null]
**Consumables** | [**[]Consumable**](Consumable.md) | Consumables in the current plan. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

