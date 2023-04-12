# RecordProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** |  | |
|**Type** | [**RecordType**](RecordType.md) |  | |
|**Content** | **string** |  | |
|**Ttl** | Pointer to **int32** | Time to live for the record, recommended 3600. | [optional] [default to 3600]|
|**Priority** | Pointer to **int32** | Priority value is between 0 and 65535. Priority is mandatory for MX, SRV and URI record types and ignored for all other types. | [optional] |
|**Enabled** | Pointer to **bool** | When true - the record is visible for lookup. | [optional] [default to true]|

## Methods

### NewRecordProperties

`func NewRecordProperties(name string, type_ RecordType, content string, ) *RecordProperties`

NewRecordProperties instantiates a new RecordProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordPropertiesWithDefaults

`func NewRecordPropertiesWithDefaults() *RecordProperties`

NewRecordPropertiesWithDefaults instantiates a new RecordProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecordProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordProperties) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RecordProperties) GetType() RecordType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordProperties) GetTypeOk() (*RecordType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordProperties) SetType(v RecordType)`

SetType sets Type field to given value.


### GetContent

`func (o *RecordProperties) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RecordProperties) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RecordProperties) SetContent(v string)`

SetContent sets Content field to given value.


### GetTtl

`func (o *RecordProperties) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordProperties) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordProperties) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordProperties) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetPriority

`func (o *RecordProperties) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RecordProperties) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RecordProperties) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RecordProperties) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEnabled

`func (o *RecordProperties) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RecordProperties) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RecordProperties) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RecordProperties) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


