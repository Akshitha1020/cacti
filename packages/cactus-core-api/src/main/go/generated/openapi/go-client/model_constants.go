/*
Hyperledger Core API

Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-core-api

import (
	"encoding/json"
	"fmt"
)

// Constants the model 'Constants'
type Constants string

// List of Constants
const (
	SocketIoConnectionPathV1 Constants = "/api/v1/async/socket-io/connect"
)

// All allowed values of Constants enum
var AllowedConstantsEnumValues = []Constants{
	"/api/v1/async/socket-io/connect",
}

func (v *Constants) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Constants(value)
	for _, existing := range AllowedConstantsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Constants", value)
}

// NewConstantsFromValue returns a pointer to a valid Constants
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConstantsFromValue(v string) (*Constants, error) {
	ev := Constants(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Constants: valid values are %v", v, AllowedConstantsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Constants) IsValid() bool {
	for _, existing := range AllowedConstantsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Constants value
func (v Constants) Ptr() *Constants {
	return &v
}

type NullableConstants struct {
	value *Constants
	isSet bool
}

func (v NullableConstants) Get() *Constants {
	return v.value
}

func (v *NullableConstants) Set(val *Constants) {
	v.value = val
	v.isSet = true
}

func (v NullableConstants) IsSet() bool {
	return v.isSet
}

func (v *NullableConstants) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstants(val *Constants) *NullableConstants {
	return &NullableConstants{value: val, isSet: true}
}

func (v NullableConstants) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstants) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

