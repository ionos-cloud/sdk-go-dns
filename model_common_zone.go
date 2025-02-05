/*
 * IONOS Cloud - DNS API
 *
 * Cloud DNS service helps IONOS Cloud customers to automate DNS Zone and Record management.
 *
 * API version: 1.17.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// CommonZone Indicates the fields for a zone to be created
type CommonZone struct {
	// The zone name
	ZoneName *string `json:"zoneName"`
	// The hosted zone is used for...
	Description *string `json:"description,omitempty"`
}

// NewCommonZone instantiates a new CommonZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonZone(zoneName string) *CommonZone {
	this := CommonZone{}

	this.ZoneName = &zoneName

	return &this
}

// NewCommonZoneWithDefaults instantiates a new CommonZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonZoneWithDefaults() *CommonZone {
	this := CommonZone{}
	return &this
}

// GetZoneName returns the ZoneName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonZone) GetZoneName() *string {
	if o == nil {
		return nil
	}

	return o.ZoneName

}

// GetZoneNameOk returns a tuple with the ZoneName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonZone) GetZoneNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ZoneName, true
}

// SetZoneName sets field value
func (o *CommonZone) SetZoneName(v string) {

	o.ZoneName = &v

}

// HasZoneName returns a boolean if a field has been set.
func (o *CommonZone) HasZoneName() bool {
	if o != nil && o.ZoneName != nil {
		return true
	}

	return false
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonZone) GetDescription() *string {
	if o == nil {
		return nil
	}

	return o.Description

}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonZone) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Description, true
}

// SetDescription sets field value
func (o *CommonZone) SetDescription(v string) {

	o.Description = &v

}

// HasDescription returns a boolean if a field has been set.
func (o *CommonZone) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

func (o CommonZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneName != nil {
		toSerialize["zoneName"] = o.ZoneName
	}

	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	return json.Marshal(toSerialize)
}

type NullableCommonZone struct {
	value *CommonZone
	isSet bool
}

func (v NullableCommonZone) Get() *CommonZone {
	return v.value
}

func (v *NullableCommonZone) Set(val *CommonZone) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonZone) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonZone(val *CommonZone) *NullableCommonZone {
	return &NullableCommonZone{value: val, isSet: true}
}

func (v NullableCommonZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
