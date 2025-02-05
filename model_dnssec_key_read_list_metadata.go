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

// DnssecKeyReadListMetadata Metadata of the resource with not state information.
type DnssecKeyReadListMetadata struct {
	// The ID (UUID) of the DNS zone of which record belongs to.
	ZoneId *string      `json:"zoneId,omitempty"`
	Items  *[]DnssecKey `json:"items,omitempty"`
}

// NewDnssecKeyReadListMetadata instantiates a new DnssecKeyReadListMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnssecKeyReadListMetadata() *DnssecKeyReadListMetadata {
	this := DnssecKeyReadListMetadata{}

	return &this
}

// NewDnssecKeyReadListMetadataWithDefaults instantiates a new DnssecKeyReadListMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnssecKeyReadListMetadataWithDefaults() *DnssecKeyReadListMetadata {
	this := DnssecKeyReadListMetadata{}
	return &this
}

// GetZoneId returns the ZoneId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DnssecKeyReadListMetadata) GetZoneId() *string {
	if o == nil {
		return nil
	}

	return o.ZoneId

}

// GetZoneIdOk returns a tuple with the ZoneId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnssecKeyReadListMetadata) GetZoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ZoneId, true
}

// SetZoneId sets field value
func (o *DnssecKeyReadListMetadata) SetZoneId(v string) {

	o.ZoneId = &v

}

// HasZoneId returns a boolean if a field has been set.
func (o *DnssecKeyReadListMetadata) HasZoneId() bool {
	if o != nil && o.ZoneId != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []DnssecKey will be returned
func (o *DnssecKeyReadListMetadata) GetItems() *[]DnssecKey {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnssecKeyReadListMetadata) GetItemsOk() (*[]DnssecKey, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *DnssecKeyReadListMetadata) SetItems(v []DnssecKey) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *DnssecKeyReadListMetadata) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o DnssecKeyReadListMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneId != nil {
		toSerialize["zoneId"] = o.ZoneId
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableDnssecKeyReadListMetadata struct {
	value *DnssecKeyReadListMetadata
	isSet bool
}

func (v NullableDnssecKeyReadListMetadata) Get() *DnssecKeyReadListMetadata {
	return v.value
}

func (v *NullableDnssecKeyReadListMetadata) Set(val *DnssecKeyReadListMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDnssecKeyReadListMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDnssecKeyReadListMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnssecKeyReadListMetadata(val *DnssecKeyReadListMetadata) *NullableDnssecKeyReadListMetadata {
	return &NullableDnssecKeyReadListMetadata{value: val, isSet: true}
}

func (v NullableDnssecKeyReadListMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnssecKeyReadListMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
