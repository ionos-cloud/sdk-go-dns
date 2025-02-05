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

// Links URLs to navigate the different pages. As of now we always only return a single page.
type Links struct {
	// URL (with offset and limit parameters) of the previous page; only present if offset is greater than 0.
	Prev *string `json:"prev,omitempty"`
	// URL (with offset and limit parameters) of the current page.
	Self *string `json:"self,omitempty"`
	// URL (with offset and limit parameters) of the next page; only present if offset + limit is less than the total number of elements.
	Next *string `json:"next,omitempty"`
}

// NewLinks instantiates a new Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinks() *Links {
	this := Links{}

	return &this
}

// NewLinksWithDefaults instantiates a new Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetPrev returns the Prev field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Links) GetPrev() *string {
	if o == nil {
		return nil
	}

	return o.Prev

}

// GetPrevOk returns a tuple with the Prev field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Links) GetPrevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Prev, true
}

// SetPrev sets field value
func (o *Links) SetPrev(v string) {

	o.Prev = &v

}

// HasPrev returns a boolean if a field has been set.
func (o *Links) HasPrev() bool {
	if o != nil && o.Prev != nil {
		return true
	}

	return false
}

// GetSelf returns the Self field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Links) GetSelf() *string {
	if o == nil {
		return nil
	}

	return o.Self

}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Links) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Self, true
}

// SetSelf sets field value
func (o *Links) SetSelf(v string) {

	o.Self = &v

}

// HasSelf returns a boolean if a field has been set.
func (o *Links) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Links) GetNext() *string {
	if o == nil {
		return nil
	}

	return o.Next

}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Links) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Next, true
}

// SetNext sets field value
func (o *Links) SetNext(v string) {

	o.Next = &v

}

// HasNext returns a boolean if a field has been set.
func (o *Links) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Prev != nil {
		toSerialize["prev"] = o.Prev
	}

	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	if o.Next != nil {
		toSerialize["next"] = o.Next
	}

	return json.Marshal(toSerialize)
}

type NullableLinks struct {
	value *Links
	isSet bool
}

func (v NullableLinks) Get() *Links {
	return v.value
}

func (v *NullableLinks) Set(val *Links) {
	v.value = val
	v.isSet = true
}

func (v NullableLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinks(val *Links) *NullableLinks {
	return &NullableLinks{value: val, isSet: true}
}

func (v NullableLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
