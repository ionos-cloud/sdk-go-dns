/*
 * IONOS Cloud - DNS API
 *
 * DNS API Specification
 *
 * API version: 1.0.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// RecordCreateRequest struct for RecordCreateRequest
type RecordCreateRequest struct {
	Properties *RecordProperties `json:"properties"`
}

// NewRecordCreateRequest instantiates a new RecordCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordCreateRequest(properties RecordProperties) *RecordCreateRequest {
	this := RecordCreateRequest{}

	this.Properties = &properties

	return &this
}

// NewRecordCreateRequestWithDefaults instantiates a new RecordCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordCreateRequestWithDefaults() *RecordCreateRequest {
	this := RecordCreateRequest{}
	return &this
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for RecordProperties will be returned
func (o *RecordCreateRequest) GetProperties() *RecordProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecordCreateRequest) GetPropertiesOk() (*RecordProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *RecordCreateRequest) SetProperties(v RecordProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *RecordCreateRequest) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o RecordCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableRecordCreateRequest struct {
	value *RecordCreateRequest
	isSet bool
}

func (v NullableRecordCreateRequest) Get() *RecordCreateRequest {
	return v.value
}

func (v *NullableRecordCreateRequest) Set(val *RecordCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordCreateRequest(val *RecordCreateRequest) *NullableRecordCreateRequest {
	return &NullableRecordCreateRequest{value: val, isSet: true}
}

func (v NullableRecordCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
