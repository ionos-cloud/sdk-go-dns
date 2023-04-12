# ZoneResponseProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ZoneName** | Pointer to **string** | The zone name | [optional] |
|**Description** | Pointer to **string** | The hosted zone is used for... | [optional] |
|**Enabled** | Pointer to **bool** | Users can activate and deactivate zones. | [optional] [default to true]|

## Methods

### NewZoneResponseProperties

`func NewZoneResponseProperties() *ZoneResponseProperties`

NewZoneResponseProperties instantiates a new ZoneResponseProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneResponsePropertiesWithDefaults

`func NewZoneResponsePropertiesWithDefaults() *ZoneResponseProperties`

NewZoneResponsePropertiesWithDefaults instantiates a new ZoneResponseProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *ZoneResponseProperties) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ZoneResponseProperties) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ZoneResponseProperties) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ZoneResponseProperties) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetDescription

`func (o *ZoneResponseProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneResponseProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneResponseProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ZoneResponseProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ZoneResponseProperties) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ZoneResponseProperties) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ZoneResponseProperties) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ZoneResponseProperties) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


