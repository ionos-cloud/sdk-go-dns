# RecordCreateRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**RecordProperties**](RecordProperties.md) |  | |

## Methods

### NewRecordCreateRequest

`func NewRecordCreateRequest(properties RecordProperties, ) *RecordCreateRequest`

NewRecordCreateRequest instantiates a new RecordCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordCreateRequestWithDefaults

`func NewRecordCreateRequestWithDefaults() *RecordCreateRequest`

NewRecordCreateRequestWithDefaults instantiates a new RecordCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *RecordCreateRequest) GetProperties() RecordProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RecordCreateRequest) GetPropertiesOk() (*RecordProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RecordCreateRequest) SetProperties(v RecordProperties)`

SetProperties sets Properties field to given value.



