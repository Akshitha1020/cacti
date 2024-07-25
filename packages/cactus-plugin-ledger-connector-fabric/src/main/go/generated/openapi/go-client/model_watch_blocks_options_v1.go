/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the WatchBlocksOptionsV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksOptionsV1{}

// WatchBlocksOptionsV1 Options passed when subscribing to block monitoring.
type WatchBlocksOptionsV1 struct {
	// Hyperledger Fabric channel to connect to.
	ChannelName string `json:"channelName"`
	GatewayOptions GatewayOptions `json:"gatewayOptions"`
	Type WatchBlocksListenerTypeV1 `json:"type"`
	// From which block start monitoring. Defaults to latest.
	StartBlock *string `json:"startBlock,omitempty"`
}

// NewWatchBlocksOptionsV1 instantiates a new WatchBlocksOptionsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksOptionsV1(channelName string, gatewayOptions GatewayOptions, type_ WatchBlocksListenerTypeV1) *WatchBlocksOptionsV1 {
	this := WatchBlocksOptionsV1{}
	this.ChannelName = channelName
	this.GatewayOptions = gatewayOptions
	this.Type = type_
	return &this
}

// NewWatchBlocksOptionsV1WithDefaults instantiates a new WatchBlocksOptionsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksOptionsV1WithDefaults() *WatchBlocksOptionsV1 {
	this := WatchBlocksOptionsV1{}
	return &this
}

// GetChannelName returns the ChannelName field value
func (o *WatchBlocksOptionsV1) GetChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
func (o *WatchBlocksOptionsV1) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelName, true
}

// SetChannelName sets field value
func (o *WatchBlocksOptionsV1) SetChannelName(v string) {
	o.ChannelName = v
}

// GetGatewayOptions returns the GatewayOptions field value
func (o *WatchBlocksOptionsV1) GetGatewayOptions() GatewayOptions {
	if o == nil {
		var ret GatewayOptions
		return ret
	}

	return o.GatewayOptions
}

// GetGatewayOptionsOk returns a tuple with the GatewayOptions field value
// and a boolean to check if the value has been set.
func (o *WatchBlocksOptionsV1) GetGatewayOptionsOk() (*GatewayOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayOptions, true
}

// SetGatewayOptions sets field value
func (o *WatchBlocksOptionsV1) SetGatewayOptions(v GatewayOptions) {
	o.GatewayOptions = v
}

// GetType returns the Type field value
func (o *WatchBlocksOptionsV1) GetType() WatchBlocksListenerTypeV1 {
	if o == nil {
		var ret WatchBlocksListenerTypeV1
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WatchBlocksOptionsV1) GetTypeOk() (*WatchBlocksListenerTypeV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WatchBlocksOptionsV1) SetType(v WatchBlocksListenerTypeV1) {
	o.Type = v
}

// GetStartBlock returns the StartBlock field value if set, zero value otherwise.
func (o *WatchBlocksOptionsV1) GetStartBlock() string {
	if o == nil || IsNil(o.StartBlock) {
		var ret string
		return ret
	}
	return *o.StartBlock
}

// GetStartBlockOk returns a tuple with the StartBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksOptionsV1) GetStartBlockOk() (*string, bool) {
	if o == nil || IsNil(o.StartBlock) {
		return nil, false
	}
	return o.StartBlock, true
}

// HasStartBlock returns a boolean if a field has been set.
func (o *WatchBlocksOptionsV1) HasStartBlock() bool {
	if o != nil && !IsNil(o.StartBlock) {
		return true
	}

	return false
}

// SetStartBlock gets a reference to the given string and assigns it to the StartBlock field.
func (o *WatchBlocksOptionsV1) SetStartBlock(v string) {
	o.StartBlock = &v
}

func (o WatchBlocksOptionsV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksOptionsV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channelName"] = o.ChannelName
	toSerialize["gatewayOptions"] = o.GatewayOptions
	toSerialize["type"] = o.Type
	if !IsNil(o.StartBlock) {
		toSerialize["startBlock"] = o.StartBlock
	}
	return toSerialize, nil
}

type NullableWatchBlocksOptionsV1 struct {
	value *WatchBlocksOptionsV1
	isSet bool
}

func (v NullableWatchBlocksOptionsV1) Get() *WatchBlocksOptionsV1 {
	return v.value
}

func (v *NullableWatchBlocksOptionsV1) Set(val *WatchBlocksOptionsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksOptionsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksOptionsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksOptionsV1(val *WatchBlocksOptionsV1) *NullableWatchBlocksOptionsV1 {
	return &NullableWatchBlocksOptionsV1{value: val, isSet: true}
}

func (v NullableWatchBlocksOptionsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksOptionsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


