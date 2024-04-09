# \ZonesApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ZonesDelete**](ZonesApi.md#ZonesDelete) | **Delete** /zones/{zoneId} | Delete a zone|
|[**ZonesFindById**](ZonesApi.md#ZonesFindById) | **Get** /zones/{zoneId} | Retrieve a zone|
|[**ZonesGet**](ZonesApi.md#ZonesGet) | **Get** /zones | Retrieve zones|
|[**ZonesPost**](ZonesApi.md#ZonesPost) | **Post** /zones | Create a zone|
|[**ZonesPut**](ZonesApi.md#ZonesPut) | **Put** /zones/{zoneId} | Update a zone|



## ZonesDelete

```go
var result map[string]interface{} = ZonesDelete(ctx, zoneId)
                      .Execute()
```

Delete a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/go"
)

func main() {
    zoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.ZonesApi.ZonesDelete(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ZonesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ZonesDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ZonesApiService.ZonesDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ZonesApiService.ZonesDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ZonesApiService.ZonesDelete": {
    "port": "8443",
},
})
```


## ZonesFindById

```go
var result ZoneRead = ZonesFindById(ctx, zoneId)
                      .Execute()
```

Retrieve a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/go"
)

func main() {
    zoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ZonesApi.ZonesFindById(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ZonesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesFindById`: ZoneRead
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ZonesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ZoneRead**](../models/ZoneRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ZonesApiService.ZonesFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ZonesApiService.ZonesFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ZonesApiService.ZonesFindById": {
    "port": "8443",
},
})
```


## ZonesGet

```go
var result ZoneReadList = ZonesGet(ctx)
                      .FilterState(filterState)
                      .FilterZoneName(filterZoneName)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve zones



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/go"
)

func main() {
    filterState := openapiclient.ProvisioningState("PROVISIONING") // ProvisioningState | Filter used to fetch all zones in a particular state. (optional)
    filterZoneName := "example.com" // string | Filter used to fetch only the zones that contain the specified zone name. (optional)
    offset := int32(56) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ZonesApi.ZonesGet(context.Background()).FilterState(filterState).FilterZoneName(filterZoneName).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ZonesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesGet`: ZoneReadList
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ZonesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiZonesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **filterState** | [**ProvisioningState**](../models/.md) | Filter used to fetch all zones in a particular state. | |
| **filterZoneName** | **string** | Filter used to fetch only the zones that contain the specified zone name. | |
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**ZoneReadList**](../models/ZoneReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ZonesApiService.ZonesGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ZonesApiService.ZonesGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ZonesApiService.ZonesGet": {
    "port": "8443",
},
})
```


## ZonesPost

```go
var result ZoneRead = ZonesPost(ctx)
                      .ZoneCreate(zoneCreate)
                      .Execute()
```

Create a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/go"
)

func main() {
    zoneCreate := *openapiclient.NewZoneCreate(*openapiclient.NewZone("example.com")) // ZoneCreate | zone

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ZonesApi.ZonesPost(context.Background()).ZoneCreate(zoneCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ZonesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesPost`: ZoneRead
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ZonesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiZonesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **zoneCreate** | [**ZoneCreate**](../models/ZoneCreate.md) | zone | |

### Return type

[**ZoneRead**](../models/ZoneRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ZonesApiService.ZonesPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ZonesApiService.ZonesPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ZonesApiService.ZonesPost": {
    "port": "8443",
},
})
```


## ZonesPut

```go
var result ZoneRead = ZonesPut(ctx, zoneId)
                      .ZoneEnsure(zoneEnsure)
                      .Execute()
```

Update a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/go"
)

func main() {
    zoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.
    zoneEnsure := *openapiclient.NewZoneEnsure(*openapiclient.NewZone("example.com")) // ZoneEnsure | update zone

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ZonesApi.ZonesPut(context.Background(), zoneId).ZoneEnsure(zoneEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ZonesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesPut`: ZoneRead
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ZonesPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **zoneEnsure** | [**ZoneEnsure**](../models/ZoneEnsure.md) | update zone | |

### Return type

[**ZoneRead**](../models/ZoneRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ZonesApiService.ZonesPut"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ZonesApiService.ZonesPut": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ZonesApiService.ZonesPut": {
    "port": "8443",
},
})
```

