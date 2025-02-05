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
	"fmt"
)

// KskBits Key signing key length in bits. kskBits >= zskBits
type KskBits int32

// List of kskBits
const (
	KSKBITS__1024 KskBits = 1024
	KSKBITS__2048 KskBits = 2048
	KSKBITS__4096 KskBits = 4096
)

func (v *KskBits) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KskBits(value)
	for _, existing := range []KskBits{1024, 2048, 4096} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KskBits", value)
}

// Ptr returns reference to kskBits value
func (v KskBits) Ptr() *KskBits {
	return &v
}

type NullableKskBits struct {
	value *KskBits
	isSet bool
}

func (v NullableKskBits) Get() *KskBits {
	return v.value
}

func (v *NullableKskBits) Set(val *KskBits) {
	v.value = val
	v.isSet = true
}

func (v NullableKskBits) IsSet() bool {
	return v.isSet
}

func (v *NullableKskBits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKskBits(val *KskBits) *NullableKskBits {
	return &NullableKskBits{value: val, isSet: true}
}

func (v NullableKskBits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKskBits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
