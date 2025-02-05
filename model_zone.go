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

// Zone struct for Zone
type Zone struct {
	// The zone name
	ZoneName *string `json:"zoneName"`
	// The hosted zone is used for...
	Description *string `json:"description,omitempty"`
	// Users can activate and deactivate zones.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewZone instantiates a new Zone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZone(zoneName string) *Zone {
	this := Zone{}

	this.ZoneName = &zoneName
	var enabled bool = true
	this.Enabled = &enabled

	return &this
}

// NewZoneWithDefaults instantiates a new Zone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneWithDefaults() *Zone {
	this := Zone{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetZoneName returns the ZoneName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Zone) GetZoneName() *string {
	if o == nil {
		return nil
	}

	return o.ZoneName

}

// GetZoneNameOk returns a tuple with the ZoneName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Zone) GetZoneNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ZoneName, true
}

// SetZoneName sets field value
func (o *Zone) SetZoneName(v string) {

	o.ZoneName = &v

}

// HasZoneName returns a boolean if a field has been set.
func (o *Zone) HasZoneName() bool {
	if o != nil && o.ZoneName != nil {
		return true
	}

	return false
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Zone) GetDescription() *string {
	if o == nil {
		return nil
	}

	return o.Description

}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Zone) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Description, true
}

// SetDescription sets field value
func (o *Zone) SetDescription(v string) {

	o.Description = &v

}

// HasDescription returns a boolean if a field has been set.
func (o *Zone) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Zone) GetEnabled() *bool {
	if o == nil {
		return nil
	}

	return o.Enabled

}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Zone) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.Enabled, true
}

// SetEnabled sets field value
func (o *Zone) SetEnabled(v bool) {

	o.Enabled = &v

}

// HasEnabled returns a boolean if a field has been set.
func (o *Zone) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

func (o Zone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneName != nil {
		toSerialize["zoneName"] = o.ZoneName
	}

	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}

	return json.Marshal(toSerialize)
}

type NullableZone struct {
	value *Zone
	isSet bool
}

func (v NullableZone) Get() *Zone {
	return v.value
}

func (v *NullableZone) Set(val *Zone) {
	v.value = val
	v.isSet = true
}

func (v NullableZone) IsSet() bool {
	return v.isSet
}

func (v *NullableZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZone(val *Zone) *NullableZone {
	return &NullableZone{value: val, isSet: true}
}

func (v NullableZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
