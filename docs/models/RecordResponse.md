# RecordResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** |  | [optional] |
|**Metadata** | Pointer to [**RecordMetadata**](RecordMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**RecordProperties**](RecordProperties.md) |  | [optional] |

## Methods

### NewRecordResponse

`func NewRecordResponse() *RecordResponse`

NewRecordResponse instantiates a new RecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordResponseWithDefaults

`func NewRecordResponseWithDefaults() *RecordResponse`

NewRecordResponseWithDefaults instantiates a new RecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecordResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecordResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *RecordResponse) GetMetadata() RecordMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecordResponse) GetMetadataOk() (*RecordMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecordResponse) SetMetadata(v RecordMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RecordResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *RecordResponse) GetProperties() RecordProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RecordResponse) GetPropertiesOk() (*RecordProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RecordResponse) SetProperties(v RecordProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *RecordResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


