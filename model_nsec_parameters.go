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

// NsecParameters Nsec parameters.
type NsecParameters struct {
	NsecMode *NsecMode `json:"nsecMode"`
	// Number of iterations for NSEC3. (between 0 and 50)
	Nsec3Iterations *int32 `json:"nsec3Iterations"`
	// Salt length in bits for NSEC3. (between 64 and 128, multiples of 8)
	Nsec3SaltBits *int32 `json:"nsec3SaltBits"`
}

// NewNsecParameters instantiates a new NsecParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsecParameters(nsecMode NsecMode, nsec3Iterations int32, nsec3SaltBits int32) *NsecParameters {
	this := NsecParameters{}

	this.NsecMode = &nsecMode
	this.Nsec3Iterations = &nsec3Iterations
	this.Nsec3SaltBits = &nsec3SaltBits

	return &this
}

// NewNsecParametersWithDefaults instantiates a new NsecParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsecParametersWithDefaults() *NsecParameters {
	this := NsecParameters{}
	return &this
}

// GetNsecMode returns the NsecMode field value
// If the value is explicit nil, the zero value for NsecMode will be returned
func (o *NsecParameters) GetNsecMode() *NsecMode {
	if o == nil {
		return nil
	}

	return o.NsecMode

}

// GetNsecModeOk returns a tuple with the NsecMode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NsecParameters) GetNsecModeOk() (*NsecMode, bool) {
	if o == nil {
		return nil, false
	}

	return o.NsecMode, true
}

// SetNsecMode sets field value
func (o *NsecParameters) SetNsecMode(v NsecMode) {

	o.NsecMode = &v

}

// HasNsecMode returns a boolean if a field has been set.
func (o *NsecParameters) HasNsecMode() bool {
	if o != nil && o.NsecMode != nil {
		return true
	}

	return false
}

// GetNsec3Iterations returns the Nsec3Iterations field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *NsecParameters) GetNsec3Iterations() *int32 {
	if o == nil {
		return nil
	}

	return o.Nsec3Iterations

}

// GetNsec3IterationsOk returns a tuple with the Nsec3Iterations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NsecParameters) GetNsec3IterationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Nsec3Iterations, true
}

// SetNsec3Iterations sets field value
func (o *NsecParameters) SetNsec3Iterations(v int32) {

	o.Nsec3Iterations = &v

}

// HasNsec3Iterations returns a boolean if a field has been set.
func (o *NsecParameters) HasNsec3Iterations() bool {
	if o != nil && o.Nsec3Iterations != nil {
		return true
	}

	return false
}

// GetNsec3SaltBits returns the Nsec3SaltBits field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *NsecParameters) GetNsec3SaltBits() *int32 {
	if o == nil {
		return nil
	}

	return o.Nsec3SaltBits

}

// GetNsec3SaltBitsOk returns a tuple with the Nsec3SaltBits field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NsecParameters) GetNsec3SaltBitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Nsec3SaltBits, true
}

// SetNsec3SaltBits sets field value
func (o *NsecParameters) SetNsec3SaltBits(v int32) {

	o.Nsec3SaltBits = &v

}

// HasNsec3SaltBits returns a boolean if a field has been set.
func (o *NsecParameters) HasNsec3SaltBits() bool {
	if o != nil && o.Nsec3SaltBits != nil {
		return true
	}

	return false
}

func (o NsecParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NsecMode != nil {
		toSerialize["nsecMode"] = o.NsecMode
	}

	if o.Nsec3Iterations != nil {
		toSerialize["nsec3Iterations"] = o.Nsec3Iterations
	}

	if o.Nsec3SaltBits != nil {
		toSerialize["nsec3SaltBits"] = o.Nsec3SaltBits
	}

	return json.Marshal(toSerialize)
}

type NullableNsecParameters struct {
	value *NsecParameters
	isSet bool
}

func (v NullableNsecParameters) Get() *NsecParameters {
	return v.value
}

func (v *NullableNsecParameters) Set(val *NsecParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableNsecParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableNsecParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsecParameters(val *NsecParameters) *NullableNsecParameters {
	return &NullableNsecParameters{value: val, isSet: true}
}

func (v NullableNsecParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsecParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
