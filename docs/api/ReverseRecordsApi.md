# \ReverseRecordsApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ReverserecordsDelete**](ReverseRecordsApi.md#ReverserecordsDelete) | **Delete** /reverserecords/{reverserecordId} | Delete a reverse DNS record|
|[**ReverserecordsFindById**](ReverseRecordsApi.md#ReverserecordsFindById) | **Get** /reverserecords/{reverserecordId} | Retrieve a reverse DNS record|
|[**ReverserecordsGet**](ReverseRecordsApi.md#ReverserecordsGet) | **Get** /reverserecords | Retrieves existing reverse DNS records|
|[**ReverserecordsPost**](ReverseRecordsApi.md#ReverserecordsPost) | **Post** /reverserecords | Create a reverse DNS record|
|[**ReverserecordsPut**](ReverseRecordsApi.md#ReverserecordsPut) | **Put** /reverserecords/{reverserecordId} | Update a reverse DNS record|



## ReverserecordsDelete

```go
var result map[string]interface{} = ReverserecordsDelete(ctx, reverserecordId)
                      .Execute()
```

Delete a reverse DNS record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dns"
)

func main() {
    reverserecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the reverse DNS record.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.ReverseRecordsApi.ReverserecordsDelete(context.Background(), reverserecordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseRecordsApi.ReverserecordsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ReverserecordsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReverseRecordsApi.ReverserecordsDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**reverserecordId** | **string** | The ID (UUID) of the reverse DNS record. | |

### Other Parameters

Other parameters are passed through a pointer to an apiReverserecordsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ReverseRecordsApiService.ReverserecordsDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ReverseRecordsApiService.ReverserecordsDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ReverseRecordsApiService.ReverserecordsDelete": {
    "port": "8443",
},
})
```


## ReverserecordsFindById

```go
var result ReverseRecordRead = ReverserecordsFindById(ctx, reverserecordId)
                      .Execute()
```

Retrieve a reverse DNS record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dns"
)

func main() {
    reverserecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the reverse DNS record.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ReverseRecordsApi.ReverserecordsFindById(context.Background(), reverserecordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseRecordsApi.ReverserecordsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ReverserecordsFindById`: ReverseRecordRead
    fmt.Fprintf(os.Stdout, "Response from `ReverseRecordsApi.ReverserecordsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**reverserecordId** | **string** | The ID (UUID) of the reverse DNS record. | |

### Other Parameters

Other parameters are passed through a pointer to an apiReverserecordsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ReverseRecordRead**](../models/ReverseRecordRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ReverseRecordsApiService.ReverserecordsFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ReverseRecordsApiService.ReverserecordsFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ReverseRecordsApiService.ReverserecordsFindById": {
    "port": "8443",
},
})
```


## ReverserecordsGet

```go
var result ReverseRecordsReadList = ReverserecordsGet(ctx)
                      .FilterRecordIp(filterRecordIp)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieves existing reverse DNS records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dns"
)

func main() {
    filterRecordIp := []string{"Inner_example"} // []string | Filter is used to fetch only the reverse records for the specified IPs. It's an array of IP records. IP can be any valid IPv4 or IPv6 address. Parameter has to be sent in query as many times as values (filter.recordIp=1.2.3.4&filter.recordIp=2.3.4.5)  (optional)
    offset := int32(56) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ReverseRecordsApi.ReverserecordsGet(context.Background()).FilterRecordIp(filterRecordIp).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseRecordsApi.ReverserecordsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ReverserecordsGet`: ReverseRecordsReadList
    fmt.Fprintf(os.Stdout, "Response from `ReverseRecordsApi.ReverserecordsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiReverserecordsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **filterRecordIp** | [**[]string**](../models/string.md) | Filter is used to fetch only the reverse records for the specified IPs. It&#39;s an array of IP records. IP can be any valid IPv4 or IPv6 address. Parameter has to be sent in query as many times as values (filter.recordIp&#x3D;1.2.3.4&amp;filter.recordIp&#x3D;2.3.4.5)  | |
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**ReverseRecordsReadList**](../models/ReverseRecordsReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ReverseRecordsApiService.ReverserecordsGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ReverseRecordsApiService.ReverserecordsGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ReverseRecordsApiService.ReverserecordsGet": {
    "port": "8443",
},
})
```


## ReverserecordsPost

```go
var result ReverseRecordRead = ReverserecordsPost(ctx)
                      .ReverseRecordCreate(reverseRecordCreate)
                      .Execute()
```

Create a reverse DNS record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dns"
)

func main() {
    reverseRecordCreate := *openapiclient.NewReverseRecordCreate(*openapiclient.NewReverseRecord("mail.example.com", "5.6.7.8")) // ReverseRecordCreate | reverserecord

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ReverseRecordsApi.ReverserecordsPost(context.Background()).ReverseRecordCreate(reverseRecordCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseRecordsApi.ReverserecordsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ReverserecordsPost`: ReverseRecordRead
    fmt.Fprintf(os.Stdout, "Response from `ReverseRecordsApi.ReverserecordsPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiReverserecordsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **reverseRecordCreate** | [**ReverseRecordCreate**](../models/ReverseRecordCreate.md) | reverserecord | |

### Return type

[**ReverseRecordRead**](../models/ReverseRecordRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ReverseRecordsApiService.ReverserecordsPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ReverseRecordsApiService.ReverserecordsPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ReverseRecordsApiService.ReverserecordsPost": {
    "port": "8443",
},
})
```


## ReverserecordsPut

```go
var result ReverseRecordRead = ReverserecordsPut(ctx, reverserecordId)
                      .ReverseRecordEnsure(reverseRecordEnsure)
                      .Execute()
```

Update a reverse DNS record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dns"
)

func main() {
    reverserecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the reverse DNS record.
    reverseRecordEnsure := *openapiclient.NewReverseRecordEnsure(*openapiclient.NewReverseRecord("mail.example.com", "5.6.7.8")) // ReverseRecordEnsure | 

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ReverseRecordsApi.ReverserecordsPut(context.Background(), reverserecordId).ReverseRecordEnsure(reverseRecordEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseRecordsApi.ReverserecordsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ReverserecordsPut`: ReverseRecordRead
    fmt.Fprintf(os.Stdout, "Response from `ReverseRecordsApi.ReverserecordsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**reverserecordId** | **string** | The ID (UUID) of the reverse DNS record. | |

### Other Parameters

Other parameters are passed through a pointer to an apiReverserecordsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **reverseRecordEnsure** | [**ReverseRecordEnsure**](../models/ReverseRecordEnsure.md) |  | |

### Return type

[**ReverseRecordRead**](../models/ReverseRecordRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ReverseRecordsApiService.ReverserecordsPut"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ReverseRecordsApiService.ReverserecordsPut": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ReverseRecordsApiService.ReverserecordsPut": {
    "port": "8443",
},
})
```

