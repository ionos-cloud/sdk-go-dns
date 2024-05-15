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

// DnssecKeyCreate struct for DnssecKeyCreate
type DnssecKeyCreate struct {
	Properties *DnssecKeyParameters `json:"properties"`
}

// NewDnssecKeyCreate instantiates a new DnssecKeyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnssecKeyCreate(properties DnssecKeyParameters) *DnssecKeyCreate {
	this := DnssecKeyCreate{}

	this.Properties = &properties

	return &this
}

// NewDnssecKeyCreateWithDefaults instantiates a new DnssecKeyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnssecKeyCreateWithDefaults() *DnssecKeyCreate {
	this := DnssecKeyCreate{}
	return &this
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for DnssecKeyParameters will be returned
func (o *DnssecKeyCreate) GetProperties() *DnssecKeyParameters {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnssecKeyCreate) GetPropertiesOk() (*DnssecKeyParameters, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *DnssecKeyCreate) SetProperties(v DnssecKeyParameters) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *DnssecKeyCreate) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o DnssecKeyCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableDnssecKeyCreate struct {
	value *DnssecKeyCreate
	isSet bool
}

func (v NullableDnssecKeyCreate) Get() *DnssecKeyCreate {
	return v.value
}

func (v *NullableDnssecKeyCreate) Set(val *DnssecKeyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDnssecKeyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDnssecKeyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnssecKeyCreate(val *DnssecKeyCreate) *NullableDnssecKeyCreate {
	return &NullableDnssecKeyCreate{value: val, isSet: true}
}

func (v NullableDnssecKeyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnssecKeyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}