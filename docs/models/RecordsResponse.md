# RecordsResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Items** | Pointer to [**[]RecordResponse**](RecordResponse.md) |  | [optional] |
|**Offset** | Pointer to **float32** | Pagination offset. | [optional] [readonly] |
|**Limit** | Pointer to **float32** | Pagination limit. | [optional] [readonly] |

## Methods

### NewRecordsResponse

`func NewRecordsResponse() *RecordsResponse`

NewRecordsResponse instantiates a new RecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordsResponseWithDefaults

`func NewRecordsResponseWithDefaults() *RecordsResponse`

NewRecordsResponseWithDefaults instantiates a new RecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RecordsResponse) GetItems() []RecordResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RecordsResponse) GetItemsOk() (*[]RecordResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RecordsResponse) SetItems(v []RecordResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *RecordsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *RecordsResponse) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RecordsResponse) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RecordsResponse) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RecordsResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *RecordsResponse) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RecordsResponse) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RecordsResponse) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RecordsResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


