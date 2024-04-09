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

// ZoneReadList struct for ZoneReadList
type ZoneReadList struct {
	// ID (UUID) created to identify this action.
	Id   *string `json:"id"`
	Type *string `json:"type"`
	Href *string `json:"href"`
	// Pagination offset.
	Offset *float32 `json:"offset"`
	// Pagination limit.
	Limit *float32    `json:"limit"`
	Links *Links      `json:"_links"`
	Items *[]ZoneRead `json:"items"`
}

// NewZoneReadList instantiates a new ZoneReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneReadList(id string, type_ string, href string, offset float32, limit float32, links Links, items []ZoneRead) *ZoneReadList {
	this := ZoneReadList{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href
	this.Offset = &offset
	this.Limit = &limit
	this.Links = &links
	this.Items = &items

	return &this
}

// NewZoneReadListWithDefaults instantiates a new ZoneReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneReadListWithDefaults() *ZoneReadList {
	this := ZoneReadList{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ZoneReadList) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneReadList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *ZoneReadList) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *ZoneReadList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ZoneReadList) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneReadList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *ZoneReadList) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ZoneReadList) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ZoneReadList) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneReadList) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *ZoneReadList) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *ZoneReadList) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetOffset returns the Offset field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ZoneReadList) GetOffset() *float32 {
	if o == nil {
		return nil
	}

	return o.Offset

}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneReadList) GetOffsetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Offset, true
}

// SetOffset sets field value
func (o *ZoneReadList) SetOffset(v float32) {

	o.Offset = &v

}

// HasOffset returns a boolean if a field has been set.
func (o *ZoneReadList) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ZoneReadList) GetLimit() *float32 {
	if o == nil {
		return nil
	}

	return o.Limit

}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneReadList) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Limit, true
}

// SetLimit sets field value
func (o *ZoneReadList) SetLimit(v float32) {

	o.Limit = &v

}

// HasLimit returns a boolean if a field has been set.
func (o *ZoneReadList) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for Links will be returned
func (o *ZoneReadList) GetLinks() *Links {
	if o == nil {
		return nil
	}

	return o.Links

}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneReadList) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}

	return o.Links, true
}

// SetLinks sets field value
func (o *ZoneReadList) SetLinks(v Links) {

	o.Links = &v

}

// HasLinks returns a boolean if a field has been set.
func (o *ZoneReadList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []ZoneRead will be returned
func (o *ZoneReadList) GetItems() *[]ZoneRead {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneReadList) GetItemsOk() (*[]ZoneRead, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *ZoneReadList) SetItems(v []ZoneRead) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *ZoneReadList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o ZoneReadList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableZoneReadList struct {
	value *ZoneReadList
	isSet bool
}

func (v NullableZoneReadList) Get() *ZoneReadList {
	return v.value
}

func (v *NullableZoneReadList) Set(val *ZoneReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneReadList(val *ZoneReadList) *NullableZoneReadList {
	return &NullableZoneReadList{value: val, isSet: true}
}

func (v NullableZoneReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
