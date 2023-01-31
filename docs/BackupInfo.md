# BackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **int64** |  | [readonly] 
**State** | [**StateCode**](StateCode.md) |  | [readonly] 
**FailureReason** | Pointer to **string** | Reason for failure if the state is &#39;FAILED&#39; | [optional] 

## Methods

### NewBackupInfo

`func NewBackupInfo(backupId int64, state StateCode, ) *BackupInfo`

NewBackupInfo instantiates a new BackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupInfoWithDefaults

`func NewBackupInfoWithDefaults() *BackupInfo`

NewBackupInfoWithDefaults instantiates a new BackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *BackupInfo) GetBackupId() int64`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *BackupInfo) GetBackupIdOk() (*int64, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *BackupInfo) SetBackupId(v int64)`

SetBackupId sets BackupId field to given value.


### GetState

`func (o *BackupInfo) GetState() StateCode`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupInfo) GetStateOk() (*StateCode, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupInfo) SetState(v StateCode)`

SetState sets State field to given value.


### GetFailureReason

`func (o *BackupInfo) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *BackupInfo) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *BackupInfo) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *BackupInfo) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


