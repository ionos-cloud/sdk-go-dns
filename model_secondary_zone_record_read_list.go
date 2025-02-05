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

// SecondaryZoneRecordReadList struct for SecondaryZoneRecordReadList
type SecondaryZoneRecordReadList struct {
	// The resource's unique identifier.
	Id       *string                              `json:"id"`
	Type     *string                              `json:"type"`
	Href     *string                              `json:"href"`
	Metadata *SecondaryZoneRecordReadListMetadata `json:"metadata"`
	Items    *[]SecondaryZoneRecordRead           `json:"items"`
	// Pagination offset.
	Offset *float32 `json:"offset"`
	// Pagination limit.
	Limit *float32 `json:"limit"`
	Links *Links   `json:"_links"`
}

// NewSecondaryZoneRecordReadList instantiates a new SecondaryZoneRecordReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecondaryZoneRecordReadList(id string, type_ string, href string, metadata SecondaryZoneRecordReadListMetadata, items []SecondaryZoneRecordRead, offset float32, limit float32, links Links) *SecondaryZoneRecordReadList {
	this := SecondaryZoneRecordReadList{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href
	this.Metadata = &metadata
	this.Items = &items
	this.Offset = &offset
	this.Limit = &limit
	this.Links = &links

	return &this
}

// NewSecondaryZoneRecordReadListWithDefaults instantiates a new SecondaryZoneRecordReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecondaryZoneRecordReadListWithDefaults() *SecondaryZoneRecordReadList {
	this := SecondaryZoneRecordReadList{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SecondaryZoneRecordReadList) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecondaryZoneRecordReadList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *SecondaryZoneRecordReadList) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *SecondaryZoneRecordReadList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SecondaryZoneRecordReadList) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecondaryZoneRecordReadList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *SecondaryZoneRecordReadList) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *SecondaryZoneRecordReadList) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SecondaryZoneRecordReadList) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecondaryZoneRecordReadList) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *SecondaryZoneRecordReadList) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *SecondaryZoneRecordReadList) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for SecondaryZoneRecordReadListMetadata will be returned
func (o *SecondaryZoneRecordReadList) GetMetadata() *SecondaryZoneRecordReadListMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecondaryZoneRecordReadList) GetMetadataOk() (*SecondaryZoneRecordReadListMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *SecondaryZoneRecordReadList) SetMetadata(v SecondaryZoneRecordReadListMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *SecondaryZoneRecordReadList) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []SecondaryZoneRecordRead will be returned
func (o *SecondaryZoneRecordReadList) GetItems() *[]SecondaryZoneRecordRead {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecondaryZoneRecordReadList) GetItemsOk() (*[]SecondaryZoneRecordRead, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *SecondaryZoneRecordReadList) SetItems(v []SecondaryZoneRecordRead) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *SecondaryZoneRecordReadList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// GetOffset returns the Offset field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SecondaryZoneRecordReadList) GetOffset() *float32 {
	if o == nil {
		return nil
	}

	return o.Offset

}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecondaryZoneRecordReadList) GetOffsetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Offset, true
}

// SetOffset sets field value
func (o *SecondaryZoneRecordReadList) SetOffset(v float32) {

	o.Offset = &v

}

// HasOffset returns a boolean if a field has been set.
func (o *SecondaryZoneRecordReadList) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SecondaryZoneRecordReadList) GetLimit() *float32 {
	if o == nil {
		return nil
	}

	return o.Limit

}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecondaryZoneRecordReadList) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Limit, true
}

// SetLimit sets field value
func (o *SecondaryZoneRecordReadList) SetLimit(v float32) {

	o.Limit = &v

}

// HasLimit returns a boolean if a field has been set.
func (o *SecondaryZoneRecordReadList) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for Links will be returned
func (o *SecondaryZoneRecordReadList) GetLinks() *Links {
	if o == nil {
		return nil
	}

	return o.Links

}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecondaryZoneRecordReadList) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}

	return o.Links, true
}

// SetLinks sets field value
func (o *SecondaryZoneRecordReadList) SetLinks(v Links) {

	o.Links = &v

}

// HasLinks returns a boolean if a field has been set.
func (o *SecondaryZoneRecordReadList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

func (o SecondaryZoneRecordReadList) MarshalJSON() ([]byte, error) {
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

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
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

	return json.Marshal(toSerialize)
}

type NullableSecondaryZoneRecordReadList struct {
	value *SecondaryZoneRecordReadList
	isSet bool
}

func (v NullableSecondaryZoneRecordReadList) Get() *SecondaryZoneRecordReadList {
	return v.value
}

func (v *NullableSecondaryZoneRecordReadList) Set(val *SecondaryZoneRecordReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableSecondaryZoneRecordReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableSecondaryZoneRecordReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecondaryZoneRecordReadList(val *SecondaryZoneRecordReadList) *NullableSecondaryZoneRecordReadList {
	return &NullableSecondaryZoneRecordReadList{value: val, isSet: true}
}

func (v NullableSecondaryZoneRecordReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecondaryZoneRecordReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
