# \ZoneFilesApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ZonesZonefileGet**](ZoneFilesApi.md#ZonesZonefileGet) | **Get** /zones/{zoneId}/zonefile | Retrieve a zone file|
|[**ZonesZonefilePut**](ZoneFilesApi.md#ZonesZonefilePut) | **Put** /zones/{zoneId}/zonefile | Updates a zone with a file|



## ZonesZonefileGet

```go
var result  = ZonesZonefileGet(ctx, zoneId)
                      .Execute()
```

Retrieve a zone file



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
    resource, resp, err := apiClient.ZoneFilesApi.ZonesZonefileGet(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneFilesApi.ZonesZonefileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesZonefileGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plainapplication/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ZoneFilesApiService.ZonesZonefileGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ZoneFilesApiService.ZonesZonefileGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ZoneFilesApiService.ZonesZonefileGet": {
    "port": "8443",
},
})
```


## ZonesZonefilePut

```go
var result RecordReadList = ZonesZonefilePut(ctx, zoneId)
                      .Body(body)
                      .Execute()
```

Updates a zone with a file



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
    body := "body_example" // string | Zone file in BIND format (RFC 1035). In order to support import files from other sources, the bind zone file can contain SOA and NS records, but these records will be ignored.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ZoneFilesApi.ZonesZonefilePut(context.Background(), zoneId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneFilesApi.ZonesZonefilePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesZonefilePut`: RecordReadList
    fmt.Fprintf(os.Stdout, "Response from `ZoneFilesApi.ZonesZonefilePut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesZonefilePutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **string** | Zone file in BIND format (RFC 1035). In order to support import files from other sources, the bind zone file can contain SOA and NS records, but these records will be ignored. | |

### Return type

[**RecordReadList**](../models/RecordReadList.md)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ZoneFilesApiService.ZonesZonefilePut"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ZoneFilesApiService.ZonesZonefilePut": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ZoneFilesApiService.ZonesZonefilePut": {
    "port": "8443",
},
})
```

