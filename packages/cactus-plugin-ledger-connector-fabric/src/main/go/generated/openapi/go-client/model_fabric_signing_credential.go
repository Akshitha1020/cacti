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

// checks if the FabricSigningCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricSigningCredential{}

// FabricSigningCredential struct for FabricSigningCredential
type FabricSigningCredential struct {
	KeychainId string `json:"keychainId"`
	KeychainRef string `json:"keychainRef"`
	Type *FabricSigningCredentialType `json:"type,omitempty"`
	VaultTransitKey *VaultTransitKey `json:"vaultTransitKey,omitempty"`
	WebSocketKey *WebSocketKey `json:"webSocketKey,omitempty"`
}

// NewFabricSigningCredential instantiates a new FabricSigningCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSigningCredential(keychainId string, keychainRef string) *FabricSigningCredential {
	this := FabricSigningCredential{}
	this.KeychainId = keychainId
	this.KeychainRef = keychainRef
	return &this
}

// NewFabricSigningCredentialWithDefaults instantiates a new FabricSigningCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSigningCredentialWithDefaults() *FabricSigningCredential {
	this := FabricSigningCredential{}
	return &this
}

// GetKeychainId returns the KeychainId field value
func (o *FabricSigningCredential) GetKeychainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeychainId
}

// GetKeychainIdOk returns a tuple with the KeychainId field value
// and a boolean to check if the value has been set.
func (o *FabricSigningCredential) GetKeychainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeychainId, true
}

// SetKeychainId sets field value
func (o *FabricSigningCredential) SetKeychainId(v string) {
	o.KeychainId = v
}

// GetKeychainRef returns the KeychainRef field value
func (o *FabricSigningCredential) GetKeychainRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeychainRef
}

// GetKeychainRefOk returns a tuple with the KeychainRef field value
// and a boolean to check if the value has been set.
func (o *FabricSigningCredential) GetKeychainRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeychainRef, true
}

// SetKeychainRef sets field value
func (o *FabricSigningCredential) SetKeychainRef(v string) {
	o.KeychainRef = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FabricSigningCredential) GetType() FabricSigningCredentialType {
	if o == nil || IsNil(o.Type) {
		var ret FabricSigningCredentialType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSigningCredential) GetTypeOk() (*FabricSigningCredentialType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FabricSigningCredential) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FabricSigningCredentialType and assigns it to the Type field.
func (o *FabricSigningCredential) SetType(v FabricSigningCredentialType) {
	o.Type = &v
}

// GetVaultTransitKey returns the VaultTransitKey field value if set, zero value otherwise.
func (o *FabricSigningCredential) GetVaultTransitKey() VaultTransitKey {
	if o == nil || IsNil(o.VaultTransitKey) {
		var ret VaultTransitKey
		return ret
	}
	return *o.VaultTransitKey
}

// GetVaultTransitKeyOk returns a tuple with the VaultTransitKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSigningCredential) GetVaultTransitKeyOk() (*VaultTransitKey, bool) {
	if o == nil || IsNil(o.VaultTransitKey) {
		return nil, false
	}
	return o.VaultTransitKey, true
}

// HasVaultTransitKey returns a boolean if a field has been set.
func (o *FabricSigningCredential) HasVaultTransitKey() bool {
	if o != nil && !IsNil(o.VaultTransitKey) {
		return true
	}

	return false
}

// SetVaultTransitKey gets a reference to the given VaultTransitKey and assigns it to the VaultTransitKey field.
func (o *FabricSigningCredential) SetVaultTransitKey(v VaultTransitKey) {
	o.VaultTransitKey = &v
}

// GetWebSocketKey returns the WebSocketKey field value if set, zero value otherwise.
func (o *FabricSigningCredential) GetWebSocketKey() WebSocketKey {
	if o == nil || IsNil(o.WebSocketKey) {
		var ret WebSocketKey
		return ret
	}
	return *o.WebSocketKey
}

// GetWebSocketKeyOk returns a tuple with the WebSocketKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSigningCredential) GetWebSocketKeyOk() (*WebSocketKey, bool) {
	if o == nil || IsNil(o.WebSocketKey) {
		return nil, false
	}
	return o.WebSocketKey, true
}

// HasWebSocketKey returns a boolean if a field has been set.
func (o *FabricSigningCredential) HasWebSocketKey() bool {
	if o != nil && !IsNil(o.WebSocketKey) {
		return true
	}

	return false
}

// SetWebSocketKey gets a reference to the given WebSocketKey and assigns it to the WebSocketKey field.
func (o *FabricSigningCredential) SetWebSocketKey(v WebSocketKey) {
	o.WebSocketKey = &v
}

func (o FabricSigningCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricSigningCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keychainId"] = o.KeychainId
	toSerialize["keychainRef"] = o.KeychainRef
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.VaultTransitKey) {
		toSerialize["vaultTransitKey"] = o.VaultTransitKey
	}
	if !IsNil(o.WebSocketKey) {
		toSerialize["webSocketKey"] = o.WebSocketKey
	}
	return toSerialize, nil
}

type NullableFabricSigningCredential struct {
	value *FabricSigningCredential
	isSet bool
}

func (v NullableFabricSigningCredential) Get() *FabricSigningCredential {
	return v.value
}

func (v *NullableFabricSigningCredential) Set(val *FabricSigningCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSigningCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSigningCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSigningCredential(val *FabricSigningCredential) *NullableFabricSigningCredential {
	return &NullableFabricSigningCredential{value: val, isSet: true}
}

func (v NullableFabricSigningCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSigningCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


