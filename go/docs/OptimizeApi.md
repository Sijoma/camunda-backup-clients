# \OptimizeApi

All URIs are relative to *http://localhost:9600/actuator/backups*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupIdDelete**](OptimizeApi.md#BackupIdDelete) | **Delete** /{backupId} | Delete a backup
[**BackupIdGet**](OptimizeApi.md#BackupIdGet) | **Get** /{backupId} | Get information of a backup
[**RootGet**](OptimizeApi.md#RootGet) | **Get** / | Lists all available backups
[**RootPost**](OptimizeApi.md#RootPost) | **Post** / | Takes a backup



## BackupIdDelete

> BackupIdDelete(ctx, backupId).Execute()

Delete a backup



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    backupId := int64(789) // int64 | Id of the backup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptimizeApi.BackupIdDelete(context.Background(), backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptimizeApi.BackupIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **int64** | Id of the backup | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupIdGet

> BackupInfo BackupIdGet(ctx, backupId).Execute()

Get information of a backup



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    backupId := int64(789) // int64 | Id of the backup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptimizeApi.BackupIdGet(context.Background(), backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptimizeApi.BackupIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupIdGet`: BackupInfo
    fmt.Fprintf(os.Stdout, "Response from `OptimizeApi.BackupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **int64** | Id of the backup | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupInfo**](BackupInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RootGet

> []BackupInfo RootGet(ctx).Execute()

Lists all available backups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptimizeApi.RootGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptimizeApi.RootGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RootGet`: []BackupInfo
    fmt.Fprintf(os.Stdout, "Response from `OptimizeApi.RootGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRootGetRequest struct via the builder pattern


### Return type

[**[]BackupInfo**](BackupInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RootPost

> TakeBackupResponse RootPost(ctx).TakeBackupRequest(takeBackupRequest).Execute()

Takes a backup



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    takeBackupRequest := *openapiclient.NewTakeBackupRequest(int64(123)) // TakeBackupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptimizeApi.RootPost(context.Background()).TakeBackupRequest(takeBackupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptimizeApi.RootPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RootPost`: TakeBackupResponse
    fmt.Fprintf(os.Stdout, "Response from `OptimizeApi.RootPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRootPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **takeBackupRequest** | [**TakeBackupRequest**](TakeBackupRequest.md) |  | 

### Return type

[**TakeBackupResponse**](TakeBackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

