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

// ZoneEnsure struct for ZoneEnsure
type ZoneEnsure struct {
	Properties *Zone `json:"properties"`
}

// NewZoneEnsure instantiates a new ZoneEnsure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneEnsure(properties Zone) *ZoneEnsure {
	this := ZoneEnsure{}

	this.Properties = &properties

	return &this
}

// NewZoneEnsureWithDefaults instantiates a new ZoneEnsure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneEnsureWithDefaults() *ZoneEnsure {
	this := ZoneEnsure{}
	return &this
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for Zone will be returned
func (o *ZoneEnsure) GetProperties() *Zone {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneEnsure) GetPropertiesOk() (*Zone, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *ZoneEnsure) SetProperties(v Zone) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *ZoneEnsure) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o ZoneEnsure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableZoneEnsure struct {
	value *ZoneEnsure
	isSet bool
}

func (v NullableZoneEnsure) Get() *ZoneEnsure {
	return v.value
}

func (v *NullableZoneEnsure) Set(val *ZoneEnsure) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneEnsure) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneEnsure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneEnsure(val *ZoneEnsure) *NullableZoneEnsure {
	return &NullableZoneEnsure{value: val, isSet: true}
}

func (v NullableZoneEnsure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneEnsure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
