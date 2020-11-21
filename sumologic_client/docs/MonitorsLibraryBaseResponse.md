# MonitorsLibraryBaseResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the monitor or folder. | [default to null]
**Name** | **string** | Identifier of the monitor or folder. | [default to null]
**Description** | **string** | Description of the monitor or folder. | [default to null]
**Version** | **int64** | Version of the monitor or folder. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**CreatedBy** | **string** | Identifier of the user who created the resource. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Last modification timestamp in UTC. | [default to null]
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | [default to null]
**ParentId** | **string** | Identifier of the parent folder. | [default to null]
**ContentType** | **string** | Type of the content. Valid values:   1) Monitor   2) Folder | [default to null]
**Type_** | **string** | Type of the object model. | [default to null]
**IsSystem** | **bool** | System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can&#x27;t be updated. | [default to null]
**IsMutable** | **bool** | Immutable objects are \&quot;READ-ONLY\&quot;. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

