# ZoneUpdateRequestProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ZoneName** | Pointer to **string** |  | [optional] |
|**Description** | Pointer to **string** |  | [optional] |
|**Enabled** | Pointer to **bool** |  | [optional] |

## Methods

### NewZoneUpdateRequestProperties

`func NewZoneUpdateRequestProperties() *ZoneUpdateRequestProperties`

NewZoneUpdateRequestProperties instantiates a new ZoneUpdateRequestProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneUpdateRequestPropertiesWithDefaults

`func NewZoneUpdateRequestPropertiesWithDefaults() *ZoneUpdateRequestProperties`

NewZoneUpdateRequestPropertiesWithDefaults instantiates a new ZoneUpdateRequestProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *ZoneUpdateRequestProperties) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ZoneUpdateRequestProperties) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ZoneUpdateRequestProperties) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ZoneUpdateRequestProperties) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetDescription

`func (o *ZoneUpdateRequestProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneUpdateRequestProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneUpdateRequestProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ZoneUpdateRequestProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ZoneUpdateRequestProperties) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ZoneUpdateRequestProperties) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ZoneUpdateRequestProperties) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ZoneUpdateRequestProperties) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


