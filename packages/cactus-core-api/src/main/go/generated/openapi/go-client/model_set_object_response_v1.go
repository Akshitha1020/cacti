/*
Hyperledger Core API

Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-core-api

import (
	"encoding/json"
)

// checks if the SetObjectResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetObjectResponseV1{}

// SetObjectResponseV1 struct for SetObjectResponseV1
type SetObjectResponseV1 struct {
	// The key that was used to set the value in the object store.
	Key string `json:"key"`
}

// NewSetObjectResponseV1 instantiates a new SetObjectResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetObjectResponseV1(key string) *SetObjectResponseV1 {
	this := SetObjectResponseV1{}
	this.Key = key
	return &this
}

// NewSetObjectResponseV1WithDefaults instantiates a new SetObjectResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetObjectResponseV1WithDefaults() *SetObjectResponseV1 {
	this := SetObjectResponseV1{}
	return &this
}

// GetKey returns the Key field value
func (o *SetObjectResponseV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SetObjectResponseV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SetObjectResponseV1) SetKey(v string) {
	o.Key = v
}

func (o SetObjectResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetObjectResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableSetObjectResponseV1 struct {
	value *SetObjectResponseV1
	isSet bool
}

func (v NullableSetObjectResponseV1) Get() *SetObjectResponseV1 {
	return v.value
}

func (v *NullableSetObjectResponseV1) Set(val *SetObjectResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSetObjectResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSetObjectResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetObjectResponseV1(val *SetObjectResponseV1) *NullableSetObjectResponseV1 {
	return &NullableSetObjectResponseV1{value: val, isSet: true}
}

func (v NullableSetObjectResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetObjectResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


