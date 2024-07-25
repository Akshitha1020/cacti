/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
	"fmt"
)

// WatchBlocksV1BlockDataTimestamp - struct for WatchBlocksV1BlockDataTimestamp
type WatchBlocksV1BlockDataTimestamp struct {
	Float32 *float32
	String *string
}

// float32AsWatchBlocksV1BlockDataTimestamp is a convenience function that returns float32 wrapped in WatchBlocksV1BlockDataTimestamp
func Float32AsWatchBlocksV1BlockDataTimestamp(v *float32) WatchBlocksV1BlockDataTimestamp {
	return WatchBlocksV1BlockDataTimestamp{
		Float32: v,
	}
}

// stringAsWatchBlocksV1BlockDataTimestamp is a convenience function that returns string wrapped in WatchBlocksV1BlockDataTimestamp
func StringAsWatchBlocksV1BlockDataTimestamp(v *string) WatchBlocksV1BlockDataTimestamp {
	return WatchBlocksV1BlockDataTimestamp{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WatchBlocksV1BlockDataTimestamp) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(WatchBlocksV1BlockDataTimestamp)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(WatchBlocksV1BlockDataTimestamp)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WatchBlocksV1BlockDataTimestamp) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WatchBlocksV1BlockDataTimestamp) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableWatchBlocksV1BlockDataTimestamp struct {
	value *WatchBlocksV1BlockDataTimestamp
	isSet bool
}

func (v NullableWatchBlocksV1BlockDataTimestamp) Get() *WatchBlocksV1BlockDataTimestamp {
	return v.value
}

func (v *NullableWatchBlocksV1BlockDataTimestamp) Set(val *WatchBlocksV1BlockDataTimestamp) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksV1BlockDataTimestamp) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksV1BlockDataTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksV1BlockDataTimestamp(val *WatchBlocksV1BlockDataTimestamp) *NullableWatchBlocksV1BlockDataTimestamp {
	return &NullableWatchBlocksV1BlockDataTimestamp{value: val, isSet: true}
}

func (v NullableWatchBlocksV1BlockDataTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksV1BlockDataTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


