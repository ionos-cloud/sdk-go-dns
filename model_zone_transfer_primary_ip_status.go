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

// ZoneTransferPrimaryIpStatus struct for ZoneTransferPrimaryIpStatus
type ZoneTransferPrimaryIpStatus struct {
	// one single IP from the primaryIps field for secondary zones
	PrimaryIp *string `json:"primaryIp"`
	// Human readable status of the zone transfer status for the IP
	Status *string `json:"status"`
	// Human readable explanation of the error when status is not ok
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewZoneTransferPrimaryIpStatus instantiates a new ZoneTransferPrimaryIpStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneTransferPrimaryIpStatus(primaryIp string, status string) *ZoneTransferPrimaryIpStatus {
	this := ZoneTransferPrimaryIpStatus{}

	this.PrimaryIp = &primaryIp
	this.Status = &status

	return &this
}

// NewZoneTransferPrimaryIpStatusWithDefaults instantiates a new ZoneTransferPrimaryIpStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneTransferPrimaryIpStatusWithDefaults() *ZoneTransferPrimaryIpStatus {
	this := ZoneTransferPrimaryIpStatus{}
	return &this
}

// GetPrimaryIp returns the PrimaryIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ZoneTransferPrimaryIpStatus) GetPrimaryIp() *string {
	if o == nil {
		return nil
	}

	return o.PrimaryIp

}

// GetPrimaryIpOk returns a tuple with the PrimaryIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneTransferPrimaryIpStatus) GetPrimaryIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PrimaryIp, true
}

// SetPrimaryIp sets field value
func (o *ZoneTransferPrimaryIpStatus) SetPrimaryIp(v string) {

	o.PrimaryIp = &v

}

// HasPrimaryIp returns a boolean if a field has been set.
func (o *ZoneTransferPrimaryIpStatus) HasPrimaryIp() bool {
	if o != nil && o.PrimaryIp != nil {
		return true
	}

	return false
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ZoneTransferPrimaryIpStatus) GetStatus() *string {
	if o == nil {
		return nil
	}

	return o.Status

}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneTransferPrimaryIpStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Status, true
}

// SetStatus sets field value
func (o *ZoneTransferPrimaryIpStatus) SetStatus(v string) {

	o.Status = &v

}

// HasStatus returns a boolean if a field has been set.
func (o *ZoneTransferPrimaryIpStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// GetErrorMessage returns the ErrorMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ZoneTransferPrimaryIpStatus) GetErrorMessage() *string {
	if o == nil {
		return nil
	}

	return o.ErrorMessage

}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneTransferPrimaryIpStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *ZoneTransferPrimaryIpStatus) SetErrorMessage(v string) {

	o.ErrorMessage = &v

}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ZoneTransferPrimaryIpStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

func (o ZoneTransferPrimaryIpStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrimaryIp != nil {
		toSerialize["primaryIp"] = o.PrimaryIp
	}

	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}

	return json.Marshal(toSerialize)
}

type NullableZoneTransferPrimaryIpStatus struct {
	value *ZoneTransferPrimaryIpStatus
	isSet bool
}

func (v NullableZoneTransferPrimaryIpStatus) Get() *ZoneTransferPrimaryIpStatus {
	return v.value
}

func (v *NullableZoneTransferPrimaryIpStatus) Set(val *ZoneTransferPrimaryIpStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneTransferPrimaryIpStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneTransferPrimaryIpStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneTransferPrimaryIpStatus(val *ZoneTransferPrimaryIpStatus) *NullableZoneTransferPrimaryIpStatus {
	return &NullableZoneTransferPrimaryIpStatus{value: val, isSet: true}
}

func (v NullableZoneTransferPrimaryIpStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneTransferPrimaryIpStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
