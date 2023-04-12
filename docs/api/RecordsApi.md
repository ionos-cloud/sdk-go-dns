# \RecordsApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RecordsGet**](RecordsApi.md#RecordsGet) | **Get** /records | Retrieve all records|
|[**ZonesRecordsDelete**](RecordsApi.md#ZonesRecordsDelete) | **Delete** /zones/{zoneId}/records/{recordId} | Delete a record|
|[**ZonesRecordsFindById**](RecordsApi.md#ZonesRecordsFindById) | **Get** /zones/{zoneId}/records/{recordId} | Retrieve a record|
|[**ZonesRecordsGet**](RecordsApi.md#ZonesRecordsGet) | **Get** /zones/{zoneId}/records | Retrieve records|
|[**ZonesRecordsPost**](RecordsApi.md#ZonesRecordsPost) | **Post** /zones/{zoneId}/records | Create a record|
|[**ZonesRecordsPut**](RecordsApi.md#ZonesRecordsPut) | **Put** /zones/{zoneId}/records/{recordId} | Ensure a record|



## RecordsGet

```go
var result RecordsResponse = RecordsGet(ctx)
                      .FilterZoneId(filterZoneId)
                      .FilterName(filterName)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all records



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
    filterZoneId := TODO // string | Filter used to fetch only the records that contain specified zoneId. (optional)
    filterName := "app" // string | Filter used to fetch only the records that contain specified record name. (optional)
    offset := int32(56) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.RecordsGet(context.Background()).FilterZoneId(filterZoneId).FilterName(filterName).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.RecordsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RecordsGet`: RecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.RecordsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiRecordsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **filterZoneId** | [**string**](../models/.md) | Filter used to fetch only the records that contain specified zoneId. | |
| **filterName** | **string** | Filter used to fetch only the records that contain specified record name. | |
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**RecordsResponse**](../models/RecordsResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesRecordsDelete

```go
var result  = ZonesRecordsDelete(ctx, zoneId, recordId)
                      .Execute()
```

Delete a record



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
    recordId := TODO // string | The ID (UUID) of the record.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.ZonesRecordsDelete(context.Background(), zoneId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | [**string**](../models/.md) | The ID (UUID) of the DNS zone. | |
|**recordId** | [**string**](../models/.md) | The ID (UUID) of the record. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesRecordsFindById

```go
var result RecordResponse = ZonesRecordsFindById(ctx, zoneId, recordId)
                      .Execute()
```

Retrieve a record



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
    recordId := TODO // string | The ID (UUID) of the record.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.ZonesRecordsFindById(context.Background(), zoneId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesRecordsFindById`: RecordResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ZonesRecordsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | [**string**](../models/.md) | The ID (UUID) of the DNS zone. | |
|**recordId** | [**string**](../models/.md) | The ID (UUID) of the record. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**RecordResponse**](../models/RecordResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesRecordsGet

```go
var result RecordsResponse = ZonesRecordsGet(ctx, zoneId)
                      .Execute()
```

Retrieve records



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
    resource, resp, err := apiClient.RecordsApi.ZonesRecordsGet(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesRecordsGet`: RecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ZonesRecordsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | [**string**](../models/.md) | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**RecordsResponse**](../models/RecordsResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesRecordsPost

```go
var result RecordResponse = ZonesRecordsPost(ctx, zoneId)
                      .RecordCreateRequest(recordCreateRequest)
                      .Execute()
```

Create a record



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
    recordCreateRequest := *openapiclient.NewRecordCreateRequest(*openapiclient.NewRecordProperties("app", openapiclient.recordType("A"), "1.2.3.4")) // RecordCreateRequest | record

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.ZonesRecordsPost(context.Background(), zoneId).RecordCreateRequest(recordCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesRecordsPost`: RecordResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ZonesRecordsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | [**string**](../models/.md) | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **recordCreateRequest** | [**RecordCreateRequest**](../models/RecordCreateRequest.md) | record | |

### Return type

[**RecordResponse**](../models/RecordResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ZonesRecordsPut

```go
var result RecordResponse = ZonesRecordsPut(ctx, zoneId, recordId)
                      .RecordUpdateRequest(recordUpdateRequest)
                      .Execute()
```

Ensure a record



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
    recordId := TODO // string | The ID (UUID) of the DNS record.
    recordUpdateRequest := *openapiclient.NewRecordUpdateRequest(*openapiclient.NewRecordProperties("app", openapiclient.recordType("A"), "1.2.3.4")) // RecordUpdateRequest | 

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.ZonesRecordsPut(context.Background(), zoneId, recordId).RecordUpdateRequest(recordUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesRecordsPut`: RecordResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ZonesRecordsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | [**string**](../models/.md) | The ID (UUID) of the DNS zone. | |
|**recordId** | [**string**](../models/.md) | The ID (UUID) of the DNS record. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **recordUpdateRequest** | [**RecordUpdateRequest**](../models/RecordUpdateRequest.md) |  | |

### Return type

[**RecordResponse**](../models/RecordResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


