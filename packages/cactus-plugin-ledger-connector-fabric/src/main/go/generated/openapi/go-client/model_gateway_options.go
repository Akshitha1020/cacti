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

// checks if the GatewayOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayOptions{}

// GatewayOptions struct for GatewayOptions
type GatewayOptions struct {
	ConnectionProfile *ConnectionProfile `json:"connectionProfile,omitempty"`
	Discovery *GatewayDiscoveryOptions `json:"discovery,omitempty"`
	EventHandlerOptions *GatewayEventHandlerOptions `json:"eventHandlerOptions,omitempty"`
	Identity string `json:"identity"`
	Wallet GatewayOptionsWallet `json:"wallet"`
}

// NewGatewayOptions instantiates a new GatewayOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayOptions(identity string, wallet GatewayOptionsWallet) *GatewayOptions {
	this := GatewayOptions{}
	this.Identity = identity
	this.Wallet = wallet
	return &this
}

// NewGatewayOptionsWithDefaults instantiates a new GatewayOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayOptionsWithDefaults() *GatewayOptions {
	this := GatewayOptions{}
	return &this
}

// GetConnectionProfile returns the ConnectionProfile field value if set, zero value otherwise.
func (o *GatewayOptions) GetConnectionProfile() ConnectionProfile {
	if o == nil || IsNil(o.ConnectionProfile) {
		var ret ConnectionProfile
		return ret
	}
	return *o.ConnectionProfile
}

// GetConnectionProfileOk returns a tuple with the ConnectionProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayOptions) GetConnectionProfileOk() (*ConnectionProfile, bool) {
	if o == nil || IsNil(o.ConnectionProfile) {
		return nil, false
	}
	return o.ConnectionProfile, true
}

// HasConnectionProfile returns a boolean if a field has been set.
func (o *GatewayOptions) HasConnectionProfile() bool {
	if o != nil && !IsNil(o.ConnectionProfile) {
		return true
	}

	return false
}

// SetConnectionProfile gets a reference to the given ConnectionProfile and assigns it to the ConnectionProfile field.
func (o *GatewayOptions) SetConnectionProfile(v ConnectionProfile) {
	o.ConnectionProfile = &v
}

// GetDiscovery returns the Discovery field value if set, zero value otherwise.
func (o *GatewayOptions) GetDiscovery() GatewayDiscoveryOptions {
	if o == nil || IsNil(o.Discovery) {
		var ret GatewayDiscoveryOptions
		return ret
	}
	return *o.Discovery
}

// GetDiscoveryOk returns a tuple with the Discovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayOptions) GetDiscoveryOk() (*GatewayDiscoveryOptions, bool) {
	if o == nil || IsNil(o.Discovery) {
		return nil, false
	}
	return o.Discovery, true
}

// HasDiscovery returns a boolean if a field has been set.
func (o *GatewayOptions) HasDiscovery() bool {
	if o != nil && !IsNil(o.Discovery) {
		return true
	}

	return false
}

// SetDiscovery gets a reference to the given GatewayDiscoveryOptions and assigns it to the Discovery field.
func (o *GatewayOptions) SetDiscovery(v GatewayDiscoveryOptions) {
	o.Discovery = &v
}

// GetEventHandlerOptions returns the EventHandlerOptions field value if set, zero value otherwise.
func (o *GatewayOptions) GetEventHandlerOptions() GatewayEventHandlerOptions {
	if o == nil || IsNil(o.EventHandlerOptions) {
		var ret GatewayEventHandlerOptions
		return ret
	}
	return *o.EventHandlerOptions
}

// GetEventHandlerOptionsOk returns a tuple with the EventHandlerOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayOptions) GetEventHandlerOptionsOk() (*GatewayEventHandlerOptions, bool) {
	if o == nil || IsNil(o.EventHandlerOptions) {
		return nil, false
	}
	return o.EventHandlerOptions, true
}

// HasEventHandlerOptions returns a boolean if a field has been set.
func (o *GatewayOptions) HasEventHandlerOptions() bool {
	if o != nil && !IsNil(o.EventHandlerOptions) {
		return true
	}

	return false
}

// SetEventHandlerOptions gets a reference to the given GatewayEventHandlerOptions and assigns it to the EventHandlerOptions field.
func (o *GatewayOptions) SetEventHandlerOptions(v GatewayEventHandlerOptions) {
	o.EventHandlerOptions = &v
}

// GetIdentity returns the Identity field value
func (o *GatewayOptions) GetIdentity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *GatewayOptions) GetIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *GatewayOptions) SetIdentity(v string) {
	o.Identity = v
}

// GetWallet returns the Wallet field value
func (o *GatewayOptions) GetWallet() GatewayOptionsWallet {
	if o == nil {
		var ret GatewayOptionsWallet
		return ret
	}

	return o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value
// and a boolean to check if the value has been set.
func (o *GatewayOptions) GetWalletOk() (*GatewayOptionsWallet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallet, true
}

// SetWallet sets field value
func (o *GatewayOptions) SetWallet(v GatewayOptionsWallet) {
	o.Wallet = v
}

func (o GatewayOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConnectionProfile) {
		toSerialize["connectionProfile"] = o.ConnectionProfile
	}
	if !IsNil(o.Discovery) {
		toSerialize["discovery"] = o.Discovery
	}
	if !IsNil(o.EventHandlerOptions) {
		toSerialize["eventHandlerOptions"] = o.EventHandlerOptions
	}
	toSerialize["identity"] = o.Identity
	toSerialize["wallet"] = o.Wallet
	return toSerialize, nil
}

type NullableGatewayOptions struct {
	value *GatewayOptions
	isSet bool
}

func (v NullableGatewayOptions) Get() *GatewayOptions {
	return v.value
}

func (v *NullableGatewayOptions) Set(val *GatewayOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayOptions(val *GatewayOptions) *NullableGatewayOptions {
	return &NullableGatewayOptions{value: val, isSet: true}
}

func (v NullableGatewayOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


