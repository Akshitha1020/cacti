/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
)

// checks if the WatchBlocksV1Options type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksV1Options{}

// WatchBlocksV1Options struct for WatchBlocksV1Options
type WatchBlocksV1Options struct {
	// Include entire block data if flag is true, otherwise just a header is returned (default)
	GetBlockData *bool `json:"getBlockData,omitempty"`
	// Block from which we want to start the monitoring process.
	LastSeenBlock *float32 `json:"lastSeenBlock,omitempty"`
	// How often to poll ethereum node for new blocks. Not used if the node supports subscription based monitoring (i.e. WebSocket).
	HttpPollInterval *float32 `json:"httpPollInterval,omitempty"`
}

// NewWatchBlocksV1Options instantiates a new WatchBlocksV1Options object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksV1Options() *WatchBlocksV1Options {
	this := WatchBlocksV1Options{}
	var getBlockData bool = false
	this.GetBlockData = &getBlockData
	return &this
}

// NewWatchBlocksV1OptionsWithDefaults instantiates a new WatchBlocksV1Options object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksV1OptionsWithDefaults() *WatchBlocksV1Options {
	this := WatchBlocksV1Options{}
	var getBlockData bool = false
	this.GetBlockData = &getBlockData
	return &this
}

// GetGetBlockData returns the GetBlockData field value if set, zero value otherwise.
func (o *WatchBlocksV1Options) GetGetBlockData() bool {
	if o == nil || IsNil(o.GetBlockData) {
		var ret bool
		return ret
	}
	return *o.GetBlockData
}

// GetGetBlockDataOk returns a tuple with the GetBlockData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksV1Options) GetGetBlockDataOk() (*bool, bool) {
	if o == nil || IsNil(o.GetBlockData) {
		return nil, false
	}
	return o.GetBlockData, true
}

// HasGetBlockData returns a boolean if a field has been set.
func (o *WatchBlocksV1Options) HasGetBlockData() bool {
	if o != nil && !IsNil(o.GetBlockData) {
		return true
	}

	return false
}

// SetGetBlockData gets a reference to the given bool and assigns it to the GetBlockData field.
func (o *WatchBlocksV1Options) SetGetBlockData(v bool) {
	o.GetBlockData = &v
}

// GetLastSeenBlock returns the LastSeenBlock field value if set, zero value otherwise.
func (o *WatchBlocksV1Options) GetLastSeenBlock() float32 {
	if o == nil || IsNil(o.LastSeenBlock) {
		var ret float32
		return ret
	}
	return *o.LastSeenBlock
}

// GetLastSeenBlockOk returns a tuple with the LastSeenBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksV1Options) GetLastSeenBlockOk() (*float32, bool) {
	if o == nil || IsNil(o.LastSeenBlock) {
		return nil, false
	}
	return o.LastSeenBlock, true
}

// HasLastSeenBlock returns a boolean if a field has been set.
func (o *WatchBlocksV1Options) HasLastSeenBlock() bool {
	if o != nil && !IsNil(o.LastSeenBlock) {
		return true
	}

	return false
}

// SetLastSeenBlock gets a reference to the given float32 and assigns it to the LastSeenBlock field.
func (o *WatchBlocksV1Options) SetLastSeenBlock(v float32) {
	o.LastSeenBlock = &v
}

// GetHttpPollInterval returns the HttpPollInterval field value if set, zero value otherwise.
func (o *WatchBlocksV1Options) GetHttpPollInterval() float32 {
	if o == nil || IsNil(o.HttpPollInterval) {
		var ret float32
		return ret
	}
	return *o.HttpPollInterval
}

// GetHttpPollIntervalOk returns a tuple with the HttpPollInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksV1Options) GetHttpPollIntervalOk() (*float32, bool) {
	if o == nil || IsNil(o.HttpPollInterval) {
		return nil, false
	}
	return o.HttpPollInterval, true
}

// HasHttpPollInterval returns a boolean if a field has been set.
func (o *WatchBlocksV1Options) HasHttpPollInterval() bool {
	if o != nil && !IsNil(o.HttpPollInterval) {
		return true
	}

	return false
}

// SetHttpPollInterval gets a reference to the given float32 and assigns it to the HttpPollInterval field.
func (o *WatchBlocksV1Options) SetHttpPollInterval(v float32) {
	o.HttpPollInterval = &v
}

func (o WatchBlocksV1Options) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksV1Options) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GetBlockData) {
		toSerialize["getBlockData"] = o.GetBlockData
	}
	if !IsNil(o.LastSeenBlock) {
		toSerialize["lastSeenBlock"] = o.LastSeenBlock
	}
	if !IsNil(o.HttpPollInterval) {
		toSerialize["httpPollInterval"] = o.HttpPollInterval
	}
	return toSerialize, nil
}

type NullableWatchBlocksV1Options struct {
	value *WatchBlocksV1Options
	isSet bool
}

func (v NullableWatchBlocksV1Options) Get() *WatchBlocksV1Options {
	return v.value
}

func (v *NullableWatchBlocksV1Options) Set(val *WatchBlocksV1Options) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksV1Options) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksV1Options) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksV1Options(val *WatchBlocksV1Options) *NullableWatchBlocksV1Options {
	return &NullableWatchBlocksV1Options{value: val, isSet: true}
}

func (v NullableWatchBlocksV1Options) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksV1Options) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


