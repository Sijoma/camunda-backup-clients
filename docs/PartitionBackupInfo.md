# PartitionBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartitionId** | **int32** |  | [readonly] 
**State** | [**StateCode**](StateCode.md) |  | [readonly] 
**FailureReason** | Pointer to **string** | Failure reason if stats is &#39;FAILED&#39; | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp at which the backup was started on this partition. | [optional] [readonly] 
**LastUpdatedAt** | Pointer to **time.Time** | The timestamp at which the backup was last updated on this partition, e.g. changed state from IN_PROGRESS to COMPLETED.  | [optional] [readonly] 
**SnapshotId** | Pointer to **string** | The ID of the snapshot which is included in this backup. | [optional] [readonly] 
**CheckpointPosition** | Pointer to **int64** | The position of the checkpoint for this backup. | [optional] [readonly] 
**BrokerId** | Pointer to **int32** | The ID of the broker from which the backup was taken for this partition. | [optional] [readonly] 
**BrokerVersion** | Pointer to **string** | The version of the broker from which the backup was taken for this partition. | [optional] [readonly] 

## Methods

### NewPartitionBackupInfo

`func NewPartitionBackupInfo(partitionId int32, state StateCode, ) *PartitionBackupInfo`

NewPartitionBackupInfo instantiates a new PartitionBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionBackupInfoWithDefaults

`func NewPartitionBackupInfoWithDefaults() *PartitionBackupInfo`

NewPartitionBackupInfoWithDefaults instantiates a new PartitionBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitionId

`func (o *PartitionBackupInfo) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *PartitionBackupInfo) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *PartitionBackupInfo) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetState

`func (o *PartitionBackupInfo) GetState() StateCode`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PartitionBackupInfo) GetStateOk() (*StateCode, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PartitionBackupInfo) SetState(v StateCode)`

SetState sets State field to given value.


### GetFailureReason

`func (o *PartitionBackupInfo) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *PartitionBackupInfo) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *PartitionBackupInfo) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *PartitionBackupInfo) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PartitionBackupInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PartitionBackupInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PartitionBackupInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PartitionBackupInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *PartitionBackupInfo) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *PartitionBackupInfo) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *PartitionBackupInfo) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *PartitionBackupInfo) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetSnapshotId

`func (o *PartitionBackupInfo) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *PartitionBackupInfo) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *PartitionBackupInfo) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *PartitionBackupInfo) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetCheckpointPosition

`func (o *PartitionBackupInfo) GetCheckpointPosition() int64`

GetCheckpointPosition returns the CheckpointPosition field if non-nil, zero value otherwise.

### GetCheckpointPositionOk

`func (o *PartitionBackupInfo) GetCheckpointPositionOk() (*int64, bool)`

GetCheckpointPositionOk returns a tuple with the CheckpointPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointPosition

`func (o *PartitionBackupInfo) SetCheckpointPosition(v int64)`

SetCheckpointPosition sets CheckpointPosition field to given value.

### HasCheckpointPosition

`func (o *PartitionBackupInfo) HasCheckpointPosition() bool`

HasCheckpointPosition returns a boolean if a field has been set.

### GetBrokerId

`func (o *PartitionBackupInfo) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *PartitionBackupInfo) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *PartitionBackupInfo) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *PartitionBackupInfo) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetBrokerVersion

`func (o *PartitionBackupInfo) GetBrokerVersion() string`

GetBrokerVersion returns the BrokerVersion field if non-nil, zero value otherwise.

### GetBrokerVersionOk

`func (o *PartitionBackupInfo) GetBrokerVersionOk() (*string, bool)`

GetBrokerVersionOk returns a tuple with the BrokerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerVersion

`func (o *PartitionBackupInfo) SetBrokerVersion(v string)`

SetBrokerVersion sets BrokerVersion field to given value.

### HasBrokerVersion

`func (o *PartitionBackupInfo) HasBrokerVersion() bool`

HasBrokerVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


