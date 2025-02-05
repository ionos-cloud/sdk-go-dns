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

// SecondaryZoneAllOf Indicates the fields for a secondary zone to be created
type SecondaryZoneAllOf struct {
	// Indicates IP addresses of primary nameservers for a secondary zone. Accepts IPv4 and IPv6 addresses
	PrimaryIps *[]string `json:"primaryIps"`
}

// NewSecondaryZoneAllOf instantiates a new SecondaryZoneAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecondaryZoneAllOf(primaryIps []string) *SecondaryZoneAllOf {
	this := SecondaryZoneAllOf{}

	this.PrimaryIps = &primaryIps

	return &this
}

// NewSecondaryZoneAllOfWithDefaults instantiates a new SecondaryZoneAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecondaryZoneAllOfWithDefaults() *SecondaryZoneAllOf {
	this := SecondaryZoneAllOf{}
	return &this
}

// GetPrimaryIps returns the PrimaryIps field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *SecondaryZoneAllOf) GetPrimaryIps() *[]string {
	if o == nil {
		return nil
	}

	return o.PrimaryIps

}

// GetPrimaryIpsOk returns a tuple with the PrimaryIps field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecondaryZoneAllOf) GetPrimaryIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PrimaryIps, true
}

// SetPrimaryIps sets field value
func (o *SecondaryZoneAllOf) SetPrimaryIps(v []string) {

	o.PrimaryIps = &v

}

// HasPrimaryIps returns a boolean if a field has been set.
func (o *SecondaryZoneAllOf) HasPrimaryIps() bool {
	if o != nil && o.PrimaryIps != nil {
		return true
	}

	return false
}

func (o SecondaryZoneAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrimaryIps != nil {
		toSerialize["primaryIps"] = o.PrimaryIps
	}

	return json.Marshal(toSerialize)
}

type NullableSecondaryZoneAllOf struct {
	value *SecondaryZoneAllOf
	isSet bool
}

func (v NullableSecondaryZoneAllOf) Get() *SecondaryZoneAllOf {
	return v.value
}

func (v *NullableSecondaryZoneAllOf) Set(val *SecondaryZoneAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecondaryZoneAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecondaryZoneAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecondaryZoneAllOf(val *SecondaryZoneAllOf) *NullableSecondaryZoneAllOf {
	return &NullableSecondaryZoneAllOf{value: val, isSet: true}
}

func (v NullableSecondaryZoneAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecondaryZoneAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
