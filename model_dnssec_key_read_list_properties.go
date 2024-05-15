/*
 * IONOS Cloud - DNS API
 *
 * Cloud DNS service helps IONOS Cloud customers to automate DNS Zone and Record management.
 *
 * API version: 1.15.4
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// DnssecKeyReadListProperties Properties of the key.
type DnssecKeyReadListProperties struct {
	KeyParameters  *DnssecKeyReadListPropertiesKeyParameters  `json:"keyParameters"`
	NsecParameters *DnssecKeyReadListPropertiesNsecParameters `json:"nsecParameters"`
}

// NewDnssecKeyReadListProperties instantiates a new DnssecKeyReadListProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnssecKeyReadListProperties(keyParameters DnssecKeyReadListPropertiesKeyParameters, nsecParameters DnssecKeyReadListPropertiesNsecParameters) *DnssecKeyReadListProperties {
	this := DnssecKeyReadListProperties{}

	this.KeyParameters = &keyParameters
	this.NsecParameters = &nsecParameters

	return &this
}

// NewDnssecKeyReadListPropertiesWithDefaults instantiates a new DnssecKeyReadListProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnssecKeyReadListPropertiesWithDefaults() *DnssecKeyReadListProperties {
	this := DnssecKeyReadListProperties{}
	return &this
}

// GetKeyParameters returns the KeyParameters field value
// If the value is explicit nil, the zero value for DnssecKeyReadListPropertiesKeyParameters will be returned
func (o *DnssecKeyReadListProperties) GetKeyParameters() *DnssecKeyReadListPropertiesKeyParameters {
	if o == nil {
		return nil
	}

	return o.KeyParameters

}

// GetKeyParametersOk returns a tuple with the KeyParameters field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnssecKeyReadListProperties) GetKeyParametersOk() (*DnssecKeyReadListPropertiesKeyParameters, bool) {
	if o == nil {
		return nil, false
	}

	return o.KeyParameters, true
}

// SetKeyParameters sets field value
func (o *DnssecKeyReadListProperties) SetKeyParameters(v DnssecKeyReadListPropertiesKeyParameters) {

	o.KeyParameters = &v

}

// HasKeyParameters returns a boolean if a field has been set.
func (o *DnssecKeyReadListProperties) HasKeyParameters() bool {
	if o != nil && o.KeyParameters != nil {
		return true
	}

	return false
}

// GetNsecParameters returns the NsecParameters field value
// If the value is explicit nil, the zero value for DnssecKeyReadListPropertiesNsecParameters will be returned
func (o *DnssecKeyReadListProperties) GetNsecParameters() *DnssecKeyReadListPropertiesNsecParameters {
	if o == nil {
		return nil
	}

	return o.NsecParameters

}

// GetNsecParametersOk returns a tuple with the NsecParameters field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnssecKeyReadListProperties) GetNsecParametersOk() (*DnssecKeyReadListPropertiesNsecParameters, bool) {
	if o == nil {
		return nil, false
	}

	return o.NsecParameters, true
}

// SetNsecParameters sets field value
func (o *DnssecKeyReadListProperties) SetNsecParameters(v DnssecKeyReadListPropertiesNsecParameters) {

	o.NsecParameters = &v

}

// HasNsecParameters returns a boolean if a field has been set.
func (o *DnssecKeyReadListProperties) HasNsecParameters() bool {
	if o != nil && o.NsecParameters != nil {
		return true
	}

	return false
}

func (o DnssecKeyReadListProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyParameters != nil {
		toSerialize["keyParameters"] = o.KeyParameters
	}

	if o.NsecParameters != nil {
		toSerialize["nsecParameters"] = o.NsecParameters
	}

	return json.Marshal(toSerialize)
}

type NullableDnssecKeyReadListProperties struct {
	value *DnssecKeyReadListProperties
	isSet bool
}

func (v NullableDnssecKeyReadListProperties) Get() *DnssecKeyReadListProperties {
	return v.value
}

func (v *NullableDnssecKeyReadListProperties) Set(val *DnssecKeyReadListProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableDnssecKeyReadListProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableDnssecKeyReadListProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnssecKeyReadListProperties(val *DnssecKeyReadListProperties) *NullableDnssecKeyReadListProperties {
	return &NullableDnssecKeyReadListProperties{value: val, isSet: true}
}

func (v NullableDnssecKeyReadListProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnssecKeyReadListProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
