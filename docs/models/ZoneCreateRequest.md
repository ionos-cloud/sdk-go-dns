# ZoneCreateRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**ZoneCreateRequestProperties**](ZoneCreateRequestProperties.md) |  | |

## Methods

### NewZoneCreateRequest

`func NewZoneCreateRequest(properties ZoneCreateRequestProperties, ) *ZoneCreateRequest`

NewZoneCreateRequest instantiates a new ZoneCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateRequestWithDefaults

`func NewZoneCreateRequestWithDefaults() *ZoneCreateRequest`

NewZoneCreateRequestWithDefaults instantiates a new ZoneCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *ZoneCreateRequest) GetProperties() ZoneCreateRequestProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ZoneCreateRequest) GetPropertiesOk() (*ZoneCreateRequestProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ZoneCreateRequest) SetProperties(v ZoneCreateRequestProperties)`

SetProperties sets Properties field to given value.



