# RoleModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**CreatedBy** | **string** | Identifier of the user who created the resource. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Last modification timestamp in UTC. | [default to null]
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | [default to null]
**Name** | **string** | Name of the role. | [default to null]
**Description** | **string** | Description of the role. | [optional] [default to null]
**FilterPredicate** | **string** | A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using &#x27;!_sourceCategory&#x3D;billing&#x27; as a filter predicate will prevent users assigned to the role from viewing logs from the source category named &#x27;billing&#x27;. | [optional] [default to null]
**Users** | **[]string** | List of user identifiers to assign the role to. | [optional] [default to null]
**Capabilities** | **[]string** | List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - dataVolumeIndex   - viewConnections   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - manageDataStreams  ### Entity management   - manageEntityTypeConfig  ### Metrics   - manageMonitors   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipWhitelisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardWhitelist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex | [optional] [default to null]
**AutofillDependencies** | **bool** | Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies. | [optional] [default to true]
**Id** | **string** | Unique identifier for the role. | [default to null]
**SystemDefined** | **bool** | Role is system or user defined. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

