# \DNSSECApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ZonesKeysDelete**](DNSSECApi.md#ZonesKeysDelete) | **Delete** /zones/{zoneId}/keys | Delete a DNSSEC key|
|[**ZonesKeysGet**](DNSSECApi.md#ZonesKeysGet) | **Get** /zones/{zoneId}/keys | Retrieve a DNSSEC key|
|[**ZonesKeysPost**](DNSSECApi.md#ZonesKeysPost) | **Post** /zones/{zoneId}/keys | Create a DNSSEC key|



## ZonesKeysDelete

```go
var result map[string]interface{} = ZonesKeysDelete(ctx, zoneId)
                      .Execute()
```

Delete a DNSSEC key



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
    zoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.DNSSECApi.ZonesKeysDelete(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSSECApi.ZonesKeysDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesKeysDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DNSSECApi.ZonesKeysDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesKeysDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"DNSSECApiService.ZonesKeysDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "DNSSECApiService.ZonesKeysDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "DNSSECApiService.ZonesKeysDelete": {
    "port": "8443",
},
})
```


## ZonesKeysGet

```go
var result DnssecKeyReadList = ZonesKeysGet(ctx, zoneId)
                      .Execute()
```

Retrieve a DNSSEC key



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
    zoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.DNSSECApi.ZonesKeysGet(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSSECApi.ZonesKeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesKeysGet`: DnssecKeyReadList
    fmt.Fprintf(os.Stdout, "Response from `DNSSECApi.ZonesKeysGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesKeysGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**DnssecKeyReadList**](../models/DnssecKeyReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"DNSSECApiService.ZonesKeysGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "DNSSECApiService.ZonesKeysGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "DNSSECApiService.ZonesKeysGet": {
    "port": "8443",
},
})
```


## ZonesKeysPost

```go
var result DnssecKeyReadCreation = ZonesKeysPost(ctx, zoneId)
                      .DnssecKeyCreate(dnssecKeyCreate)
                      .Execute()
```

Create a DNSSEC key



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
    zoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.
    dnssecKeyCreate := *openapiclient.NewDnssecKeyCreate(*openapiclient.NewDnssecKeyParameters(*openapiclient.NewKeyParameters(openapiclient.algorithm("RSASHA256"), openapiclient.kskBits(1024), openapiclient.zskBits(1024)), *openapiclient.NewNsecParameters(openapiclient.nsecMode("NSEC"), int32(21), int32(128)), int32(120))) // DnssecKeyCreate | Enable DNSSEC request.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.DNSSECApi.ZonesKeysPost(context.Background(), zoneId).DnssecKeyCreate(dnssecKeyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSSECApi.ZonesKeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesKeysPost`: DnssecKeyReadCreation
    fmt.Fprintf(os.Stdout, "Response from `DNSSECApi.ZonesKeysPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesKeysPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dnssecKeyCreate** | [**DnssecKeyCreate**](../models/DnssecKeyCreate.md) | Enable DNSSEC request. | |

### Return type

[**DnssecKeyReadCreation**](../models/DnssecKeyReadCreation.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"DNSSECApiService.ZonesKeysPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "DNSSECApiService.ZonesKeysPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "DNSSECApiService.ZonesKeysPost": {
    "port": "8443",
},
})
```

