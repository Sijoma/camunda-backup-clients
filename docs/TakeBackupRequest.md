# TakeBackupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **int64** | The ID of the backup to be taken | 

## Methods

### NewTakeBackupRequest

`func NewTakeBackupRequest(backupId int64, ) *TakeBackupRequest`

NewTakeBackupRequest instantiates a new TakeBackupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTakeBackupRequestWithDefaults

`func NewTakeBackupRequestWithDefaults() *TakeBackupRequest`

NewTakeBackupRequestWithDefaults instantiates a new TakeBackupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *TakeBackupRequest) GetBackupId() int64`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *TakeBackupRequest) GetBackupIdOk() (*int64, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *TakeBackupRequest) SetBackupId(v int64)`

SetBackupId sets BackupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


