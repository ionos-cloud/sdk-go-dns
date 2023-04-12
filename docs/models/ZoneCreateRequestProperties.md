# ZoneCreateRequestProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ZoneName** | **string** |  | |
|**Description** | Pointer to **string** |  | [optional] |
|**Enabled** | Pointer to **bool** |  | [optional] [default to true]|

## Methods

### NewZoneCreateRequestProperties

`func NewZoneCreateRequestProperties(zoneName string, ) *ZoneCreateRequestProperties`

NewZoneCreateRequestProperties instantiates a new ZoneCreateRequestProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateRequestPropertiesWithDefaults

`func NewZoneCreateRequestPropertiesWithDefaults() *ZoneCreateRequestProperties`

NewZoneCreateRequestPropertiesWithDefaults instantiates a new ZoneCreateRequestProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *ZoneCreateRequestProperties) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ZoneCreateRequestProperties) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ZoneCreateRequestProperties) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetDescription

`func (o *ZoneCreateRequestProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneCreateRequestProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneCreateRequestProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ZoneCreateRequestProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ZoneCreateRequestProperties) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ZoneCreateRequestProperties) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ZoneCreateRequestProperties) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ZoneCreateRequestProperties) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


