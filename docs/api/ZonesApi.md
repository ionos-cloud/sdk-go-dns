# \ZonesApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ZonesDelete**](ZonesApi.md#ZonesDelete) | **Delete** /zones/{zoneId} | Delete a zone|
|[**ZonesFindById**](ZonesApi.md#ZonesFindById) | **Get** /zones/{zoneId} | Retrieve a zone|
|[**ZonesGet**](ZonesApi.md#ZonesGet) | **Get** /zones | Retrieve zones|
|[**ZonesPost**](ZonesApi.md#ZonesPost) | **Post** /zones | Create a zone|
|[**ZonesPut**](ZonesApi.md#ZonesPut) | **Put** /zones/{zoneId} | Ensure a zone|



## ZonesDelete

```go
var result  = ZonesDelete(ctx, zoneId)
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

    ionoscloud "github.com/ionos-cloud/sdk-go-dnsaas"
)

func main() {
    zoneId := TODO // string | The ID (UUID) of the DNS zone.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ZonesApi.ZonesDelete(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ZonesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | [**string**](../models/.md) | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesFindById

```go
var result ZoneResponse = ZonesFindById(ctx, zoneId)
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

    ionoscloud "github.com/ionos-cloud/sdk-go-dnsaas"
)

func main() {
    zoneId := TODO // string | The ID (UUID) of the DNS zone.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ZonesApi.ZonesFindById(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ZonesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesFindById`: ZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ZonesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | [**string**](../models/.md) | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ZoneResponse**](../models/ZoneResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesGet

```go
var result ZonesResponse = ZonesGet(ctx)
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

    ionoscloud "github.com/ionos-cloud/sdk-go-dnsaas"
)

func main() {
    filterState := "CREATED" // string | Filter used to fetch all zones in a particular state (PROVISIONING, DEPROVISIONING, CREATED, FAILED). (optional)
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
    // response from `ZonesGet`: ZonesResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ZonesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiZonesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **filterState** | **string** | Filter used to fetch all zones in a particular state (PROVISIONING, DEPROVISIONING, CREATED, FAILED). | |
| **filterZoneName** | **string** | Filter used to fetch only the zones that contain the specified zone name. | |
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**ZonesResponse**](../models/ZonesResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesPost

```go
var result ZoneResponse = ZonesPost(ctx)
                      .ZoneCreateRequest(zoneCreateRequest)
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

    ionoscloud "github.com/ionos-cloud/sdk-go-dnsaas"
)

func main() {
    zoneCreateRequest := *openapiclient.NewZoneCreateRequest(*openapiclient.NewZoneCreateRequestProperties("example.com")) // ZoneCreateRequest | zone

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ZonesApi.ZonesPost(context.Background()).ZoneCreateRequest(zoneCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ZonesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesPost`: ZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ZonesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiZonesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **zoneCreateRequest** | [**ZoneCreateRequest**](../models/ZoneCreateRequest.md) | zone | |

### Return type

[**ZoneResponse**](../models/ZoneResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ZonesPut

```go
var result ZoneResponse = ZonesPut(ctx, zoneId)
                      .ZoneUpdateRequest(zoneUpdateRequest)
                      .Execute()
```

Ensure a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dnsaas"
)

func main() {
    zoneId := TODO // string | The ID (UUID) of the DNS zone.
    zoneUpdateRequest := *openapiclient.NewZoneUpdateRequest(*openapiclient.NewZoneUpdateRequestProperties()) // ZoneUpdateRequest | update zone

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ZonesApi.ZonesPut(context.Background(), zoneId).ZoneUpdateRequest(zoneUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ZonesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesPut`: ZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ZonesPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | [**string**](../models/.md) | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **zoneUpdateRequest** | [**ZoneUpdateRequest**](../models/ZoneUpdateRequest.md) | update zone | |

### Return type

[**ZoneResponse**](../models/ZoneResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


