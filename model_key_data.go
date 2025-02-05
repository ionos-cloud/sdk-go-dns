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

// KeyData Represents the separate components of the RDATA for a DNSKEY. The values must conform to the guidelines in [RFC-4034 Section 2.1](https://www.rfc-editor.org/rfc/rfc4034#section-2.1).
type KeyData struct {
	// Represents the key's metadata and usage information.
	Flags *int32 `json:"flags,omitempty"`
	// Represents the public key data in Base64 encoding.
	PubKey *string `json:"pubKey,omitempty"`
}

// NewKeyData instantiates a new KeyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyData() *KeyData {
	this := KeyData{}

	return &this
}

// NewKeyDataWithDefaults instantiates a new KeyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyDataWithDefaults() *KeyData {
	this := KeyData{}
	return &this
}

// GetFlags returns the Flags field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *KeyData) GetFlags() *int32 {
	if o == nil {
		return nil
	}

	return o.Flags

}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyData) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Flags, true
}

// SetFlags sets field value
func (o *KeyData) SetFlags(v int32) {

	o.Flags = &v

}

// HasFlags returns a boolean if a field has been set.
func (o *KeyData) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// GetPubKey returns the PubKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KeyData) GetPubKey() *string {
	if o == nil {
		return nil
	}

	return o.PubKey

}

// GetPubKeyOk returns a tuple with the PubKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyData) GetPubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PubKey, true
}

// SetPubKey sets field value
func (o *KeyData) SetPubKey(v string) {

	o.PubKey = &v

}

// HasPubKey returns a boolean if a field has been set.
func (o *KeyData) HasPubKey() bool {
	if o != nil && o.PubKey != nil {
		return true
	}

	return false
}

func (o KeyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}

	if o.PubKey != nil {
		toSerialize["pubKey"] = o.PubKey
	}

	return json.Marshal(toSerialize)
}

type NullableKeyData struct {
	value *KeyData
	isSet bool
}

func (v NullableKeyData) Get() *KeyData {
	return v.value
}

func (v *NullableKeyData) Set(val *KeyData) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyData) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyData(val *KeyData) *NullableKeyData {
	return &NullableKeyData{value: val, isSet: true}
}

func (v NullableKeyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
