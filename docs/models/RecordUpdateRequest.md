# RecordUpdateRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**RecordProperties**](RecordProperties.md) |  | |

## Methods

### NewRecordUpdateRequest

`func NewRecordUpdateRequest(properties RecordProperties, ) *RecordUpdateRequest`

NewRecordUpdateRequest instantiates a new RecordUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordUpdateRequestWithDefaults

`func NewRecordUpdateRequestWithDefaults() *RecordUpdateRequest`

NewRecordUpdateRequestWithDefaults instantiates a new RecordUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *RecordUpdateRequest) GetProperties() RecordProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RecordUpdateRequest) GetPropertiesOk() (*RecordProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RecordUpdateRequest) SetProperties(v RecordProperties)`

SetProperties sets Properties field to given value.



