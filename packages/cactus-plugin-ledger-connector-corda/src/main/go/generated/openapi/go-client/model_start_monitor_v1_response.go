/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the StartMonitorV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartMonitorV1Response{}

// StartMonitorV1Response struct for StartMonitorV1Response
type StartMonitorV1Response struct {
	// Flag set to true if monitoring started correctly.
	Success bool `json:"success"`
	// Message describing operation status or any errors that occurred.
	Msg string `json:"msg"`
}

// NewStartMonitorV1Response instantiates a new StartMonitorV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartMonitorV1Response(success bool, msg string) *StartMonitorV1Response {
	this := StartMonitorV1Response{}
	this.Success = success
	this.Msg = msg
	return &this
}

// NewStartMonitorV1ResponseWithDefaults instantiates a new StartMonitorV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartMonitorV1ResponseWithDefaults() *StartMonitorV1Response {
	this := StartMonitorV1Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *StartMonitorV1Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *StartMonitorV1Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *StartMonitorV1Response) SetSuccess(v bool) {
	o.Success = v
}

// GetMsg returns the Msg field value
func (o *StartMonitorV1Response) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *StartMonitorV1Response) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *StartMonitorV1Response) SetMsg(v string) {
	o.Msg = v
}

func (o StartMonitorV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartMonitorV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["msg"] = o.Msg
	return toSerialize, nil
}

type NullableStartMonitorV1Response struct {
	value *StartMonitorV1Response
	isSet bool
}

func (v NullableStartMonitorV1Response) Get() *StartMonitorV1Response {
	return v.value
}

func (v *NullableStartMonitorV1Response) Set(val *StartMonitorV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableStartMonitorV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableStartMonitorV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartMonitorV1Response(val *StartMonitorV1Response) *NullableStartMonitorV1Response {
	return &NullableStartMonitorV1Response{value: val, isSet: true}
}

func (v NullableStartMonitorV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartMonitorV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


