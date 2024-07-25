/*
Hyperledger Cactus - Keychain API

Contains/describes the Keychain API types/paths for Hyperledger Cactus.

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-keychain-azure-kv

import (
	"encoding/json"
)

// checks if the DeleteKeychainEntryResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteKeychainEntryResponseV1{}

// DeleteKeychainEntryResponseV1 struct for DeleteKeychainEntryResponseV1
type DeleteKeychainEntryResponseV1 struct {
	// The key that was deleted from the keychain.
	Key string `json:"key"`
}

// NewDeleteKeychainEntryResponseV1 instantiates a new DeleteKeychainEntryResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteKeychainEntryResponseV1(key string) *DeleteKeychainEntryResponseV1 {
	this := DeleteKeychainEntryResponseV1{}
	this.Key = key
	return &this
}

// NewDeleteKeychainEntryResponseV1WithDefaults instantiates a new DeleteKeychainEntryResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteKeychainEntryResponseV1WithDefaults() *DeleteKeychainEntryResponseV1 {
	this := DeleteKeychainEntryResponseV1{}
	return &this
}

// GetKey returns the Key field value
func (o *DeleteKeychainEntryResponseV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DeleteKeychainEntryResponseV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DeleteKeychainEntryResponseV1) SetKey(v string) {
	o.Key = v
}

func (o DeleteKeychainEntryResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteKeychainEntryResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableDeleteKeychainEntryResponseV1 struct {
	value *DeleteKeychainEntryResponseV1
	isSet bool
}

func (v NullableDeleteKeychainEntryResponseV1) Get() *DeleteKeychainEntryResponseV1 {
	return v.value
}

func (v *NullableDeleteKeychainEntryResponseV1) Set(val *DeleteKeychainEntryResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteKeychainEntryResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteKeychainEntryResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteKeychainEntryResponseV1(val *DeleteKeychainEntryResponseV1) *NullableDeleteKeychainEntryResponseV1 {
	return &NullableDeleteKeychainEntryResponseV1{value: val, isSet: true}
}

func (v NullableDeleteKeychainEntryResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteKeychainEntryResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


