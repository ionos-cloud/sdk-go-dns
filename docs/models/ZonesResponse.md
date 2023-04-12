# ZonesResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Items** | Pointer to [**[]ZoneResponse**](ZoneResponse.md) |  | [optional] |
|**Offset** | Pointer to **float32** | Pagination offset. | [optional] [readonly] |
|**Limit** | Pointer to **float32** | Pagination limit. | [optional] [readonly] |

## Methods

### NewZonesResponse

`func NewZonesResponse() *ZonesResponse`

NewZonesResponse instantiates a new ZonesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonesResponseWithDefaults

`func NewZonesResponseWithDefaults() *ZonesResponse`

NewZonesResponseWithDefaults instantiates a new ZonesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ZonesResponse) GetItems() []ZoneResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ZonesResponse) GetItemsOk() (*[]ZoneResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ZonesResponse) SetItems(v []ZoneResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *ZonesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *ZonesResponse) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ZonesResponse) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ZonesResponse) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ZonesResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *ZonesResponse) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ZonesResponse) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ZonesResponse) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ZonesResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


